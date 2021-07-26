package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
	StickerDatabase "sticker_board/sticker/database"
)

func InitializeStickerCommon(app *iris.Application)  {
	stickerPlainText := app.Party("/api/sticker/v1")

	stickerPlainText.Use(ApiMiddleware.AuthAccountMiddleware)

	stickerPlainText.Post("/delete", deleteSticker)
}

func deleteSticker(ctx iris.Context) {
	type RequestParams struct {
		StickerID uint `json:"sticker_id"`
	}
	type ResponseParams struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
	}

	authResult := ApiMiddleware.AuthAccountMiddleWareGetResponse(ctx)

	// parse params
	requestParams := RequestParams{}
	err := ctx.ReadJSON(&requestParams)
	if err != nil {
		ctx.JSON(ResponseParams{
			Code: 406,
			Message: "Params Error",
		})
		return
	}

	databaseResponse := StickerDatabase.DeleteSticker(authResult.AccountID, requestParams.StickerID)
	if databaseResponse.Code != 200 {
		ctx.JSON(&ResponseParams{
			Code: databaseResponse.Code,
			Message: databaseResponse.Message,
		})
		return
	}

	ctx.JSON(&ResponseParams{
		Code: 200,
		Message: "Success",
	})

}
