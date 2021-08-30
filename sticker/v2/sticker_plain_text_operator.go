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

func (operator *StickerOperator) CreatePlainTextSticker(
	accountID string,
	star int,
	isPinned bool,
	stickerStatus int,
	title string,
	background string,
	tagIDs []string,
	categoryID string,
	text string,
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

	// 2. Write a new plain text sticker to database
	currentMillisecond := Formatter.CurrentTimestampMillisecond()
	insertSticker := StickerV2Model.StickerDatabaseModel{
		ID:         primitive.NewObjectID(),
		Type:       StickerModuleConst.StickerTypePlainText,
		AccountID:  accountID,
		Star:       star,
		IsPinned:   isPinned,
		Status:     stickerStatus,
		Title:      title,
		Background: background,
		CreateTime: currentMillisecond,
		UpdateTime: currentMillisecond,
		SearchText: text,
		Sort:       10000,
		TagIDs:     tagIDs,
		CategoryID: categoryID,
		PlainText:  text,
	}
	_, err := ApplicationDB.MongoDB.Collection(ApplicationDB.CollectionSticker).InsertOne(context.TODO(), bson.M{
		"_id":         insertSticker.ID,
		"type":        insertSticker.Type,
		"account_id":  insertSticker.AccountID,
		"star":        insertSticker.Star,
		"is_pinned":   insertSticker.IsPinned,
		"status":      insertSticker.Status,
		"title":       insertSticker.Title,
		"background":  insertSticker.Background,
		"create_time": insertSticker.CreateTime,
		"update_time": insertSticker.UpdateTime,
		"search_text": insertSticker.SearchText,
		"sort":        insertSticker.Sort,
		"tags":        insertSticker.TagIDs,
		"category":    insertSticker.CategoryID,
		"plain_text":  insertSticker.PlainText,
	})
	if err != nil {
		return StickerModuleResponse.StickerSingleResponse{
			StickerResponse: StickerModuleResponse.CreateInternalErrorResponseWithMessage(
				"Error while inserting plain text sticker",
			),
		}
	}

	return StickerModuleResponse.StickerSingleResponse{
		StickerResponse: StickerModuleResponse.CreateSuccessResponse(),
		Sticker: StickerModuleModel.StickerPlainTextModel{
			StickerBasicModel : StickerModuleModel.StickerBasicModel{
				ID: insertSticker.ID.Hex(),
				Type: insertSticker.Type,
				AccountID: insertSticker.AccountID,
				Star: insertSticker.Star,
				IsPinned: insertSticker.IsPinned,
				Status: insertSticker.Status,
				Title: insertSticker.Title,
				Background: insertSticker.Background,
				CreateTime: insertSticker.CreateTime,
				UpdateTime: insertSticker.UpdateTime,
			},
			Text: insertSticker.PlainText,
		},
	}
}
