package StickerDatabase

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	Application "sticker_board/application/database"
	LogService "sticker_board/lib/log_service"
	"sticker_board/sticker/v1/model"
)

func migrateDatabaseWithInstance(db *gorm.DB, databaseModel interface{}, databaseName string){
	err := db.AutoMigrate(databaseModel)
	if err != nil {
		LogService.Error("Error occurred while auto migrating ", databaseName, " database.")
		panic(err)
	}
}

func migrateMongoDBWithInstance(mongoClient *mongo.Client, mongoDB *mongo.Database, mongoCtx context.Context, collectionName string){
	mongoDB.Collection(collectionName)
}

func Initialize(){
	db := Application.GetDB()

	migrateDatabaseWithInstance(db, StickerDatabase.CategoryModel{}, "CategoryModel")
	migrateDatabaseWithInstance(db, StickerDatabase.TagModel{}, "TagModel")
	migrateDatabaseWithInstance(db, StickerDatabase.StickerBasicModel{}, "StickerBasicModel")
	migrateDatabaseWithInstance(db, StickerDatabase.StickerPlainTextModel{}, "StickerPlainTextModel")
	migrateDatabaseWithInstance(db, StickerDatabase.StickerTagModel{}, "StickerTagModel")
	migrateDatabaseWithInstance(db, StickerDatabase.StickerCategoryModel{}, "StickerCategoryModel")

	mongoDB, mongoClient, mongoCtx := Application.GetMongoDB()
	migrateMongoDBWithInstance(mongoDB, mongoClient ,mongoCtx, "Category")
	migrateMongoDBWithInstance(mongoDB, mongoClient ,mongoCtx, "Tag")
	migrateMongoDBWithInstance(mongoDB, mongoClient ,mongoCtx, "Sticker")
}

