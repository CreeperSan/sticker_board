package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
)

func InitializeStickerCommon(app *iris.Application)  {
	stickerPlainText := app.Party("/api/sticker/v1")

	stickerPlainText.Use(ApiMiddleware.AuthAccountMiddleware)

	stickerPlainText.Post("delete", deleteSticker)
}

func deleteSticker(ctx iris.Context)  {
	ctx.WriteString("TODO")
}
