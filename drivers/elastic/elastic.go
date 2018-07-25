package elastic

import (
	"context"
	"fmt"
	"time"

	elastic "gopkg.in/olivere/elastic.v5"
)

// Connect to the elastic search
func Connect(connectionString string) {
	// Starting with elastic.v5, you must pass a context to execute each service
	ctx := context.Background()

	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course you can configure your client to connect
	// to other hosts and configure it in various other ways.
	defaultOptions := []elastic.ClientOptionFunc{
		elastic.SetURL(connectionString),
		elastic.SetSniff(false),
		elastic.SetHealthcheckTimeoutStartup(10 * time.Second),
		elastic.SetHealthcheckTimeout(2 * time.Second),
	}
	client, err := elastic.NewClient(defaultOptions...)
	if err != nil {
		// Handle error
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(connectionString).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion(connectionString)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	names, err := client.IndexNames()
	if err != nil {
		// Handle error
		panic(err)
	}
	for _, name := range names {
		fmt.Printf("%s\n", name)
	}
}
