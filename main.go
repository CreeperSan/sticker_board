package main

import (
	StickerAccount "sticker_board/account/v1"
	ApiGeneral "sticker_board/api"
	Sticker "sticker_board/sticker"
	"sticker_board/account/manager"
	"sticker_board/account/v2"
	"sticker_board/lib/log_service"
)


func main() {
	LogService.Info("========= Sticker Board =========")

	// Initialize account module
	AccountModule.InstallOperator(&AccountV2.AccountOperator{})
	AccountModule.GetOperator().Initialize()

	StickerAccount.Init()
	Sticker.Init()
	ApiGeneral.Initialize()
}
