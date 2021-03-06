package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
	StickerModule "sticker_board/sticker/manager"
)

func InitializeStickerPlainText(app *iris.Application)  {
	stickerPlainText := app.Party("/api/sticker/v1/plain_text")

	stickerPlainText.Use(ApiMiddleware.LanguageMiddleware)
	stickerPlainText.Use(ApiMiddleware.AuthVersionMiddleware)
	stickerPlainText.Use(ApiMiddleware.AuthAccountMiddleware)

	stickerPlainText.Post("/create", createOrUpdatePlainTextSticker)
	stickerPlainText.Post("/update", createOrUpdatePlainTextSticker)
}

func createOrUpdatePlainTextSticker(ctx iris.Context){
	type RequestParams struct {
		StickerID  string   `json:"sticker_id"`
		Star       int      `json:"star"`
		Status     int      `json:"status"`
		Title      string   `json:"title"`
		Background string   `json:"background"`
		Text       string   `json:"text"`
		CategoryID string   `json:"category_id"`
		TagID      []string `json:"tag_id"`
		IsPinned   bool     `json:"is_pinned"`
	}
	type ResponseParams struct {
		Code    int         `json:"code"`
		Message string      `json:"msg"`
		Data    interface{} `json:"data"`
	}

	authResult := ApiMiddleware.AuthAccountMiddleWareGetResponse(ctx)

	// parse params
	requestParams := RequestParams{}
	err := ctx.ReadJSON(&requestParams)
	if err != nil {
		ctx.JSON(ResponseParams{
			Code: 406,
			Message: ApiMiddleware.LanguageMiddlewareTrText(ctx, "common_params_error"),
		})
		return
	}

	createResult := StickerModule.GetOperator().CreateOrUpdatePlainTextSticker(
		requestParams.StickerID,
		authResult.AccountID,
		requestParams.Star,
		requestParams.IsPinned,
		requestParams.Status,
		requestParams.Title,
		requestParams.Background,
		requestParams.TagID,
		requestParams.CategoryID,
		requestParams.Text,
	)

	if !createResult.IsSuccess() {
		ctx.JSON(ResponseParams{
			Code: createResult.Code,
			Message: createResult.Message,
		})
	} else {
		ctx.JSON(ResponseParams{
			Code: 200,
			Message: createResult.Message,
			Data: createResult.Sticker,
		})
	}

}


