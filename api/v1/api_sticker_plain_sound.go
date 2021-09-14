package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
	LogService "sticker_board/lib/log_service"
	StickerModule "sticker_board/sticker/manager"
)

func InitializeStickerPlainSound(app *iris.Application) {
	stickerPlainText := app.Party("api/sticker/v1/plain_sound")

	stickerPlainText.Use(ApiMiddleware.AuthVersionMiddleware)
	stickerPlainText.Use(ApiMiddleware.AuthAccountMiddleware)

	stickerPlainText.Post("/create", createPlainSoundSticker)
}

func createPlainSoundSticker(ctx iris.Context) {
	type RequestParams struct {
		Star             int      `json:"star"`
		Status           int      `json:"status"`
		Title            string   `json:"title"`
		Background       string   `json:"background"`
		CategoryID       string   `json:"category_id"`
		TagID            []string `json:"tag_id"`
		IsPinned         bool     `json:"is_pinned"`
		SoundDescription string   `json:"description"`
		SoundPath        string   `json:"path"`
		SoundDuration    int      `json:"duration"`
	}
	type ResponseParams struct {
		Code    int         `json:"code"`
		Message string      `json:"msg"`
		Data    interface{} `json:"data"`
	}

	authResult := ApiMiddleware.AuthAccountMiddleWareGetResponse(ctx)

	// parse params
	var requestParams RequestParams
	err := ctx.ReadJSON(&requestParams)
	if err != nil {
		ctx.JSON(ResponseParams{
			Code: 401,
			Message: "Params Error",
		})
		return
	}

	// Validate params
	if len(requestParams.SoundPath) <= 0 {
		ctx.JSON(ResponseParams{
			Code: 401,
			Message: "Sound path is empty",
		})
		return
	}

	LogService.Info("Request Params : ", requestParams)

	// Extract file from form request
	createResult := StickerModule.GetOperator().CreatePlainSoundSticker(
		authResult.AccountID,
		requestParams.Star,
		requestParams.IsPinned,
		requestParams.Status,
		requestParams.Title,
		requestParams.Background,
		requestParams.TagID,
		requestParams.CategoryID,
		requestParams.SoundPath,
		requestParams.SoundDescription,
		requestParams.SoundDuration,
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
