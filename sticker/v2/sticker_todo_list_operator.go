package StickerV2

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	AccountModule "sticker_board/account/manager"
	ApplicationDB "sticker_board/application/mongodb"
	Formatter "sticker_board/lib/formatter"
	StickerModuleConst "sticker_board/sticker/manager/const"
	StickerModuleModel "sticker_board/sticker/manager/model"
	StickerModuleResponse "sticker_board/sticker/manager/response"
	StickerV2Model "sticker_board/sticker/v2/model"
)

func (operator *StickerOperator) CreateTodoListSticker(
	accountID string,
	star int,
	isPinned bool,
	stickerStatus int,
	title string,
	background string,
	tagIDs []string,
	categoryID string,
	description string,
	todos []StickerModuleModel.StickerTodoListItemModel,
) StickerModuleResponse.StickerSingleResponse {
	// 1. Check the account is existed
	result := AccountModule.GetOperator().IsAccountExist(accountID)
	if !result.IsSuccess() {
		return StickerModuleResponse.StickerSingleResponse{
			StickerResponse: StickerModuleResponse.CreateParamsErrorResponseWithMessage(
				"Account not existed",
			),
		}
	}

	// 2. Convert interface model to module model
	var todoListAction []StickerV2Model.TodoListAction
	for _, todoItems := range todos {
		todoListAction = append(todoListAction, StickerV2Model.TodoListAction{
			State:       todoItems.State,
			Message:     todoItems.Message,
			Description: todoItems.Description,
		})
	}

	// 3. Write a new todo list sticker to database
	currentMillisecond := Formatter.CurrentTimestampMillisecond()
	insertSticker := StickerV2Model.StickerDatabaseModel{
		ID:                  primitive.NewObjectID(),
		Type:                StickerModuleConst.StickerTypePlainImage,
		AccountID:           accountID,
		Star:                star,
		IsPinned:            isPinned,
		Status:              stickerStatus,
		Title:               title,
		Background:          background,
		CreateTime:          currentMillisecond,
		UpdateTime:          currentMillisecond,
		SearchText:          title,
		Sort:                10000,
		TagIDs:              tagIDs,
		CategoryID:          categoryID,
		TodoListDescription: description,
		TodoListAction:      todoListAction,
	}
	_, err := ApplicationDB.MongoDB.Collection(ApplicationDB.CollectionSticker).InsertOne(context.TODO(), bson.M{
		"_id":                   insertSticker.ID,
		"type":                  insertSticker.Type,
		"account_id":            insertSticker.AccountID,
		"star":                  insertSticker.Star,
		"is_pinned":             insertSticker.IsPinned,
		"status":                insertSticker.Status,
		"title":                 insertSticker.Title,
		"background":            insertSticker.Background,
		"create_time":           insertSticker.CreateTime,
		"update_time":           insertSticker.UpdateTime,
		"search_text":           insertSticker.SearchText,
		"sort":                  insertSticker.Sort,
		"tags":                  insertSticker.TagIDs,
		"category":              insertSticker.CategoryID,
		"todo_list_description": insertSticker.TodoListDescription,
		"todo_list_action":      insertSticker.TodoListAction,
	})
	if err != nil {
		return StickerModuleResponse.StickerSingleResponse{
			StickerResponse: StickerModuleResponse.CreateInternalErrorResponseWithMessage(
				"Error while inserting todo list sticker",
			),
		}
	}

	return StickerModuleResponse.StickerSingleResponse{
		StickerResponse: StickerModuleResponse.CreateSuccessResponse(),
		Sticker: insertSticker,
	}
}