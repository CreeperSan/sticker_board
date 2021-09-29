package ApiV1

import (
	"github.com/kataras/iris/v12"
	ApiMiddleware "sticker_board/api/middleware"
	StickerModule "sticker_board/sticker/manager"
	StickerModuleModel "sticker_board/sticker/manager/model"
)

func InitializeTodoListSticker(app *iris.Application){
	stickerTodoList := app.Party("/api/sticker/v1/todo_list")

	stickerTodoList.Use(ApiMiddleware.LanguageMiddleware)
	stickerTodoList.Use(ApiMiddleware.AuthVersionMiddleware)
	stickerTodoList.Use(ApiMiddleware.AuthAccountMiddleware)

	stickerTodoList.Post("/create", createOrUpdateTodoListSticker)
	stickerTodoList.Post("/update", createOrUpdateTodoListSticker)
}

func createOrUpdateTodoListSticker(ctx iris.Context){
	type RequestParamTodoItem struct {
		State       int    `json:"state"` // 0 = not Finish, 1 = Finish
		Message     string `json:"message"`
		Description string `json:"description"`
	}
	type RequestParams struct {
		StickerID   string                 `json:"sticker_id"`
		Star        int                    `json:"star"`
		Status      int                    `json:"status"`
		Title       string                 `json:"title"`
		Background  string                 `json:"background"`
		Text        string                 `json:"text"`
		CategoryID  string                 `json:"category_id"`
		TagID       []string               `json:"tag_id"`
		IsPinned    bool                   `json:"is_pinned"`
		Description string                 `json:"description"`
		Todos       []RequestParamTodoItem `json:"todos"`
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
			Message: "Params Error",
		})
		return
	}

	// Convert todos from request params
	var todoSlice []StickerModuleModel.StickerTodoListItemModel
	for _, todoItem := range requestParams.Todos {
		todoSlice = append(todoSlice, StickerModuleModel.StickerTodoListItemModel{
			State: todoItem.State,
			Message: todoItem.Message,
			Description: todoItem.Description,
		})
	}

	createResult := StickerModule.GetOperator().CreateOrUpdateTodoListSticker(
		requestParams.StickerID,
		authResult.AccountID,
		requestParams.Star,
		requestParams.IsPinned,
		requestParams.Status,
		requestParams.Title,
		requestParams.Background,
		requestParams.TagID,
		requestParams.CategoryID,
		requestParams.Description,
		todoSlice,
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
