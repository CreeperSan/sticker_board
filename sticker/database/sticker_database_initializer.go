package StickerDatabase

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	Application "sticker_board/application/database"
	LogService "sticker_board/lib/log_service"
	Sticker "sticker_board/sticker/model"
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

	migrateDatabaseWithInstance(db, Sticker.CategoryModel{}, "CategoryModel")
	migrateDatabaseWithInstance(db, Sticker.TagModel{}, "TagModel")
	migrateDatabaseWithInstance(db, Sticker.StickerBasicModel{}, "StickerBasicModel")
	migrateDatabaseWithInstance(db, Sticker.StickerPlainTextModel{}, "StickerPlainTextModel")
	migrateDatabaseWithInstance(db, Sticker.StickerTagModel{}, "StickerTagModel")
	migrateDatabaseWithInstance(db, Sticker.StickerCategoryModel{}, "StickerCategoryModel")

	mongoDB, mongoClient, mongoCtx := Application.GetMongoDB()
	migrateMongoDBWithInstance(mongoDB, mongoClient ,mongoCtx, "Category")
	migrateMongoDBWithInstance(mongoDB, mongoClient ,mongoCtx, "Tag")
	migrateMongoDBWithInstance(mongoDB, mongoClient ,mongoCtx, "Sticker")
}

