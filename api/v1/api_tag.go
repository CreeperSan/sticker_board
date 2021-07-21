package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
)

func InitializeTag(app *iris.Application){
	tagApi := app.Party("/api/tag/v1")

	tagApi.Use(ApiMiddleware.AuthAccountMiddleware)

	tagApi.Post("/create_tag", createTag)
}

func createTag(ctx iris.Context){
	type RequestParams struct {
		TagName string `json:"tag_name"`
		Icon    string `json:"icon"`
		Color   int `json:"color"`
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

	// can't make tag color to transparent
	if requestParams.Color <= 0 {
		requestParams.Color = 0xFF000000
	}

	requestParams.Icon = "icon:tag_default"

	ctx.JSON(ResponseParams{
		Code: 200,
		Message: "Todo createTag -> " + requestParams.TagName + "  " + string(authResult.AccountID),
	})
}
