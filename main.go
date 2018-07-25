package main

import (
	"fmt"

	"github.com/jigtools/tail/drivers/elastic"

	"github.com/integrii/flaggy"
)

// make a variable for the version which will be set at build time
var version = "development build"
var connectionString = "http://localhost:9200"

func init() {
	// Set your program's name and description, if you want to.
	// This shows when you run help
	flaggy.SetName("Tail")
	flaggy.SetDescription("Tail, for more than just files")

	// you can disable various things by changing bools on the default parser (or your own parser if you have created one)
	flaggy.DefaultParser.ShowHelpOnUnexpected = true

	// you can set a help prepend or append on the default parser
	flaggy.DefaultParser.AdditionalHelpPrepend = "http://github.com/jigtools/tail"

	// set the version and parse
	flaggy.SetVersion(version)

	// Global flags
	flaggy.String(&connectionString, "c", "connect", fmt.Sprintf("Elasticsearch connection string, defaults to %s", connectionString))

	// ls - Show Indexes
	showIndexes := flaggy.NewSubcommand("ls")
	flaggy.AttachSubcommand(showIndexes, 1)

	flaggy.Parse()

	if showIndexes.Used {
		elastic.List(connectionString)
	} else {
		elastic.Connect(connectionString)
	}
}

func main() {
	println("Done.")
}
