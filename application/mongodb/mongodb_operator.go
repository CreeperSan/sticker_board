package ApplicationDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	StickerBoard "sticker_board/application/const"
	SharedPreferences "sticker_board/lib/shared_preferences"
	"time"
)

const CollectionAccount = "account"
const CollectionStickerTag = "sticker_tag"
const CollectionStickerCategory = "sticker_category"
const CollectionAccountToken = "account_token"

var mongodbDatabaseDatabaseName = ""
var mongodbDatabaseAddress = ""
var mongodbDatabasePort = ""
var MongoClient *mongo.Client
var MongoDB *mongo.Database
func ConnectDB() (mongoClient *mongo.Client, mongoDB *mongo.Database) {
	if MongoClient != nil && MongoDB != nil {
		return MongoClient, MongoDB
	}

	if mongodbDatabaseDatabaseName == "" {
		mongodbDatabaseDatabaseName = SharedPreferences.GetString(StickerBoard.SPMongoDBDatabaseName, mongodbDatabaseDatabaseName)
	}
	if mongodbDatabaseAddress == "" {
		mongodbDatabaseAddress = SharedPreferences.GetString(StickerBoard.SPMongoDBDatabaseAddress, mongodbDatabaseAddress)
	}
	if mongodbDatabasePort == "" {
		mongodbDatabasePort = SharedPreferences.GetString(StickerBoard.SPMongoDBDatabasePort, mongodbDatabasePort)
	}

	//uri := "mongodb+srv://"+mongodbDatabaseAddress+":"+mongodbDatabasePort
	uri := "mongodb://"+ mongodbDatabaseAddress +":"+ mongodbDatabasePort
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	database := client.Database(mongodbDatabaseDatabaseName)

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")

	MongoClient = client
	MongoDB = database

	return MongoClient, MongoDB
}



