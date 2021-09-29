package main

import (
	"sticker_board/account/manager"
	"sticker_board/account/v2"
	ApiGeneral "sticker_board/api"
	"sticker_board/application/oss_alicloud"
	"sticker_board/i18n"
	Localization "sticker_board/lib/localization"
	"sticker_board/lib/log_service"
	StickerModule "sticker_board/sticker/manager"
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

	// Initialize Alicloud OSS
	OSSAlicloud.Initialize()

	// Initialize i18n
	Localization.RegisterBundle(&i18n.BundleEnglish{})
	Localization.RegisterBundle(&i18n.BundleSimplifyChinese{})
	Localization.SetDefaultBundleName("en_US")

	ApiGeneral.Initialize()
}
