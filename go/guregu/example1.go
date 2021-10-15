package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

const (
	tableName   = "thing"
	endpointUrl = "http://localhost:4566"
	region      = "us-east-1"
)

func main() {
	DropTable()
	CreateTable()
	UpdateTable()
	//DropTable()
}

type Thing struct {
	Id     int       `json:"id" dynamo:"id,hash" index:"foo-id-index,range"`
	Search string    `json:"search,omitempty" dynamo:"search,range" index:"idType-search-index,range"`
	IdType string    `json:"idType" dynamo:"idType,omitempty" index:"idType-search-index,hash"`
	Time   time.Time `json:"time" dynamo:"time,omiyempty"`
	Foo    int       `json:"foo,omitempty" dynamo:"foo,omitempty" index:"foo-id-index,hash"`
	Bar    string    `json:"bar,omitempty" dynamo:"bar,omitempty"`
}

func CreateTable() {
	const alreadyCreatedErrorStr = "ResourceInUseException: Table already created"
	dynamoClient := getDynamoClient()
	err := dynamoClient.CreateTable(tableName, Thing{}).Run()
	if err != nil && err.Error() != alreadyCreatedErrorStr {
		panic(err)
	}
	fmt.Println("Table created successfully!")
}

func DropTable() {
	const doesNotExistErrorStr = "ResourceNotFoundException: Cannot do operations on a non-existent table"
	db := getDynamoClient()

	if err := db.Table(tableName).DeleteTable().Run(); err != nil && err.Error() != doesNotExistErrorStr {
		panic(err)
	}

	fmt.Println("Table dropped successfully!")
}

const (
	IdType1 = "A"
	IdType2 = "B"
)

func UpdateTable() {
	db := getDynamoClient()
	table := db.Table(tableName)

	//
	// Put things
	//
	var thing Thing
	rand.Seed(time.Now().UnixNano())
	idTypes := []string{IdType1, IdType2}
	var count int
	for _, idType := range idTypes {
		for i := 0; i < 5; i++ {
			thing = Thing{
				IdType: idType,
				Id:     count,
				Search: strconv.Itoa(count),
				Time:   time.Now(),
				Foo:    count,
				Bar:    strconv.Itoa(count),
			}
			err := table.Put(thing).Run()
			if err != nil {
				log.Fatal(err)
				fmt.Println("Failure")
				os.Exit(1)
			}
			count++
		}
	}

	// Get the last item.
	var result Thing
	table.
		Get("id", thing.Id).
		Range("search", dynamo.Equal, thing.Search).
		One(&result)
	fmt.Println("Last item inserted")
	fmt.Printf("\t%v\n", result)

	fmt.Println("All items:")
	{
		var results []Thing
		if err := table.
			Scan().
			All(&results); err != nil {
			log.Fatal(err)
		}
		for i, item := range results {
			fmt.Printf("%d: \t%v\n", i, item)
		}
	}

	const minCount = 3
	fmt.Printf("Items with foo >= %d:\n", minCount)
	{
		var filtered []Thing
		if err := table.
			Scan().
			Filter("'foo' >= ?", minCount).
			All(&filtered); err != nil {
			log.Fatal(err)
		}
		for i, item := range filtered {
			fmt.Printf("%d: \t%v\n", i, item)
		}
	}

	fmt.Printf("Indexed things:\n")
	{
		var indexed []Thing
		if err := table.
			Get("idType", IdType2).
			Index("idType-search-index").
			All(&indexed); err != nil {
			log.Fatal(err)
		}
		for i, item := range indexed {
			fmt.Printf("%d: \t%v\n", i, item)
		}
	}

	fmt.Printf("Found ONE item by (pk,sk) combination:\n")
	{
		var thing Thing
		if err := table.
			Get("id", 7).
			Range("search", dynamo.Equal, "7").
			//Index("idType-search-index").
			One(&thing); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\t%v\n", thing)
	}

	fmt.Printf("Found ALL items by (pk,sk) combination:\n")
	{
		var things []Thing
		if err := table.
			Get("id", 7).
			Range("search", dynamo.Equal, "7").
			//Index("idType-search-index").
			All(&things); err != nil {
			log.Fatal(err)
		}
		for i, item := range things {
			fmt.Printf("%d: \t%v\n", i, item)
		}
	}

	fmt.Printf("Found ALL items by (pk,sk) combination using secondary index:\n")
	{
		var things []Thing
		if err := table.
			Get("idType", IdType2).
			Range("search", dynamo.GreaterOrEqual, "7").
			Index("idType-search-index").
			All(&things); err != nil {
			log.Fatal(err)
		}
		for i, item := range things {
			fmt.Printf("%d: \t%v\n", i, item)
		}
	}

	fmt.Printf("Found ALL items by (pk,sk) combination using secondary index and use FILTER:\n")
	{
		var things []Thing
		if err := table.
			Get("idType", IdType2).
			Range("search", dynamo.LessOrEqual, "7").
			Index("idType-search-index").
			Filter("'foo' >= ?", 3). // That that IdType starts with 5, not already at 3.
			All(&things); err != nil {
			log.Fatal(err)
		}
		for i, item := range things {
			fmt.Printf("%d: \t%v\n", i, item)
		}
	}
}

func getDynamoClient() *dynamo.DB {
	config := aws.Config{
		Endpoint: aws.String(endpointUrl),
		Region:   aws.String(region),
	}
	session := session.Must(session.NewSession())
	return dynamo.New(session, &config)
}
