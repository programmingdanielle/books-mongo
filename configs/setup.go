package configs

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Port = ":3000"

func ConnectDB() *mongo.Client {
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(config))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(banner() + "\nRunning on port " + Port)

	return client
}

// // Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Publishing").Collection(collectionName)
	return collection
}

func banner() string {
	b, err := ioutil.ReadFile("ascii.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func getConfig() (string, error) {
	plan, _ := ioutil.ReadFile("config.json")
	var data interface{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		return "", err
	}

	var cred string

	m, _ := data.(map[string]interface{})
	for _, v := range m {
		StrVal := fmt.Sprint(v)
		cred = StrVal
	}

	return cred, nil
}
