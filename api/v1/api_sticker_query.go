package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
	StickerModule "sticker_board/sticker/manager"
)

func InitializeQuery(app *iris.Application)  {
	stickerPlainText := app.Party("/api/sticker/v1")

	stickerPlainText.Use(ApiMiddleware.AuthVersionMiddleware)
	stickerPlainText.Use(ApiMiddleware.AuthAccountMiddleware)

	stickerPlainText.Post("/query", querySticker)
}

func querySticker(ctx iris.Context){
	type RequestParams struct {
		Page      int    `json:"page"`
		PageSize  int    `json:"page_size"`
	}
	type ResponseParams struct {
		Code    int           `json:"code"`
		Message string        `json:"msg"`
		Data    []interface{} `json:"data"`
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

	queryResult := StickerModule.GetOperator().FindSticker(
		authResult.AccountID,
		requestParams.Page,
		requestParams.PageSize,
	)

	if !queryResult.IsSuccess() {
		ctx.JSON(ResponseParams{
			Code: queryResult.Code,
			Message: queryResult.Message,
		})
		return
	}

	ctx.JSON(ResponseParams{
		Code: queryResult.Code,
		Message: queryResult.Message,
		Data: queryResult.Stickers,
	})

}
