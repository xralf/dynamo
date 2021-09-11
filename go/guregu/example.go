package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

const (
	tableName   = "widget"
	endpointUrl = "http://localhost:4566"
	region      = "us-east-1"
)

func main() {
	DropTable()
	CreateTable()
	UpdateTable()
	//DropTable()
}

// Use struct tags much like the standard JSON library,
// you can embed anonymous structs too!
// type Widget struct {
// 	UserID int       // Hash key, a.k.a. partition key
// 	Time   time.Time // Range key, a.k.a. sort key
// 	Msg       string              `dynamo:"Message"`    // Change name in the database
// 	Count     int                 `dynamo:",omitempty"` // Omits if zero value
// 	Children  []Widget            // Lists
// 	Friends   []string            `dynamo:",set"` // Sets
// 	Set       map[string]struct{} `dynamo:",set"` // Map sets, too!
// 	SecretKey string              `dynamo:"-"`    // Ignored
// }

type Widget struct {
	Id         int                 `json:"id" dynamo:"id,hash" index:"time-id-index,range"`                                  // Hash key, a.k.a. partition key
	Time       time.Time           `json:"time" dynamo:"time,range" index:"time-id-index,hash" index:"msg-time-index,range"` // Range key, a.k.a. sort key
	Msg        string              `json:"msg,omitempty" dynamo:"msg,omitempty" index:"msg-time-index,hash"`                 // Change name in the database
	Count      int                 `json:"count,omitempty" dynamo:"count,omitempty"`                                         // Omits if zero value
	Children   []Widget            `json:"children,omitempty" dynamo:"children,omitempty"`                                   // Lists
	Friends    []string            `json:"friends,omitempty" dynamo:"friends,set"`                                           // Sets
	Set        map[string]struct{} `json:"set,omitempty" dynamo:"set,set"`                                                   // Map sets, too!
	SecretKey  string              `json:"secretKey,omitempty" dynamo:"-"`                                                   // Ignored
	SecretKey2 string              `json:"secretKey2,omitempty" dynamo:"secretKey2"`                                         // Ignored
}

func CreateTable() {
	const alreadyCreatedErrorStr = "ResourceInUseException: Table already created"
	dynamoClient := getDynamoClient()
	err := dynamoClient.CreateTable(tableName, Widget{}).Run()
	//err := dynamoClient.CreateTable("my-crane", CraneRecord{}).Run()
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

func UpdateTable() {
	db := getDynamoClient()
	table := db.Table(tableName)

	// put item
	var w Widget
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		children := []Widget{
			Widget{Id: 101, Time: time.Now(), Msg: "aaa", Count: rand.Intn(10)},
			Widget{Id: 102, Time: time.Now(), Msg: "bbb", Count: rand.Intn(10), SecretKey2: "foo"},
		}
		friends := []string{"Alice", "Bob", "Cindy"}
		secretKey := "keep_me_secret"
		w = Widget{
			Id:        613,
			Time:      time.Now(),
			Msg:       "hello",
			Count:     rand.Intn(10),
			Children:  children,
			Friends:   friends,
			SecretKey: secretKey,
		}
		err := table.Put(w).Run()
		if err != nil {
			log.Fatal(err)
			fmt.Println("Failure")
			os.Exit(1)
		}
	}

	// Get the last item.
	var result Widget
	table.Get("id", w.Id).Range("time", dynamo.Equal, w.Time).One(&result)
	fmt.Println("Last item inserted")
	fmt.Printf("\t%v\n", result)

	// Get all items.
	var results []Widget
	if err := table.Scan().All(&results); err != nil {
		log.Fatal(err)
	}
	fmt.Println("All items:")
	for i, item := range results {
		fmt.Printf("%d: \t%v\n", i, item)
	}

	// Search for items.
	// Use placeholders in filter expressions.
	var filtered []Widget
	const minCount = 6
	if err := table.Scan().Filter("'count' >= ?", minCount).All(&filtered); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Items with count >= %d:\n", minCount)
	for i, item := range filtered {
		fmt.Printf("%d: \t%v\n", i, item)
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
