package generate

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Symbol struct {
	Name string
	Path string
	Kind string

	Description string
	Type        string
	Parent      string
	Modules     []string
	Platforms   []Platform
	Declaration string
	Parameters  []Parameter
	Return      string
	Deprecated  bool
}

type Platform struct {
	Name         string
	IntroducedAt string
	Current      string
	Beta         bool
	Deprecated   bool
	DeprecatedAt string
}

type Parameter struct {
	Name        string
	Description string
}

func (s Symbol) DocURL() string {
	return fmt.Sprintf("https://developer.apple.com/documentation/%s?language=objc", s.Path)
}

func (s Symbol) HasPlatform(name string) bool {
	for _, p := range s.Platforms {
		if strings.ToLower(p.Name) == strings.ToLower(name) {
			return true
		}
	}
	return false
}

func LoadSymbolFrom(file *zip.File) (v Symbol, err error) {
	var reader io.ReadCloser
	reader, err = file.Open()
	if err != nil {
		return v, err
	}
	defer reader.Close()

	b, err := io.ReadAll(reader)
	if err != nil {
		return
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return v, err
	}
	return
}
