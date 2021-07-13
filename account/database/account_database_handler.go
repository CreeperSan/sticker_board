package StickerBoardAccount

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	StickerBoardAccount "sticker_board/account/database/model"
	ActionResponse "sticker_board/account/response"
	StickerBoard "sticker_board/application/const"
	Formatter "sticker_board/lib/formatter"
	LogService "sticker_board/lib/log_service"
	SharedPreferences "sticker_board/lib/shared_preferences"
	"strings"
)

var databaseName string = ""
var databaseUsername string = ""
var databasePassword string = ""
var databasePort int = 0
var databaseAddress string = ""
func getDB() *gorm.DB {
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

func Initialize()  {
	// Connect to MySQL database
	db := getDB()

	// Auto migrate -> Account Table
	err := db.AutoMigrate(&StickerBoardAccount.AccountModel{})
	if err != nil {
		LogService.Error("Error occurred while auto migrating table : AccountModel.")
		panic(err)
	}

	// Auto migrate -> Account Auth Table
	err = db.AutoMigrate(&StickerBoardAccount.AccountTokenModel{})
	if err != nil {
		LogService.Error("Error occurred while auto migrating table : AccountTokenModel.")
		panic(err)
	}
}

func RegisterAccount(
	account string,
	password string,
	userName string,
	email string,
) ActionResponse.ActionResponse {
	account = strings.TrimSpace(account)
	password = strings.TrimSpace(password)
	userName = strings.TrimSpace(userName)
	email = strings.TrimSpace(email)

	if !Formatter.CheckStringWithLength(account, 6, 20) {
		LogService.Warming("Fail to register account, account not valid. account=",account," password=",password," userName=",userName," email=",email,".")
		return ActionResponse.CreateActionFailResponse("Account field length error.")
	}

	if !Formatter.CheckStringWithLength(password, 8, 30) {
		LogService.Warming("Fail to register account, password not valid. account=",account," password=",password," userName=",userName," email=",email,".")
		return ActionResponse.CreateActionFailResponse("Password field length error.")
	}

	if !Formatter.CheckStringWithLength(userName, 2, 20) {
		LogService.Warming("Fail to register account, username not valid. account=",account," password=",password," userName=",userName," email=",email,".")
		return ActionResponse.CreateActionFailResponse("User name field length error.")
	}

	if !Formatter.CheckStringWithLength(email, 6, 30) {
		LogService.Warming("Fail to register account, email not valid. account=",account," password=",password," userName=",userName," email=",email,".")
		return ActionResponse.CreateActionFailResponse("Email field length error.")
	}

	if !Formatter.CheckStringIsValidEmail(email) {
		LogService.Warming("Fail to register account, email. account=",account," password=",password," userName=",userName," email=",email,".")
		return ActionResponse.CreateActionFailResponse("Email address is invalid.")
	}

	// encrypt should run after length check
	password = Formatter.FormatPassword(password)

	var queryCount int64 = 0
	// check account weather already been register
	err := getDB().Model(&StickerBoardAccount.AccountModel{}).Where(StickerBoardAccount.ColumnAccountModelAccount+" = ?", account).Count(&queryCount)
	if queryCount > 0 {
		LogService.Warming("Fail to register account, account already exist. account=",account," password=",password," userName=",userName," email=",email,".")
		return ActionResponse.CreateActionFailResponse("Account already register, please login directly")
	}

	// check account weather already been register
	err = getDB().Model(&StickerBoardAccount.AccountModel{}).Where(StickerBoardAccount.ColumnAccountModelEmail+" = ?", email).Count(&queryCount)
	if queryCount > 0 {
		LogService.Warming("Fail to register account, email address already in used. account=",account," password=",password," userName=",userName," email=",email,".")
		return ActionResponse.CreateActionFailResponse("Email already registered")
	}

	// insert user info to database
	var accountModel = StickerBoardAccount.AccountModel{}
	accountModel.Account = account
	accountModel.Password = password
	accountModel.UserName = userName
	accountModel.Email = email

	result := getDB().Create(&accountModel)
	if result.Error != nil {
		LogService.Failure("Fail to register account, insert into database error. account =",account," password =",password," userName =",userName," email =",email,".")
		LogService.Error(err)
		return ActionResponse.CreateActionFailResponse("Service internal error")
	}

	LogService.Success("Account has been successfully registered. account =",account, " username =", userName, " email =", email)

	response := ActionResponse.CreateActionFailResponse("todo")
	return response
}

func LoginAccount(
	account string,
	password string,
	platform int8,
	brand string,
	deviceName string,
	machineCode string,
) ActionResponse.ActionResponse {
	account = strings.TrimSpace(account)
	password = strings.TrimSpace(password)
	password = Formatter.FormatPassword(password)

	// check whether the account and database correct
	var queryAccountModel = StickerBoardAccount.AccountModel{}
	queryAccountModel.ID = 0
	db := getDB()
	db.Where(StickerBoardAccount.ColumnAccountModelAccount+" = ? and "+StickerBoardAccount.ColumnAccountModelPassword+" = ?", account, password).Limit(1).Find(&queryAccountModel)

	// if the password or account does not correct
	if queryAccountModel.ID == 0 {
		LogService.Warming("Login failed. account =", account, ",password =", password)
		return ActionResponse.CreateActionFailResponse("Account or password error.")
	}

	// check whether the token count is out of range (maximum 5 tokens each account)
	var queryTokenModelList []StickerBoardAccount.AccountTokenModel
	queryTokenModelListResult := db.Where(StickerBoardAccount.ColumnAccountTokenModelAccountID+" = ?", queryAccountModel.ID).Order(StickerBoardAccount.ColumnAccountTokenModelUpdateTime+" desc").Find(&queryTokenModelList)
	if queryTokenModelListResult.RowsAffected >= 5 {
		// remove the oldest token

	}

	// create and insert new token
	tokenUUID, errUUID := uuid.NewRandom()
	if errUUID != nil {
		LogService.Warming("Login failed, can not generate token. account =", account, ",password =", password)
		return ActionResponse.CreateActionFailResponse("Can not generate token.")
	}
	var insertAccountTokenModel = StickerBoardAccount.AccountTokenModel{
		Token:                 fmt.Sprintf("%s", tokenUUID),
		AccountID:             queryAccountModel.ID,
		Platform:              platform,
		Brand:                 brand,
		DeviceName:            deviceName,
		MachineCode:           machineCode,
		ExpireTimeMilliSecond: 1000 * 60 * 60 * 24 * 7,
	}
	db.Create(&insertAccountTokenModel)

	return ActionResponse.CreateActionSuccessResponse()
}

