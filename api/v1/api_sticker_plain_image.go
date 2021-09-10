package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
	LogService "sticker_board/lib/log_service"
	StickerModule "sticker_board/sticker/manager"
)

func InitializeStickerPlainImage(app *iris.Application)  {
	stickerPlainText := app.Party("/api/sticker/v1/plain_image")

	stickerPlainText.Use(ApiMiddleware.AuthVersionMiddleware)
	stickerPlainText.Use(ApiMiddleware.AuthAccountMiddleware)

	stickerPlainText.Post("/create", createPlainImageSticker)
}

func createPlainImageSticker(ctx iris.Context){
	type RequestParams struct {
		Star             int      `json:"star"`
		Status           int      `json:"status"`
		Title            string   `json:"title"`
		Background       string   `json:"background"`
		CategoryID       string   `json:"category_id"`
		TagID            []string `json:"tag_id"`
		IsPinned         bool     `json:"is_pinned"`
		ImageDescription string   `json:"description"`
		ImagePath        string   `json:"image_path"`
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
	if len(requestParams.ImagePath) <= 0 {
		ctx.JSON(ResponseParams{
			Code: 401,
			Message: "Image path is empty",
		})
		return
	}

	LogService.Info("Request Params : ", requestParams)

	// Extract file from form request
	createResult := StickerModule.GetOperator().CreatePlainImageSticker(
		authResult.AccountID,
		requestParams.Star,
		requestParams.IsPinned,
		requestParams.Status,
		requestParams.Title,
		requestParams.Background,
		requestParams.TagID,
		requestParams.CategoryID,
		requestParams.ImagePath,
		requestParams.ImageDescription,
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


