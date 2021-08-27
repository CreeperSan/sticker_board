package AccountV2

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	_ "sticker_board/account/manager"
	AccountModule "sticker_board/account/manager/model"
	AccountResponse "sticker_board/account/manager/response"
	AccountV2Model "sticker_board/account/v2/model"
	StickerBoard "sticker_board/application/const"
	"sticker_board/application/mongodb"
	Formatter "sticker_board/lib/formatter"
	"sticker_board/lib/log_service"
	"strings"
)

type AccountOperator struct { }


func (operator AccountOperator) Initialize(){
	LogService.Info("Initializing Account Module ...")
	AccountV2DB.ConnectDB()
	// Check the database connection
	if AccountV2DB.MongoClient ==nil || AccountV2DB.MongoDB == nil{
		LogService.Error("Initializing Account Module Failed! Can not connect to mongoDB.")
		os.Exit(StickerBoard.ExitCodeDatabaseCreateClientConnectionFailed)
	}
	if err:= AccountV2DB.MongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		LogService.Error("Initializing Account Module Failed! Can not ping mongoDB. err =", err)
		os.Exit(StickerBoard.ExitCodeDatabasePingFailed)
	}

	// Test method
	//operator.RegisterAccount("account_01", "Aa123456", "UserName_01", "text_01@mail.com")
	//operator.LoginAccount("account_01", "Aa123456",  1, "Test Brand", "Test Device", "123456789012345678")
	//operator.AuthToken("6126e612a9192b1b0c9628be", "df245984-f653-4abb-a58c-3e9edd3642cb",  1, "Test Brand", "Test Device", "123456789012345678")
	//operator.IsAccountExist("6126e612a9192b1b0c9628be")

	LogService.Info("Initializing Account Module Succeed.")
}


func (operator AccountOperator) RegisterAccount(
	account string,
	password string,
	userName string,
	email string,
) AccountResponse.AccountResponse {
	// Format params
	account = strings.TrimSpace(account)
	password = strings.TrimSpace(password)
	userName = strings.TrimSpace(userName)
	email = strings.TrimSpace(email)

	// Validate params correctness
	if !Formatter.CheckStringWithLength(account, 6, 20) {
		LogService.Warming("Fail to register account, account not valid. account=",account," password=",password," userName=",userName," email=",email,".")
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseParamsErrorWithMessage("Account field length error."),
		}
	}
	if !Formatter.CheckStringWithLength(password, 8, 30) {
		LogService.Warming("Fail to register account, password not valid. account=",account," password=",password," userName=",userName," email=",email,".")
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseParamsErrorWithMessage("Password field length error."),
		}
	}
	if !Formatter.CheckStringWithLength(userName, 2, 20) {
		LogService.Warming("Fail to register account, username not valid. account=",account," password=",password," userName=",userName," email=",email,".")
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseParamsErrorWithMessage("User name field length error."),
		}
	}
	if !Formatter.CheckStringWithLength(email, 6, 30) {
		LogService.Warming("Fail to register account, email not valid. account=",account," password=",password," userName=",userName," email=",email,".")
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseParamsErrorWithMessage("Email field length error."),
		}
	}
	if !Formatter.CheckStringIsValidEmail(email) {
		LogService.Warming("Fail to register account, email. account=",account," password=",password," userName=",userName," email=",email,".")
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseParamsErrorWithMessage("Email address is invalid."),
		}
	}

	// encrypt should run after length check
	password = Formatter.FormatPassword(password)

	// 1. Find out whether this account has been registered
	var queryResult []AccountV2Model.AccountDatabaseModel
	cursor, err := AccountV2DB.MongoDB.Collection(AccountV2DB.CollectionAccount).Find(context.TODO(), bson.M{
		"account": account,
	})
	if err != nil {
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while checking account in database",
			),
		}
	}
	if err=cursor.All(context.TODO(), &queryResult); err!=nil {
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while decoding check account database result",
			),
		}
	}
	if len(queryResult) > 0 {
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseParamsErrorWithMessage(
				"Account already been register",
			),
		}
	}

	// 2. Find out whether this email has been used
	cursor, err = AccountV2DB.MongoDB.Collection(AccountV2DB.CollectionAccount).Find(context.TODO(), bson.M{
		"email": email,
	})
	if err != nil {
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while checking email in database",
			),
		}
	}
	if err=cursor.All(context.TODO(), &queryResult); err!=nil {
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while decoding check email database result",
			),
		}
	}
	if len(queryResult) > 0 {
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseParamsErrorWithMessage(
				"Email already been register",
			),
		}
	}

	// 3. Register account and email
	var writeAccount = AccountV2Model.AccountDatabaseModel{
		ID:           primitive.NewObjectID(),
		Account:      account,
		Password:     password,
		Username:     userName,
		RegisterTime: Formatter.CurrentTimestampMillisecond(),
		Avatar:       "",
		Email:        email,
		Type:         1,
	}
	_, err = AccountV2DB.MongoDB.Collection(AccountV2DB.CollectionAccount).InsertOne(context.TODO(), writeAccount)
	if err != nil {
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while creating account",
			),
		}
	}

	return AccountResponse.AccountResponse{
		AccountBasicResponse: AccountResponse.CreateBasicResponseSuccess(),
		Account: AccountModule.AccountModel{
			ID: writeAccount.ID.Hex(),
			Account: writeAccount.Account,
			Password: writeAccount.Password,
			Username: writeAccount.Username,
			RegisterTime: writeAccount.RegisterTime,
			Avatar: writeAccount.Avatar,
			Email: writeAccount.Email,
			Type: writeAccount.Type,
		},
	}
}

const _maxTokenPerAccount = 5
const _tokenEffectTimeMillisecond = 1000 * 60 * 60 * 24 * 7
func (operator AccountOperator) LoginAccount(
	account string,
	password string,
	platform int,
	brand string,
	deviceName string,
	machineCode string,
) AccountResponse.AccountTokenResponse {
	// Format params
	account = strings.TrimSpace(account)
	password = strings.TrimSpace(password)
	password = Formatter.FormatPassword(password)

	// 1. Find out is there an account match the condition
	var queryAccountResult []AccountV2Model.AccountDatabaseModel
	cursor, err := AccountV2DB.MongoDB.Collection(AccountV2DB.CollectionAccount).Find(context.TODO(), bson.M{
		"account" : account,
		"password" : password,
	})
	if err!= nil {
		return AccountResponse.AccountTokenResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while validating account information",
			),
		}
	}
	if err=cursor.All(context.TODO(), &queryAccountResult); err!=nil {
		return AccountResponse.AccountTokenResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while decoding account information",
			),
		}
	}
	if len(queryAccountResult)<=0 {
		return AccountResponse.AccountTokenResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseParamsErrorWithMessage(
				"Account or password not correct",
			),
		}
	}

	// 2. Checkout the token is already reach the limit of token count
	var accountModel = queryAccountResult[0] // The account that login
	var queryTokenResult []AccountV2Model.AccountTokenDatabaseModel
	cursor, err = AccountV2DB.MongoDB.Collection(AccountV2DB.CollectionAccountToken).Find(context.TODO(), bson.M{
		"account_id" : accountModel.ID.Hex(),
	}, &options.FindOptions{
		Sort: bson.M{
			"expire_time" : -1,
		},
	})
	if err!= nil {
		return AccountResponse.AccountTokenResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while finding account token",
			),
		}
	}
	if err=cursor.All(context.TODO(), &queryTokenResult); err!=nil {
		return AccountResponse.AccountTokenResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while decoding account token information",
			),
		}
	}
	if len(queryTokenResult) > _maxTokenPerAccount - 1 {
		//     2 - If so, delete the oldest token
		var i int
		for i=_maxTokenPerAccount-1; i<len(queryTokenResult); i++ {
			tmpTokenModel := queryTokenResult[i]
			one, err := AccountV2DB.MongoDB.Collection(AccountV2DB.CollectionAccountToken).DeleteOne(context.TODO(), bson.M{
				"_id": tmpTokenModel.ID,
			})
			if err != nil {
				LogService.Warming("Error occur while deleting token. token =", tmpTokenModel.Token, " accountID =", tmpTokenModel.AccountID)
			}
			LogService.Debug("Token =", tmpTokenModel.Token, "is Expired. deleteCount =", one.DeletedCount, " accountID =", tmpTokenModel.AccountID)
		}
	}
	// 3. generate and insert one to database
	tokenUUID, err := uuid.NewRandom()
	if err != nil {
		LogService.Warming("Login failed, can not generate token. account =", account, ",password =", password)
		return AccountResponse.AccountTokenResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while decoding account token information",
			),
		}
	}
	var currentMillisecond = Formatter.CurrentTimestampMillisecond();
	var generateTokenModel = AccountV2Model.AccountTokenDatabaseModel{
		ID:          primitive.NewObjectID(),
		Token:       fmt.Sprintf("%s", tokenUUID),
		AccountID:   accountModel.ID.Hex(),
		UpdateTime:  currentMillisecond,
		Platform:    platform,
		Brand:       brand,
		DeviceName:  deviceName,
		MachineCode: machineCode,
		ExpireTime:  currentMillisecond + _tokenEffectTimeMillisecond,
	}
	_, err = AccountV2DB.MongoDB.Collection(AccountV2DB.CollectionAccountToken).InsertOne(context.TODO(), generateTokenModel)
	if err != nil {
		LogService.Warming("Error occur while adding token into database, err =", err)
		return AccountResponse.AccountTokenResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while adding token information",
			),
		}
	}
	return AccountResponse.AccountTokenResponse{
		AccountBasicResponse: AccountResponse.CreateBasicResponseSuccess(),
		AccountTokenModel: AccountModule.AccountTokenModel{
			ID:          generateTokenModel.ID.Hex(),
			Token:       generateTokenModel.Token,
			AccountID:   generateTokenModel.AccountID,
			UpdateTime:  generateTokenModel.UpdateTime,
			Platform:    generateTokenModel.Platform,
			Brand:       generateTokenModel.Brand,
			DeviceName:  generateTokenModel.DeviceName,
			MachineCode: generateTokenModel.MachineCode,
			ExpireTime:  generateTokenModel.ExpireTime,
		},
	}
}


func (operator AccountOperator) AuthToken(
	accountID string,
	token string,
	platform int,
	brand string,
	deviceName string,
	machineCode string,
) AccountResponse.AccountTokenResponse {
	// Format params
	token = strings.TrimSpace(token)
	brand = strings.TrimSpace(brand)
	deviceName = strings.TrimSpace(deviceName)
	machineCode = strings.TrimSpace(machineCode)

	// 1. Checkout whether the token is exist
	cursor, err := AccountV2DB.MongoDB.Collection(AccountV2DB.CollectionAccountToken).Find(context.TODO(), bson.M{
		"account_id" : accountID,
		"token" : token,
		"platform" : platform,
		"brand" : brand,
		"device_name" : deviceName,
		"machine_code" : machineCode,
	})
	if err != nil {
		return AccountResponse.AccountTokenResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while checking token",
			),
		}
	}
	var tokenQueryResult []AccountV2Model.AccountTokenDatabaseModel
	err = cursor.All(context.TODO(), &tokenQueryResult)
	if err != nil {
		return AccountResponse.AccountTokenResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while decoding token",
			),
		}
	}
	if len(tokenQueryResult) > 0 {
		// 1.       - If existed, checkout whether the token is expired.
		tokenQuery := tokenQueryResult[0]
		currentMillisecond := Formatter.CurrentTimestampMillisecond()
		if currentMillisecond > tokenQuery.ExpireTime {
			// 1.             - If so, return failed.
			return AccountResponse.AccountTokenResponse{
				AccountBasicResponse: AccountResponse.CreateBasicResponseTokenExpired(),
			}
		} else {
			// 1.             - If not, return succeed and refresh expired time.
			updateResult, err := AccountV2DB.MongoDB.Collection(AccountV2DB.CollectionAccountToken).UpdateMany(context.TODO(), bson.M{
				"account_id" : accountID,
				"token" : token,
				"platform" : platform,
				"brand" : brand,
				"device_name" : deviceName,
				"machine_code" : machineCode,
			}, bson.M{
				"$set" : bson.M{
					"expire_time" : currentMillisecond + _tokenEffectTimeMillisecond,
				},
			})
			if err!=nil {
				LogService.Warming("Error occur while updating token's expired time. Err = ", err)
			}
			LogService.Info("Token's expired time has been update. Token = ", token, " updateResult =", updateResult)
			return AccountResponse.AccountTokenResponse{
				AccountBasicResponse : AccountResponse.CreateBasicResponseSuccess(),
			}
		}
	}
	// 1.       - If not, return failed.
	return AccountResponse.AccountTokenResponse{
		AccountBasicResponse: AccountResponse.CreateBasicResponseTokenExpired(),
	}
}


func (operator AccountOperator) IsAccountExist(
	accountID string,
) AccountResponse.AccountResponse {

	// 1. Checkout whether the account is existed.
	accountIDQuery, err := primitive.ObjectIDFromHex(accountID)
	if err != nil {
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"AccountID not correct",
			),
		}
	}
	cursor, err := AccountV2DB.MongoDB.Collection(AccountV2DB.CollectionAccount).Find(context.TODO(), bson.M{
		"_id" : accountIDQuery,
	})
	if err != nil {
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while finding account",
			),
		}
	}
	var queryAccountResult []AccountV2Model.AccountDatabaseModel
	err = cursor.All(context.TODO(), &queryAccountResult)
	if err != nil {
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseInternalErrorWithMessage(
				"Error occur while parsing account",
			),
		}
	}
	if len(queryAccountResult) <= 0 {
		return AccountResponse.AccountResponse{
			AccountBasicResponse: AccountResponse.CreateBasicResponseParamsErrorWithMessage(
				"Account not exist",
			),
		}
	}
	queryAccount := queryAccountResult[0]
	return AccountResponse.AccountResponse{
		AccountBasicResponse: AccountResponse.CreateBasicResponseSuccess(),
		Account: AccountModule.AccountModel{
			ID:           queryAccount.ID.Hex(),
			Account:      queryAccount.Account,
			Password:     queryAccount.Password,
			Username:     queryAccount.Username,
			RegisterTime: queryAccount.RegisterTime,
			Avatar:       queryAccount.Avatar,
			Email:        queryAccount.Email,
			Type:         queryAccount.Type,
		},
	}
}



