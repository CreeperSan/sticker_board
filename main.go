package main

import (
	"sticker_board/account/manager"
	StickerAccount "sticker_board/account/v1"
	"sticker_board/account/v2"
	ApiGeneral "sticker_board/api"
	"sticker_board/lib/log_service"
	StickerModule "sticker_board/sticker/manager"
	Sticker "sticker_board/sticker/v1"
	StickerV2 "sticker_board/sticker/v2"
)


func main() {
	LogService.Info("========= Sticker Board =========")

	// Initialize account module
	AccountModule.InstallOperator(&AccountV2.AccountOperator{})
	AccountModule.GetOperator().Initialize()

	// Initialize sticker module
	StickerModule.InstallOperator(&StickerV2.StickerOperator{})
	StickerModule.GetOperator().Initialize()

	StickerAccount.Init()
	Sticker.Init()
	ApiGeneral.Initialize()
}
