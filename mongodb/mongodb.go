package mongodb

import (
	"context"
	"os"
	"reflect"
	"time"

	"github.com/anyshare/anyshare-admin-api/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() {
	if client == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		c, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
		if err != nil {
			panic(err)
		}
		client = c
	}
}

func Disconnect() {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := client.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
		client = nil
	}
}

func GetCollection(v interface{}) *mongo.Collection {
	return client.Database(os.Getenv("MONGO_DB")).Collection(getCollectionName(v))
}

func getCollectionName(v interface{}) string {
	var col string
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Pointer || t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		col = helpers.ToSnakeCase(t.Elem().Name()) + "s"
	} else if t.Name() == "string" {
		col = v.(string)
	} else {
		col = helpers.ToSnakeCase(t.Name()) + "s"
	}
	return col
}
