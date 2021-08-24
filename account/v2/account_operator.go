package AccountV2

import (
	_ "sticker_board/account/manager"
	AccountResponse "sticker_board/account/manager/response"
	AccountV2 "sticker_board/account/v2/mongodb"
	"sticker_board/lib/log_service"
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
	return AccountResponse.AccountResponse{
		AccountBasicResponse: AccountResponse.CreateBasicResponseTodo(),
	}
}

func (operator AccountOperator) LoginAccount(
	account string,
	password string,
	platform int,
	brand string,
	deviceName string,
	machineCode string,
) AccountResponse.AccountTokenResponse {
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
	return AccountResponse.AccountTokenResponse{
		AccountBasicResponse: AccountResponse.CreateBasicResponseTodo(),
	}
}


func (operator AccountOperator) IsAccountExist(
	accountID int,
) AccountResponse.AccountResponse {
	return AccountResponse.AccountResponse{
		AccountBasicResponse: AccountResponse.CreateBasicResponseTodo(),

	}
}



