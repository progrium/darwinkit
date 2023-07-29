package main

import (
	"fmt"

	"github.com/progrium/macdriver/core"
)

func main() {
	ud := core.NSUserDefaults_StandardUserDefaults()
	ud.SetURLForKey(core.URL("https://github.com/progrium/macdriver"), "macdriver")
	u := ud.URLForKey("macdriver")
	fmt.Println("looked up 'macdriver' key in NSUserDefaults and got:", u)
}
