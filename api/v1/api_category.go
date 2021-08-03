package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
	Formatter "sticker_board/lib/formatter"
	StickerDatabase "sticker_board/sticker/database"
)

func InitializeCategory(app *iris.Application){
	categoryApi := app.Party("/api/category/v1")

	categoryApi.Use(ApiMiddleware.AuthVersionMiddleware)
	categoryApi.Use(ApiMiddleware.AuthAccountMiddleware)

	categoryApi.Post("/create", createCategory)
	categoryApi.Post("/delete", deleteCategory)
	categoryApi.Post("/list", queryCategoryList)
}


func createCategory(ctx iris.Context){
	type RequestParams struct {
		CategoryName string `json:"category_name"`
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

	// can't make category color to transparent
	if requestParams.Color <= 0 {
		requestParams.Color = 0xFF000000
	}
	requestParams.Color = 0xFF000000 // Currently not support custom category color
	requestParams.Icon = "icon:category_default" // Currently not support custom category color

	if !Formatter.CheckStringWithLength(requestParams.CategoryName, 1, 30) {
		ctx.JSON(ResponseParams{
			Code: 406,
			Message: "Params Error",
		})
	}

	StickerDatabase.CreateCategory(authResult.AccountID, 0, requestParams.CategoryName, requestParams.Icon, requestParams.Color, "")

	ctx.JSON(ResponseParams{
		Code: 200,
		Message: "Success",
	})
}

func deleteCategory(ctx iris.Context){
	type RequestParams struct {
		CategoryID uint `json:"category_id"`
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

		StickerDatabase.DeleteCategory(authResult.AccountID, requestParams.CategoryID)

	ctx.JSON(ResponseParams{
		Code: 200,
		Message: "Success",
	})
}

func queryCategoryList(ctx iris.Context){
	type ResponseParamsItem struct {
		CategoryID uint   `json:"category_id"`
		ParentID   uint   `json:"parent_id"`
		CreateTime int    `json:"create_time"`
		UpdateTime int    `json:"update_time"`
		Name       string `json:"name"`
		Icon       string `json:"icon"`
		Color      int    `json:"color"`
		Sort       int    `json:"sort"`
		Extra      string `json:"extra"`
	}
	type ResponseParams struct {
		Code    int                  `json:"code"`
		Message string               `json:"msg"`
		Data    []ResponseParamsItem `json:"data"`
	}

	authResult := ApiMiddleware.AuthAccountMiddleWareGetResponse(ctx)

	queryResponse := StickerDatabase.QueryAllCategory(authResult.AccountID)

	if queryResponse.Code != 200 {
		ctx.JSON(ResponseParams{
			Code: 500,
			Message: queryResponse.Message,
		})
	}

	var dataList []ResponseParamsItem
	for _, tmpItem := range queryResponse.Data {
		dataList = append(dataList, ResponseParamsItem{
			CategoryID: tmpItem.ID,
			CreateTime: tmpItem.CreateTime,
			UpdateTime: tmpItem.UpdateTime,
			Name: tmpItem.Name,
			Icon: tmpItem.Icon,
			Color: tmpItem.Color,
			Sort: tmpItem.Sort,
			Extra: tmpItem.Extra,
			ParentID: tmpItem.ParentID,
		})
	}

	ctx.JSON(ResponseParams{
		Code: 200,
		Message: "Success",
		Data: dataList,
	})
}

