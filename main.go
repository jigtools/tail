package main

import (
  "fmt"

	"github.com/integrii/flaggy"
)

// make a variable for the version which will be set at build time
var version = "development build"

func init() {
  // Set your program's name and description, if you want to.
  // This shows when you run help
  flaggy.SetName("Tail")
  flaggy.SetDescription("Tail, for more than just files")

  // you can disable various things by changing bools on the default parser (or your own parser if you have created one)
  flaggy.DefaultParser.ShowHelpOnUnexpected = false

  // you can set a help prepend or append on the default parser
  flaggy.DefaultParser.AdditionalHelpPrepend = "http://github.com/jigtools/tail"

  // set the version and parse
  flaggy.SetVersion(version)
  flaggy.Parse()
}

func main() {
	fmt.Printf("Hello, world.\n")
}