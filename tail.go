package main

import (
	"fmt"

	"github.com/jigtools/tail/drivers/elastic"

	"github.com/integrii/flaggy"
)

// make a variable for the version which will be set at build time
var version = "development build"
var connectionString = "http://localhost:9200"
var index = "*"
var format = "@timestamp log"
var timestampField = "@timestamp"

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
	flaggy.String(&index, "i", "index", fmt.Sprintf("Index name / filter, defaults to %s", index))
	flaggy.String(&format, "f", "format", fmt.Sprintf("Format output using space separated key names, defaults to %s (set to * to see all fields)", format))

	cmdPosition := 0

	cmdPosition = cmdPosition + 1
	// ls - Show Indexes
	showIndexes := flaggy.NewSubcommand("ls")
	flaggy.AttachSubcommand(showIndexes, cmdPosition)

	flaggy.Parse()

	if showIndexes.Used {
		// TODO: add index here as a way to limit the results
		elastic.List(connectionString)
		return
	}

	// Default
	elastic.Tail(connectionString, index, format, timestampField)
}

func main() {
	println("Done.")
}
