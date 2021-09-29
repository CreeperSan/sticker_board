package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
	StickerModule "sticker_board/sticker/manager"
)

func InitializeStickerCommon(app *iris.Application)  {
	stickerPlainText := app.Party("/api/sticker/v1")

	stickerPlainText.Use(ApiMiddleware.LanguageMiddleware)
	stickerPlainText.Use(ApiMiddleware.AuthVersionMiddleware)
	stickerPlainText.Use(ApiMiddleware.AuthAccountMiddleware)

	stickerPlainText.Post("/delete", deleteSticker)
}

func deleteSticker(ctx iris.Context) {
	type RequestParams struct {
		StickerID string `json:"sticker_id"`
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

	deleteResult := StickerModule.GetOperator().DeleteSticker(authResult.AccountID, requestParams.StickerID)
	if !deleteResult.IsSuccess() {
		ctx.JSON(&ResponseParams{
			Code: deleteResult.Code,
			Message: deleteResult.Message,
		})
		return
	}

	ctx.JSON(&ResponseParams{
		Code: 200,
		Message: "Success",
	})

}
