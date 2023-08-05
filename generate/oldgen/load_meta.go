package oldgen

import (
	"embed"
	"encoding/json"
	"io"
	"strings"

	"github.com/progrium/macdriver/generate/oldgen/data"
)

//go:embed meta
var metaFS embed.FS

type Module struct {
	Name  string          // module name
	Types []data.TypeInfo // types in this module
}

// load all module meta datas
func LoadAllMeta() []Module {
	entries, err := metaFS.ReadDir("meta")
	if err != nil {
		panic(err)
	}
	var moduleTypeses []Module
	for _, de := range entries {
		mts := loadModule(de.Name())
		moduleTypeses = append(moduleTypeses, mts)
	}
	return moduleTypeses
}

func loadModule(name string) Module {
	//log.Println("load module:", name)
	entries, err := metaFS.ReadDir("meta/" + name)
	if err != nil {
		panic(err)
	}
	var typeInfos []data.TypeInfo
	for _, e := range entries {
		ti := loadOneTypeInfo(name, strings.TrimSuffix(e.Name(), ".json"))
		typeInfos = append(typeInfos, ti)
	}
	return Module{
		Name:  name,
		Types: typeInfos,
	}
}

var typeInfoCache map[string]data.TypeInfo = map[string]data.TypeInfo{}

func loadOneTypeInfo(module string, name string) data.TypeInfo {
	key := module + "." + name
	ti, ok := typeInfoCache[key]
	if ok {
		return ti
	}
	ti = loadOneTypeInner(module, name)
	typeInfoCache[key] = ti
	return ti
}

func loadOneTypeInner(module string, name string) data.TypeInfo {
	//log.Println("load type meta:", name)
	f, err := metaFS.Open("meta/" + module + "/" + name + ".json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	d, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var kt data.KindType
	err = json.Unmarshal(d, &kt)
	if err != nil {
		panic(err)
	}
	switch kt.Kind {
	case data.KindClass:
		var c data.Class
		err = json.Unmarshal(d, &c)
		if err != nil {
			panic(err)
		}
		return &c
	case data.KindProtocol:
		var p data.Protocol
		err = json.Unmarshal(d, &p)
		if err != nil {
			panic(err)
		}
		return &p
	case data.KindAlias:
		var a data.Alias
		err = json.Unmarshal(d, &a)
		if err != nil {
			panic(err)
		}
		return &a
	case data.KindStruct:
		var s data.Struct
		err = json.Unmarshal(d, &s)
		if err != nil {
			panic(err)
		}
		return &s
	default:
		panic("unkown type: " + string(kt.Kind))
	}
}
