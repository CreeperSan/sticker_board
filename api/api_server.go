package ApiGeneral

import (
	"github.com/kataras/iris/v12"
	ApiV1 "sticker_board/api/v1"
	LogService "sticker_board/lib/log_service"
)

func Initialize()  {
	app := iris.New()

	// Initializing web server
	ApiV1.InitializeAccount(app)
	ApiV1.InitializeTag(app)
	ApiV1.InitializeCategory(app)
	ApiV1.InitializeStickerCommon(app)
	ApiV1.InitializeStickerPlainText(app)
	ApiV1.InitializeStickerPlainImage(app)
	ApiV1.InitializeStickerPlainSound(app)
	ApiV1.InitializeTodoListSticker(app)
	ApiV1.InitializeQuery(app)
	ApiV1.InitializeOSS(app)

	startServerError := app.Listen(":8080")
	if startServerError != nil {
		LogService.Error("Error occurred while starting web server.", startServerError)
		panic(startServerError)
		return
	}
}

