package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
	Formatter "sticker_board/lib/formatter"
	StickerModule "sticker_board/sticker/manager"
)

func InitializeCategory(app *iris.Application){
	categoryApi := app.Party("/api/category/v1")

	categoryApi.Use(ApiMiddleware.LanguageMiddleware)
	categoryApi.Use(ApiMiddleware.AuthVersionMiddleware)
	categoryApi.Use(ApiMiddleware.AuthAccountMiddleware)

	categoryApi.Post("/create", createCategory)
	categoryApi.Post("/delete", deleteCategory)
	categoryApi.Post("/list", queryCategoryList)
}


func createCategory(ctx iris.Context){
	type RequestParams struct {
		CategoryName string `json:"category_name"`
		Icon         string `json:"icon"`
		Color        int    `json:"color"`
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
			Message: ApiMiddleware.LanguageMiddlewareTrText(ctx, "common_params_error"),
		})
		return
	}

	// can't make category color to transparent
	if requestParams.Color <= 0 {
		requestParams.Color = 0xFF000000
	}
	requestParams.Color = 0xFF000000 // Currently not support custom category color
	requestParams.Icon = "icon:category_default" // Currently not support custom category color

	if !Formatter.CheckStringWithLength(requestParams.CategoryName, 1, 30) {
		ctx.JSON(ResponseParams{
			Code: 406,
			Message: ApiMiddleware.LanguageMiddlewareTrText(ctx, "common_params_error"),
		})
	}

	createResult := StickerModule.GetOperator().CreateCategory(authResult.AccountID, "", requestParams.CategoryName, requestParams.Icon, requestParams.Color)

	ctx.JSON(ResponseParams{
		Code: createResult.Code,
		Message: createResult.Message,
	})
}

func deleteCategory(ctx iris.Context){
	type RequestParams struct {
		CategoryID string `json:"category_id"`
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
			Message: ApiMiddleware.LanguageMiddlewareTrText(ctx, "common_params_error"),
		})
		return
	}

	deleteResult :=  StickerModule.GetOperator().DeleteCategory(authResult.AccountID, requestParams.CategoryID)

	ctx.JSON(ResponseParams{
		Code: deleteResult.Code,
		Message: deleteResult.Message,
	})
}

func queryCategoryList(ctx iris.Context){
	type ResponseParamsItem struct {
		CategoryID string   `json:"category_id"`
		CreateTime int64    `json:"create_time"`
		UpdateTime int64    `json:"update_time"`
		Name       string `json:"name"`
		Icon       string `json:"icon"`
		Color      int    `json:"color"`
	}
	type ResponseParams struct {
		Code    int                  `json:"code"`
		Message string               `json:"msg"`
		Data    []ResponseParamsItem `json:"data"`
	}

	authResult := ApiMiddleware.AuthAccountMiddleWareGetResponse(ctx)
	categoriesListResult := StickerModule.GetOperator().FindAllCategory(authResult.AccountID)

	if !categoriesListResult.IsSuccess() {
		ctx.JSON(ResponseParams{
			Code: categoriesListResult.Code,
			Message: categoriesListResult.Message,
		})
		return
	}

	dataList := []ResponseParamsItem{}
	for _, tmpItem := range categoriesListResult.Categories {
		dataList = append(dataList, ResponseParamsItem{
			CategoryID: tmpItem.ID,
			CreateTime: tmpItem.CreateTime,
			UpdateTime: tmpItem.UpdateTime,
			Name: tmpItem.Name,
			Icon: tmpItem.Icon,
			Color: tmpItem.Color,
		})
	}

	ctx.JSON(ResponseParams{
		Code: 200,
		Message: ApiMiddleware.LanguageMiddlewareTrText(ctx, "common_operate_success"),
		Data: dataList,
	})
}

