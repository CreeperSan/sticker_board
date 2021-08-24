package AccountV2

import (
	_ "sticker_board/account/manager"
	AccountResponse "sticker_board/account/manager/response"
	AccountV2 "sticker_board/account/v2/mongodb"
	Formatter "sticker_board/lib/formatter"
	"sticker_board/lib/log_service"
	"strings"
)

type AccountOperator struct { }


func (operator AccountOperator) Initialize(){
	LogService.Info("Initializing Account Module ...")
	client, mongoDB, mongoCtx := AccountV2.GetDB()
	if client==nil || mongoDB == nil || mongoCtx == nil {
		LogService.Error("Initializing Account Module Failed! Can not connect to mongodb database.")
		panic("Application exit.")
	}
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

	// TODO
	// 1. Find out whether this account has been registered
	// 2. Find out whether this email has been used
	// 3. Register account and email

	return AccountResponse.AccountResponse{
		AccountBasicResponse: AccountResponse.CreateBasicResponseTodo(),
	}
}

const _maxTokenPerAccount = 5
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

	// TODO
	// 1. Find out is there an account match the condition
	// 2. Checkout the token is already reach the limit of token count
	//      - If so, delete the oldest token then generate and insert one to database
	//      - If not, generate and insert one to database
	return AccountResponse.AccountTokenResponse{
		AccountBasicResponse: AccountResponse.CreateBasicResponseTodo(),
	}
}


func (operator AccountOperator) AuthToken(
	accountID int,
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

	// TODO
	// 1. Checkout whether the token is exist
	//        - If exist, checkout whether the token is expired.
	//              - If so, return failed.
	//              - If not, return succeed and refresh expired time.
	//        - If not, return failed.


	return AccountResponse.AccountTokenResponse{
		AccountBasicResponse: AccountResponse.CreateBasicResponseTodo(),
	}
}


func (operator AccountOperator) IsAccountExist(
	accountID int,
) AccountResponse.AccountResponse {

	// TODO
	// 1. Checkout whether the account is exist.

	return AccountResponse.AccountResponse{
		AccountBasicResponse: AccountResponse.CreateBasicResponseTodo(),
	}
}



