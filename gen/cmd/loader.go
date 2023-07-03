package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/progrium/macdriver/gen"
	"github.com/progrium/macschema/schema"
)

type pkg struct {
	Name     string
	Contents []schemaLoader
}

type schemaLoader func() (*schema.Schema, error)

func (s schemaLoader) Then(fn schemaUpdater) schemaLoader {
	return func() (*schema.Schema, error) {
		in, err := s()
		if err != nil {
			return nil, err
		}
		if err := fn(in); err != nil {
			return nil, err
		}
		return in, err
	}
}

type schemaUpdater func(s *schema.Schema) error

func (su schemaUpdater) Then(next schemaUpdater) schemaUpdater {
	return func(s *schema.Schema) error {
		if err := su(s); err != nil {
			return err
		}
		return next(s)
	}
}

func loadFileBase(filename string) schemaLoader {
	return func() (*schema.Schema, error) {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		var input schema.Schema
		if err := json.Unmarshal(data, &input); err != nil {
			return nil, err
		}
		return &input, nil
	}
}

func loadFile(filename string) schemaLoader {
	return loadFileBase(filename).Then(filterDeprecated)
}

func filterClassMethods(pred func(schema.Method) bool) schemaUpdater {
	filter := func(in []schema.Method) []schema.Method {
		var out []schema.Method
		for _, m := range in {
			if pred(m) {
				out = append(out, m)
			}
		}
		return out
	}
	return func(s *schema.Schema) error {
		if s.Class == nil {
			return nil
		}
		s.Class.TypeMethods = filter(s.Class.TypeMethods)
		s.Class.InstanceMethods = filter(s.Class.InstanceMethods)
		return nil
	}
}

func filterClassProperties(pred func(schema.Property) bool) schemaUpdater {
	filter := func(in []schema.Property) []schema.Property {
		var out []schema.Property
		for _, m := range in {
			if pred(m) {
				out = append(out, m)
			}
		}
		return out
	}
	return func(s *schema.Schema) error {
		if s.Class == nil {
			return nil
		}
		s.Class.TypeProperties = filter(s.Class.TypeProperties)
		s.Class.InstanceProperties = filter(s.Class.InstanceProperties)
		return nil
	}
}

// filterDeprecated removes deprecated methods and properties from a class.
var filterDeprecated = filterClassMethods(func(m schema.Method) bool {
	return !m.Deprecated
}).Then(filterClassProperties(func(p schema.Property) bool {
	return !p.Deprecated
}))

func loadSchemas(contents []schemaLoader) ([]*schema.Schema, error) {
	schemas := make([]*schema.Schema, len(contents))
	for i, loader := range contents {
		input, err := loader()
		if err != nil {
			return nil, err
		}
		schemas[i] = input
	}
	return schemas, nil
}

func definedClasses(schemas []*schema.Schema) map[string]bool {
	r := map[string]bool{}
	for _, input := range schemas {
		if input.Class != nil {
			r[input.Class.Name] = true
		}
	}
	return r
}

func generate(basePackage string, packages []pkg) error {
	var imports []gen.PackageContents
	for _, p := range packages {
		schemas, err := loadSchemas(p.Contents)
		if err != nil {
			return fmt.Errorf("loading schemas for package %q: %w", p.Name, err)
		}
		if err := generatePackage(p.Name, schemas, imports); err != nil {
			return err
		}
		// after generating this package, add its contents to imports available to
		// later packages
		imports = append(imports, gen.PackageContents{
			Import: &gen.Import{
				Path:  fmt.Sprintf("%s/%s", basePackage, p.Name),
				Alias: p.Name,
			},
			Classes: definedClasses(schemas),
		})
	}
	return nil
}

func generatePackage(name string, schemas []*schema.Schema, imports []gen.PackageContents) error {
	if err := os.MkdirAll(name, fs.ModePerm); err != nil {
		return err
	}

	defaultImport := gen.PackageContents{
		Import:  nil,
		Classes: definedClasses(schemas),
	}

	combinedImports := append([]gen.PackageContents{defaultImport}, imports...)

	desc := gen.PackageDescription{
		Name: name,
	}
	seenFrameworks := make(map[string]bool)
	addFramework := func(fw string) {
		if seenFrameworks[fw] {
			return
		}
		seenFrameworks[fw] = true
		desc.LinkFrameworks = append(desc.LinkFrameworks, fw)
		desc.CIncludes = append(desc.CIncludes, fmt.Sprintf("%s/%s.h", fw, fw))
	}
	if name == "core" {
		// HACK: the types in the "core" package are defined in "Foundation", but
		// some have methods contributed by "AppKit". We don't currenly have the
		// framework info at the method-level to detect this, so just manually
		// inject the AppKit dependency here so that those methods will compile.
		addFramework("AppKit")
	}
	for _, input := range schemas {
		if input.Class == nil {
			continue
		}
		for _, fw := range input.Class.Frameworks {
			fw = strings.ReplaceAll(fw, " ", "")
			// FIXME is there a better way to determine which includes and frameworks
			// should go in the Objective-C code?
			switch fw {
			case "CoreAnimation":
				fw = "QuartzCore"
			case "UIKit":
				continue
			}
			addFramework(fw)
		}
	}
	pkg, err := gen.Convert(desc, combinedImports, schemas...)
	if err != nil {
		return fmt.Errorf("error converting package %s: %w", name, err)
	}
	outPath := path.Join(name, desc.Name+"_objc.gen.go")
	var b bytes.Buffer
	if err := pkg.Generate(&b); err != nil {
		return fmt.Errorf("error generating package %s: %w", name, err)
	}
	code, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(outPath, code, fs.ModePerm); err != nil { // FIXME file mode?
		return err
	}
	return nil
}
