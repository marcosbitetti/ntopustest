package database

import (
	"context"
	"log"
	"os"
	"reflect"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// "go.mongodb.org/mongo-driver/bson"
// singleten derivated from https://medium.com/golang-issue/how-singleton-pattern-works-with-golang-2fdd61cd5a7f

var client *mongo.Client
var instance *mongo.Database
var once sync.Once

func DB() *mongo.Database {

	once.Do(func() {
		var uri string = "mongodb://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@localhost:27017/" // + os.Getenv("DB_NAME")
		_client, err := mongo.NewClient(options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = _client.Connect(ctx)
		if err != nil {
			panic(err)
		}

		// defer client.Disconnect(ctx)
		client = _client
		instance = client.Database(os.Getenv("DB_NAME"))
	})

	return instance
}

func UserCollection() *mongo.Collection {
	db := DB()
	userCollection := db.Collection("healt")
	/*
		cur, err := userCollection.Find(context.Background(), bson.D{})
		if err != nil {
			log.Fatal(err)
		}

		defer cur.Close(context.Background())

		for cur.Next(context.Background()) {
			var result user.User = user.User{}
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(result.Nome())
		}*/
	return userCollection
}

func Find(param bson.M, u interface{}, coll *mongo.Collection) []interface{} {
	opts := options.Find()
	opts.SetSort(bson.D{{"nome", 1}})

	list := make([]interface{}, 0)
	cur, err := coll.Find(context.Background(), param, opts)
	if err != nil {
		log.Fatal(err)
	}

	objType := reflect.TypeOf(u).Elem()
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		result := reflect.New(objType).Interface()
		err := cur.Decode(result)
		if err != nil {
			return list
		}
		list = append(list, result)
	}

	if err := cur.Err(); err != nil {
		return list
	}
	return list
}

// geric find
func FindAll(u interface{}, coll *mongo.Collection) []interface{} {
	return Find(bson.M{}, u, coll)
	/*list := make([]interface{}, 0)
	cur, err := coll.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	objType := reflect.TypeOf(u).Elem()

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		result := reflect.New(objType).Interface()
		err := cur.Decode(result)
		if err != nil {
			return list
		}

		list = append(list, result)
	}
	if err := cur.Err(); err != nil {
		return list
	}
	return list*/
}

/*
 * generic insert
 */
func Insert(u interface{}, coll *mongo.Collection) bool {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	data, err := bson.Marshal(u)
	_, _err := coll.InsertOne(ctx, data)
	if _err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
