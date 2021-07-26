package StickerDatabase

import (
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

func Initialize(){
	db := Application.GetDB()

	migrateDatabaseWithInstance(db, Sticker.CategoryModel{}, "CategoryModel")
	migrateDatabaseWithInstance(db, Sticker.TagModel{}, "TagModel")
	migrateDatabaseWithInstance(db, Sticker.StickerBasicModel{}, "StickerBasicModel")
	migrateDatabaseWithInstance(db, Sticker.StickerPlainTextModel{}, "StickerPlainTextModel")
	migrateDatabaseWithInstance(db, Sticker.StickerTagModel{}, "StickerTagModel")
	migrateDatabaseWithInstance(db, Sticker.StickerCategoryModel{}, "StickerCategoryModel")

	//CreateStickerPlainText(2, 0, 0, "Title", "Content")
}

