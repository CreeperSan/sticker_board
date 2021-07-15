package Application

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	StickerBoard "sticker_board/application/const"
	SharedPreferences "sticker_board/lib/shared_preferences"
)


var databaseName string = ""
var databaseUsername string = ""
var databasePassword string = ""
var databasePort int = 0
var databaseAddress string = ""
func GetDB() *gorm.DB {
	if databaseName == "" {
		databaseName = SharedPreferences.GetString(StickerBoard.SPMySQLDatabaseName, databaseName)
	}
	if databaseUsername == "" {
		databaseUsername = SharedPreferences.GetString(StickerBoard.SPMySQLDatabaseUserName, databaseUsername)
	}
	if databasePassword == "" {
		databasePassword = SharedPreferences.GetString(StickerBoard.SPMySQLDatabasePassword, databasePassword)
	}
	if databaseAddress == "" {
		databaseAddress = SharedPreferences.GetString(StickerBoard.SPMySQLDatabaseAddress, databaseAddress)
	}
	if databasePort == 0 {
		databasePort = SharedPreferences.GetInt(StickerBoard.SPMySQLDatabasePort, databasePort)
	}

	dsn := databaseUsername+":"+databasePassword+"@/"+databaseName+"?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
