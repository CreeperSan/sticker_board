package Account

import (
	"fmt"
	"github.com/google/uuid"
	"sticker_board/account/v1/database/model"
	"sticker_board/account/v1/database/response"
	ActionResponse "sticker_board/account/v1/response"
	Application "sticker_board/application/database"
	Formatter "sticker_board/lib/formatter"
	LogService "sticker_board/lib/log_service"
	"strings"
	"time"
)

// TODO : when using the same params to login should not only generate a new token but also delete the old token

const _MAX_TOKEN_PER_ACCOUNT = 5

func Initialize()  {
	// Connect to MySQL database
	db := Application.GetDB()

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

	// remove all cached tokens
	//db.Where("1 = 1").Delete(&Account.AccountTokenModel{})
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
	err := Application.GetDB().Model(&StickerBoardAccount.AccountModel{}).Where(StickerBoardAccount.ColumnAccountModelAccount+" = ?", account).Count(&queryCount)
	if queryCount > 0 {
		LogService.Warming("Fail to register account, account already exist. account=",account," password=",password," userName=",userName," email=",email,".")
		return ActionResponse.CreateActionFailResponse("Account already register, please login directly")
	}

	// check account weather already been register
	err = Application.GetDB().Model(&StickerBoardAccount.AccountModel{}).Where(StickerBoardAccount.ColumnAccountModelEmail+" = ?", email).Count(&queryCount)
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

	result := Application.GetDB().Create(&accountModel)
	if result.Error != nil {
		LogService.Failure("Fail to register account, insert into database error. account =",account," password =",password," userName =",userName," email =",email,".")
		LogService.Error(err)
		return ActionResponse.CreateActionFailResponse("Service internal error")
	}

	LogService.Success("Account has been successfully registered. account =",account, " username =", userName, " email =", email)

	response := ActionResponse.CreateActionSuccessResponse()
	return response
}

func LoginAccount(
	account string,
	password string,
	platform int8,
	brand string,
	deviceName string,
	machineCode string,
) Account.LoginDatabaseResponse {
	account = strings.TrimSpace(account)
	password = strings.TrimSpace(password)
	password = Formatter.FormatPassword(password)

	// check whether the account and database correct
	var queryAccountModel = StickerBoardAccount.AccountModel{}
	queryAccountModel.ID = 0
	db := Application.GetDB()
	db.Where(StickerBoardAccount.ColumnAccountModelAccount+" = ? and "+StickerBoardAccount.ColumnAccountModelPassword+" = ?", account, password).Limit(1).Find(&queryAccountModel)

	// if the password or account does not correct
	if queryAccountModel.ID == 0 {
		LogService.Warming("Login failed. account =", account, ",password =", password)
		return Account.LoginDatabaseResponse{
			Code: 400,
			Message: "Account or password not correct.",
		}
	}

	// check whether the token count is out of range (maximum 5 tokens each account)
	var queryTokenModelList []StickerBoardAccount.AccountTokenModel
	queryTokenModelListResult := db.Where(StickerBoardAccount.ColumnAccountTokenModelAccountID+" = ?", queryAccountModel.ID).Order(StickerBoardAccount.ColumnAccountTokenModelUpdateTime +" desc").Find(&queryTokenModelList)
	if queryTokenModelListResult.RowsAffected >= _MAX_TOKEN_PER_ACCOUNT {
		// remove the oldest token
		db.Delete(queryTokenModelList[_MAX_TOKEN_PER_ACCOUNT- 1])
	}

	// create and insert new token
	tokenUUID, errUUID := uuid.NewRandom()
	if errUUID != nil {
		LogService.Warming("Login failed, can not generate token. account =", account, ",password =", password)
		return Account.LoginDatabaseResponse{
			Code: 500,
			Message: "Server internal error.",
		}
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

	return Account.LoginDatabaseResponse{
		Code: 200,
		Message: "Login success",
		Token: insertAccountTokenModel.Token,
		EffectiveTime: insertAccountTokenModel.ExpireTimeMilliSecond,
		UID: insertAccountTokenModel.AccountID,
	}
}

func AuthToken(accountID uint, token string, platform int, brand string, deviceName string, machineCode string) Account.AuthDatabaseResponse {
	db := Application.GetDB()

	var queryList []StickerBoardAccount.AccountTokenModel
	queryListResult := db.Where(StickerBoardAccount.ColumnAccountTokenModelToken+" = ? and "+
		StickerBoardAccount.ColumnAccountTokenModelPlatform+ " = ? and " +
		StickerBoardAccount.ColumnAccountTokenModelBrand+ " = ? and " +
		StickerBoardAccount.ColumnAccountTokenModelDeviceName+ " = ? and " +
		StickerBoardAccount.ColumnAccountTokenModelMachineCode+ " = ?",
		token, platform, brand, deviceName, machineCode,
	).Order(StickerBoardAccount.ColumnAccountTokenModelUpdateTime +" desc").Limit(1).Find(&queryList)

	// check whether the token exists
	if queryListResult.RowsAffected >= 1 {
		var tokenModel = queryList[0]
		var currentTimestamp = time.Now().UnixNano() / 1000_000 // convert to millisecond
		// check whether the token is expired
		if currentTimestamp - tokenModel.UpdateTime > tokenModel.ExpireTimeMilliSecond {
			// token is expired
			db.Delete(tokenModel)
			LogService.Success("Auth account failed, token expired. currentTimestamp =", currentTimestamp, " updateTimestamp =", tokenModel.UpdateTime, " expiredTimeMilliSecond =", tokenModel.ExpireTimeMilliSecond)
			return Account.AuthDatabaseResponse{
				Code: 400,
				Message: "Token was expired, please login.",
			}
		} else {
			// token is not expired
			// then check the account is the same as database's account id
			if accountID != tokenModel.AccountID {
				return Account.AuthDatabaseResponse{
					Code: 400,
					Message: "Token was expired, please login.",
				}
			}
			// auth pass
			tokenModel.UpdateTime = currentTimestamp
			db.Save(tokenModel)
			LogService.Success("Auth account success. token =", token)
			return Account.AuthDatabaseResponse{
				Code: 200,
				Message: "Auth succeed",
				UpdateTime: tokenModel.UpdateTime,
				ExpireTimeMilliSecond: tokenModel.ExpireTimeMilliSecond,
				AccountID: tokenModel.AccountID,
			}
		}
	}
	LogService.Success("Auth account success. token not exist. token =", token)
	return Account.AuthDatabaseResponse{
		Code: 400,
		Message: "Token was expired, please login.",
	}
}

func IsAccountExist(accountID uint) bool {
	db := Application.GetDB()

	var queryCount int64 = 0

	db.Model(&StickerBoardAccount.AccountModel{}).Where(StickerBoardAccount.ColumnAccountModelID+" = ?", accountID).Count(&queryCount)

	return queryCount > 0
}
