package main

import (
	"fmt"

	"github.com/progrium/macdriver/core"
)

func main() {
	ud := core.NSUserDefaults_StandardUserDefaults()
	ud.SetURLForKey(core.URL("https://github.com/progrium/macdriver"), core.String("macdriver"))
	u := ud.URLForKey(core.String("macdriver"))
	fmt.Println("looked up 'macdriver' key in NSUserDefaults and got:", u)
}
