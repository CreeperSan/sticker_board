package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
	StickerDatabase "sticker_board/sticker/database"
)

func InitializeStickerPlainText(app *iris.Application)  {
	stickerPlainText := app.Party("/api/sticker/v1/plain_text")

	stickerPlainText.Use(ApiMiddleware.AuthAccountMiddleware)

	stickerPlainText.Post("/create", createPlainTextSticker)
}

func createPlainTextSticker(ctx iris.Context){
	type RequestParams struct {
		Star int `json:"star"`
		Pinned int `json:"pinned"`
		Status int `json:"status"`
		Title string `json:"title"`
		Background string `json:"background"`
		Extra string `json:"extra"`
		Text string `json:"text"`
		CategoryID uint `json:"category_id"`
		TagID []uint `json:"tag_id"`
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

	databaseResponse := StickerDatabase.CreateStickerPlainText(
		authResult.AccountID,
		requestParams.Star,
		requestParams.Pinned,
		requestParams.Title,
		requestParams.Text,
		requestParams.CategoryID,
		requestParams.TagID,
	)

	if databaseResponse.Code == 200 {
		ctx.JSON(ResponseParams{
			Code: 200,
			Message: "Success",
		})
	} else {
		ctx.JSON(ResponseParams{
			Code: 200,
			Message: databaseResponse.Message,
		})
	}

}


