package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
)

func InitializeTag(app *iris.Application){
	tagApi := app.Party("/api/tag/v1")

	tagApi.Use(ApiMiddleware.AuthVersionMiddleware)
	tagApi.Use(ApiMiddleware.AuthAccountMiddleware)

	tagApi.Post("/create", createTag)
	tagApi.Post("/delete", deleteTag)
	tagApi.Post("/list", queryTagList)
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

	//authResult := ApiMiddleware.AuthAccountMiddleWareGetResponse(ctx)
	//
	//// parse params
	//requestParams := RequestParams{}
	//err := ctx.ReadJSON(&requestParams)
	//if err != nil {
	//	ctx.JSON(ResponseParams{
	//		Code: 406,
	//		Message: "Params Error",
	//	})
	//	return
	//}
	//
	//// can't make tag color to transparent
	//if requestParams.Color <= 0 {
	//	requestParams.Color = 0xFF000000
	//}
	//requestParams.Color = 0xFF000000 // Currently not support custom tag color
	//requestParams.Icon = "icon:tag_default" // Currently not support custom tag color
	//
	//if !Formatter.CheckStringWithLength(requestParams.TagName, 1, 30) {
	//	ctx.JSON(ResponseParams{
	//		Code: 406,
	//		Message: "Params Error",
	//	})
	//}
	//
	//StickerDatabase.CreateTag(authResult.AccountID, requestParams.TagName, requestParams.Icon, requestParams.Color, "")
	//
	//ctx.JSON(ResponseParams{
	//	Code: 200,
	//	Message: "Success",
	//})
}

func deleteTag(ctx iris.Context){
	type RequestParams struct {
		TagID uint `json:"tag_id"`
	}
	type ResponseParams struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
	}

	//authResult := ApiMiddleware.AuthAccountMiddleWareGetResponse(ctx)
	//
	//// parse params
	//requestParams := RequestParams{}
	//err := ctx.ReadJSON(&requestParams)
	//if err != nil {
	//	ctx.JSON(ResponseParams{
	//		Code: 406,
	//		Message: "Params Error",
	//	})
	//	return
	//}
	//
	//StickerDatabase.DeleteTag(authResult.AccountID, requestParams.TagID)
	//
	//ctx.JSON(ResponseParams{
	//	Code: 200,
	//	Message: "Success",
	//})
}

func queryTagList(ctx iris.Context){
	type ResponseParamsItem struct {
		TagID      uint   `json:"tag_id"`
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

	//authResult := ApiMiddleware.AuthAccountMiddleWareGetResponse(ctx)
	//
	//queryResponse := StickerDatabase.QueryAllTag(authResult.AccountID)
	//
	//if queryResponse.Code != 200 {
	//	ctx.JSON(ResponseParams{
	//		Code: 500,
	//		Message: queryResponse.Message,
	//	})
	//}
	//
	//var dataList []ResponseParamsItem
	//for _, tmpItem := range queryResponse.Data {
	//	dataList = append(dataList, ResponseParamsItem{
	//		TagID: tmpItem.ID,
	//		CreateTime: tmpItem.CreateTime,
	//		UpdateTime: tmpItem.UpdateTime,
	//		Name: tmpItem.Name,
	//		Icon: tmpItem.Icon,
	//		Color: tmpItem.Color,
	//		Sort: tmpItem.Sort,
	//		Extra: tmpItem.Extra,
	//	})
	//}
	//
	//ctx.JSON(ResponseParams{
	//	Code: 200,
	//	Message: "Success",
	//	Data: dataList,
	//})
}
