//go:build ignore

package main

import (
	"flag"
	"fmt"

	"github.com/progrium/macdriver/generate/modules"
)

// go run ./generate/tools/constant.go <platform> <framework> <constant>
func main() {
	flag.Parse()
	c := modules.LookupConstant(flag.Arg(0), flag.Arg(1), flag.Arg(2))
	if c != nil {
		if c.ArmValue != "" {
			fmt.Println("Value[amd64]:", c.Value)
			fmt.Println("Value[arm64]:", c.ArmValue)
		} else {
			fmt.Println("Value:", c.Value)
		}
	}

}
