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

func (operator StickerOperator) CreateCategory (
	accountID string,
	parentCategoryID string,
	name string,
	icon string,
	color int,
) StickerModuleResponse.StickerCategoryResponse {
	// 1. Check the account is existed
	result := AccountModule.GetOperator().IsAccountExist(accountID)
	if !result.IsSuccess() {
		return StickerModuleResponse.StickerCategoryResponse{
			StickerResponse : StickerModuleResponse.CreateParamsErrorResponseWithMessage(
				"Account not existed",
			),
		}
	}

	// 2. Write a new category into database
	var currentTimestamp = Formatter.CurrentTimestampMillisecond()
	var categoryModelInsert = StickerV2Model.CategoryDatabaseModel{
		ID:         primitive.NewObjectID(),
		AccountID:  accountID,
		CreateTime: currentTimestamp,
		UpdateTime: currentTimestamp,
		Name:       name,
		Icon:       icon,
		Color:      color,
		Sort:       10000,
	}
	_, err := ApplicationDB.MongoDB.Collection(ApplicationDB.CollectionStickerCategory).InsertOne(context.TODO(), categoryModelInsert)
	if err != nil {
		LogService.Warming("Error occur while inserting category model")
		return StickerModuleResponse.StickerCategoryResponse{
			StickerResponse : StickerModuleResponse.CreateInternalErrorResponseWithMessage(
				"Account not existed",
			),
		}
	}

	return StickerModuleResponse.StickerCategoryResponse{
		StickerResponse : StickerModuleResponse.CreateSuccessResponse(),
		Category: StickerModuleModel.CategoryModel{

		},
	}
}

func (operator StickerOperator) DeleteCategory (
	accountID string,
	categoryID string,
) StickerModuleResponse.StickerResponse {
	// 1. Checkout the account is existed
	result := AccountModule.GetOperator().IsAccountExist(accountID)
	if !result.IsSuccess() {
		return StickerModuleResponse.CreateParamsErrorResponseWithMessage(
			"Account not existed",
		)
	}

	// 2. Delete category
	categoryObjectID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return StickerModuleResponse.CreateParamsErrorResponseWithMessage(
			"category ID not exist",
		)
	}
	_, err = ApplicationDB.MongoDB.Collection(ApplicationDB.CollectionStickerCategory).DeleteMany(context.TODO(), bson.M{
		"_id" : categoryObjectID,
		"account_id" : accountID,
	})
	if err != nil {
		return StickerModuleResponse.CreateInternalErrorResponseWithMessage(
			"Category ID not exist",
		)
	}

	return StickerModuleResponse.CreateSuccessResponse()
}

func (operator StickerOperator) FindAllCategory (
	accountID string,
) StickerModuleResponse.StickerCategoryArrayResponse {
	// 1. Checkout the account is existed
	result := AccountModule.GetOperator().IsAccountExist(accountID)
	if !result.IsSuccess() {
		return StickerModuleResponse.StickerCategoryArrayResponse{
			StickerResponse: StickerModuleResponse.CreateParamsErrorResponseWithMessage(
				"Account not existed",
			),
		}
	}

	// 2. Find all categories in database
	var queryResult []StickerV2Model.CategoryDatabaseModel
	cursor, err := ApplicationDB.MongoDB.Collection(ApplicationDB.CollectionStickerCategory).Find(context.TODO(), bson.M{})
	if err!=nil {
		return StickerModuleResponse.StickerCategoryArrayResponse{
			StickerResponse: StickerModuleResponse.CreateInternalErrorResponseWithMessage(
				"Error occur while finding categories in database",
			),
		}
	}

	err = cursor.All(context.TODO(), &queryResult)
	if err != nil {
		return StickerModuleResponse.StickerCategoryArrayResponse{
			StickerResponse: StickerModuleResponse.CreateInternalErrorResponseWithMessage(
				"Error occur while finding categories in database",
			),
		}
	}

	var dataList []StickerModuleModel.CategoryModel
	for _, tmpItem := range queryResult {
		dataList = append(dataList, StickerModuleModel.CategoryModel{
			ID: tmpItem.ID.Hex(),
			CreateTime: tmpItem.CreateTime,
			UpdateTime: tmpItem.UpdateTime,
			Name: tmpItem.Name,
			Icon: tmpItem.Icon,
			Color: tmpItem.Color,
		})
	}

	return StickerModuleResponse.StickerCategoryArrayResponse{
		StickerResponse : StickerModuleResponse.CreateSuccessResponse(),
		Categories: dataList,
	}
}
