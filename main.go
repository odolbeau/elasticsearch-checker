package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/olivere/elastic.v2"
	"log"
)

func main() {
	fmt.Println("Let's start!")

	client, err := elastic.NewClient()
	if err != nil {
		log.Fatalf("Cannot connect to elasticsearch: %s", err)
	}

	termQuery := elastic.NewTermQuery("trip_from_address", "paris")
	searchResult, err := client.Search().
		Index("trips").   // search in index "twitter"
		Query(termQuery). // specify the query
		From(0).Size(10). // take documents 0-9
		Pretty(true).     // pretty print request and response JSON
		Do()              // execute
	if err != nil {
		log.Fatal(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	db, err := sql.Open("mysql", "root:@/covoiturage_dev")
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var id int64
	row := db.QueryRow("SELECT id FROM TripOffer WHERE id = ?", "1000000")
	if err := row.Scan(&id); err != nil {
		log.Fatalf("Unable to run MySQL query: %s", err)
	}
	fmt.Printf("%d\n", id)
}
