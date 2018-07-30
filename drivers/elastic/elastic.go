package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	elastic "gopkg.in/olivere/elastic.v5"
)

// Connect to the elastic search
func Connect(connectionString string) *elastic.Client {
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
	return client
}

// List Indexes
func List(connectionString string) {
	client := Connect(connectionString)
	if client != nil {
		names, err := client.IndexNames()
		if err != nil {
			// Handle error
			panic(err)
		}
		for _, name := range names {
			fmt.Printf("%s\n", name)
		}
	}
}

// Tail from indexes
func Tail(connectionString, index, format, timestampField string) {
	fmt.Printf("Connecting to %s: index %s, format: %s\n", connectionString, index, format)
	ctx := context.Background()
	initialLogCount := 20 // get 20 log entries initially

	client := Connect(connectionString)
	if client != nil {
		searchResult, err := client.Search().
			Index(index). // search in index "twitter"
			//Query(termQuery).   // specify the query
			Sort(timestampField, false).
			From(0).Size(initialLogCount). // take documents 0-9
			Pretty(true).                  // pretty print request and response JSON
			Do(ctx)                        // execute
		if err != nil {
			// TODO: if not found, and no wildcard, add it - ie `--index infra` becomes `--index infra*`
			// Handle error
			panic(err)
		}
		// Iterate through results
		//for _, hit := range searchResult.Hits.Hits {
		for i := len(searchResult.Hits.Hits) - 1; i >= 0; i-- {
			fmt.Println(formatHit(format, searchResult.Hits.Hits[i]))
		}
	}
}

func formatHit(format string, hit *elastic.SearchHit) string {
	if format == "*" {
		return string(*hit.Source)
	}
	// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
	var entry map[string]interface{}
	err := json.Unmarshal(*hit.Source, &entry)
	if err != nil {
		return err.Error()
	}
	// TODO: expand to more than one key
	keys := strings.Split(format, " ")
	for i, k := range keys {
		value := entry[k]
		if value != nil {
			keys[i] = fmt.Sprint(value)
		}
	}
	return strings.Join(keys, " ")
}
