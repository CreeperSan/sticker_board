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

	type ResponseParams struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
	}

	//authResult := ApiMiddleware.ParseFromAuthMiddleware(ctx)

	authResult := ctx.Values().Get("_AccountInfo").(ApiMiddleware.AuthAccountMiddlewareResult)

	ctx.JSON(ResponseParams{
		Code: 200,
		Message: "Todo" + string(authResult.AccountID),
	})
}
