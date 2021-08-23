package AccountPackage

import AccountPackage "sticker_board/account/manager/model"

type AccountInterface interface {


	Initialize()


	RegisterAccount(
		account string,
		password string,
		userName string,
		email string,
	) (isSuccess bool, message string, accountModel AccountPackage.AccountDatabaseModel)


	LoginAccount(
		account string,
		password string,
		platform int,
		brand string,
		deviceName string,
		machineCode string,
	) (isSuccess bool, message string, accountModel AccountPackage.AccountTokenDatabaseModel)


	AuthToken(
		accountID int,
		token string,
		platform int,
		brand string,
		deviceName string,
		machineCode string,
	) (isSuccess bool, message string, accountModel AccountPackage.AccountTokenDatabaseModel)


	IsAccountExist(
		accountID int,
	) (isSuccess bool, message string, accountModel AccountPackage.AccountDatabaseModel)


}



