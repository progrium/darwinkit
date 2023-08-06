package modules

import (
	"bufio"
	"embed"
	"fmt"
	"strings"
)

type Constant struct {
	Value    string
	ArmValue string
}

//go:embed enums
var enums embed.FS

var constantCache = map[string]*Constant{}

func loadConstants(filepath string) bool {
	arm := strings.HasSuffix(filepath, "_arm64")
	f, err := enums.Open("enums/" + filepath)
	if err != nil {
		return false
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), " ", 2)
		if len(parts) == 2 {
			k := strings.ToLower(fmt.Sprintf("%s/%s", strings.TrimSuffix(filepath, "_arm64"), parts[0]))
			_, ok := constantCache[k]
			if !ok {
				constantCache[k] = &Constant{}
			}
			if arm {
				constantCache[k].ArmValue = parts[1]
			} else {
				constantCache[k].Value = parts[1]
			}
		}
	}
	return true
}

func LookupConstant(platform, framework, name string) *Constant {
	filepath := fmt.Sprintf("%s/%s", platform, framework)
	key := strings.ToLower(fmt.Sprintf("%s/%s", filepath, name))
	if c, ok := constantCache[key]; ok {
		return c
	}
	if !loadConstants(filepath) {
		return nil
	}
	loadConstants(fmt.Sprintf("%s_arm64", filepath))
	if c, ok := constantCache[key]; ok {
		return c
	}
	return nil
}
