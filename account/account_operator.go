package StickerBoardAccount

import (
	StickerBoardAccount "sticker_board/account/const"
	StickerBoardAccountDatabase "sticker_board/account/database"
	LogService "sticker_board/lib/log_service"
	SharedPreferences "sticker_board/lib/shared_preferences"
)

func GetVersionCode() int {
	return SharedPreferences.GetInt(StickerBoardAccount.SPVersionCode, 0)
}

func GetVersionName() string {
	return SharedPreferences.GetString(StickerBoardAccount.SPVersionName, "")
}

func Init(){
	LogService.Info(StickerBoardAccount.TAG, "Initializing ...")
	// Initializing Database
	StickerBoardAccountDatabase.Initialize()
	//StickerBoardAccountDatabase.RegisterAccount("acd", "1231", "s6d5a", "email.com")
	//StickerBoardAccountDatabase.RegisterAccount("account", "123213121", "UserName", "7s29@email.com")
	//StickerBoardAccountDatabase.RegisterAccount("account2", "12346578", "UserName2", "2@email.com")
	//StickerBoardAccountDatabase.LoginAccount("account2", "12346578", StickerBoardAccount.PlatformUndefined, "test", "test", "test")
	LogService.Info(StickerBoardAccount.TAG, "Initialized successful.")
}
