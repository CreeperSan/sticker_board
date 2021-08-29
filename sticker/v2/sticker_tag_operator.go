package StickerV2

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	AccountModule "sticker_board/account/manager"
	ApplicationDB "sticker_board/application/mongodb"
	Formatter "sticker_board/lib/formatter"
	LogService "sticker_board/lib/log_service"
	StickerModuleModel "sticker_board/sticker/manager/model"
	StickerModuleResponse "sticker_board/sticker/manager/response"
	StickerV2Model "sticker_board/sticker/v2/model"
)

func (operator *StickerOperator) CreateTag (accountID string, name string, icon string, color int) StickerModuleResponse.StickerTagResponse {
	// 1. Check the account is existed
	result := AccountModule.GetOperator().IsAccountExist(accountID)
	if !result.IsSuccess() {
		return StickerModuleResponse.StickerTagResponse{
			StickerResponse : StickerModuleResponse.CreateParamsErrorResponseWithMessage(
				"Account not existed",
			),
		}
	}

	// 2. Write a new tag into database
	var currentTimestamp = Formatter.CurrentTimestampMillisecond()
	var tarModelInsert = StickerV2Model.TagDatabaseModel{
		ID:         primitive.NewObjectID(),
		AccountID:  accountID,
		CreateTime: currentTimestamp,
		UpdateTime: currentTimestamp,
		Name:       name,
		Icon:       icon,
		Color:      color,
		Sort:       10000,
	}
	_, err := ApplicationDB.MongoDB.Collection(ApplicationDB.CollectionStickerTag).InsertOne(context.TODO(), tarModelInsert)
	if err != nil {
		LogService.Warming("Error occur while inserting tag model")
		return StickerModuleResponse.StickerTagResponse{
			StickerResponse : StickerModuleResponse.CreateInternalErrorResponseWithMessage(
				"Account not existed",
			),
		}
	}

	return StickerModuleResponse.StickerTagResponse{
		StickerResponse : StickerModuleResponse.CreateSuccessResponse(),
		Tag: StickerModuleModel.TagModel{
			ID: tarModelInsert.ID.Hex(),
			AccountID: tarModelInsert.AccountID,
			CreateTime: tarModelInsert.CreateTime,
			UpdateTime: tarModelInsert.UpdateTime,
			Name: tarModelInsert.Name,
			Icon: tarModelInsert.Icon,
			Color: tarModelInsert.Color,
		},
	}
}

func (operator *StickerOperator) DeleteTag (accountID string, tagID string) StickerModuleResponse.StickerResponse {
	// 1. Checkout the account is existed
	result := AccountModule.GetOperator().IsAccountExist(accountID)
	if !result.IsSuccess() {
		return StickerModuleResponse.CreateParamsErrorResponseWithMessage(
			"Account not existed",
		)
	}

	// 2. Delete tag
	tagObjectID, err := primitive.ObjectIDFromHex(tagID)
	if err != nil {
		return StickerModuleResponse.CreateParamsErrorResponseWithMessage(
			"Tag ID not exist",
		)
	}
	_, err = ApplicationDB.MongoDB.Collection(ApplicationDB.CollectionStickerTag).DeleteMany(context.TODO(), bson.M{
		"_id" : tagObjectID,
		"account_id" : accountID,
	})
	if err != nil {
		return StickerModuleResponse.CreateInternalErrorResponseWithMessage(
			"Tag ID not exist",
		)
	}

	return StickerModuleResponse.CreateSuccessResponse()
}

func (operator *StickerOperator) FindAllTag (accountID string) StickerModuleResponse.StickerTagArrayResponse {
	// 1. Checkout the account is existed
	result := AccountModule.GetOperator().IsAccountExist(accountID)
	if !result.IsSuccess() {
		return StickerModuleResponse.StickerTagArrayResponse{
			StickerResponse: StickerModuleResponse.CreateParamsErrorResponseWithMessage(
				"Account not existed",
			),
		}
	}

	// 2. Find all tags in database
	var queryResult []StickerV2Model.TagDatabaseModel
	cursor, err := ApplicationDB.MongoDB.Collection(ApplicationDB.CollectionStickerTag).Find(context.TODO(), bson.M{})
	if err!=nil {
		return StickerModuleResponse.StickerTagArrayResponse{
			StickerResponse: StickerModuleResponse.CreateInternalErrorResponseWithMessage(
				"Error occur while finding tags in database",
			),
		}
	}

	err = cursor.All(context.TODO(), &queryResult)
	if err != nil {
		return StickerModuleResponse.StickerTagArrayResponse{
			StickerResponse: StickerModuleResponse.CreateInternalErrorResponseWithMessage(
				"Error occur while finding tags in database",
			),
		}
	}

	return StickerModuleResponse.StickerTagArrayResponse{
		StickerResponse : StickerModuleResponse.CreateSuccessResponse(),
		Tags: queryResult,
	}
}
