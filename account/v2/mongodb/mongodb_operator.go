package AccountV2

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	StickerBoard "sticker_board/application/const"
	SharedPreferences "sticker_board/lib/shared_preferences"
	"time"
)

var mongodbDatabaseDatabaseName = ""
var mongodbDatabaseAddress = ""
var mongodbDatabasePort = ""
func GetDB() (mongoClient *mongo.Client, mongoDB *mongo.Database, mongoCtx context.Context) {
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
	uri := "mongodb://"+mongodbDatabaseAddress+":"+mongodbDatabasePort
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	database := client.Database(mongodbDatabaseDatabaseName)

	// Ping the primary
	//if err := client.Ping(ctx, readpref.Primary()); err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Successfully connected and pinged.")

	return client, database, ctx
}



