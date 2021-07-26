package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
)

func InitializeQuery(app *iris.Application)  {
	stickerPlainText := app.Party("/api/query/v1")

	stickerPlainText.Use(ApiMiddleware.AuthAccountMiddleware)

	stickerPlainText.Post("sticker", querySticker)
}

func querySticker(ctx iris.Context){
	ctx.WriteString("TODO")
}
