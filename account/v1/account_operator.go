package v1

import (
	"sticker_board/account/v1/const"
	StickerBoardAccountDatabase "sticker_board/account/v1/database"
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
	StickerBoardAccountDatabase.Initialize()
	LogService.Info(StickerBoardAccount.TAG, "Initialized successful.")
}

func IsAccountExist(accountID uint) bool {
	return StickerBoardAccountDatabase.IsAccountExist(accountID)
}
