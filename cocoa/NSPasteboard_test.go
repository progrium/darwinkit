package cocoa_test

import (
	"fmt"
	"github.com/progrium/macdriver/cocoa"
	"io/ioutil"
	"log"
)

func ExampleNSPasteboard() {
	// get general pasteboard
	gp := cocoa.NSPasteboard_GeneralPasteboard()

	// clear pasteboard and write to pasteboard
	gp.ClearContents()
	gp.SetStringForType("test-content", cocoa.NSPasteboardTypeString)

	// read data from pasteboard
	fmt.Println(gp.StringForType(cocoa.NSPasteboardTypeString))
	// Output: test-content
}

func ExampleNSPasteboard_DataForType() {
	// get general pasteboard
	gp := cocoa.NSPasteboard_GeneralPasteboard()
	// check which is the best available type (list all the types you can to handle here)
	pbType := gp.AvailableTypeFromArray([]cocoa.NSPasteboardType{cocoa.NSPasteboardTypePNG, cocoa.NSPasteboardTypeTIFF})
	if pbType == "" {
		// check which types are in the pasteboard right now
		log.Printf("no matching type found. only found: %s \n", gp.Types())
		return
	}
	// fetch data
	data := gp.DataForType(pbType)
	var bs []byte
	// turn into go bytes
	bs = data.Bytes()

	f, err := ioutil.TempFile("", "out-img")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = f.Write(bs)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("png written to file: %s\n", f.Name())
}
