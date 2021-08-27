package AccountModule

import AccountModule "sticker_board/account/manager/response"

type AccountInterface interface {


	Initialize()


	RegisterAccount(
		account string,
		password string,
		userName string,
		email string,
	) AccountModule.AccountResponse


	LoginAccount(
		account string,
		password string,
		platform int,
		brand string,
		deviceName string,
		machineCode string,
	) AccountModule.AccountTokenResponse


	AuthToken(
		accountID string,
		token string,
		platform int,
		brand string,
		deviceName string,
		machineCode string,
	) AccountModule.AccountTokenResponse


	IsAccountExist(
		accountID string,
	) AccountModule.AccountResponse


}



