package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
	OSSAlicloud "sticker_board/application/oss_alicloud"
	StickerModule "sticker_board/sticker/manager"
	StickerModuleModel "sticker_board/sticker/manager/model"
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

	// Format oss path to HTTP URL
	for index, tmpSticker := range queryResult.Stickers {
		switch tmpSticker.(type) {
		case StickerModuleModel.StickerPlainImageModel:
			var sticker = tmpSticker.(StickerModuleModel.StickerPlainImageModel)
			sticker.Url = "https://" + OSSAlicloud.BucketName + "." + OSSAlicloud.Endpoint + "/" + sticker.Url
			queryResult.Stickers[index] = sticker
		}
	}

	// Response
	ctx.JSON(ResponseParams{
		Code: queryResult.Code,
		Message: queryResult.Message,
		Data: queryResult.Stickers,
	})

}
