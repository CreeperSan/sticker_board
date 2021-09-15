package StickerV2

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	AccountModule "sticker_board/account/manager"
	ApplicationDB "sticker_board/application/mongodb"
	StickerModuleConst "sticker_board/sticker/manager/const"
	StickerModuleModel "sticker_board/sticker/manager/model"
	StickerModuleResponse "sticker_board/sticker/manager/response"
	StickerV2Model "sticker_board/sticker/v2/model"
)

func (operator *StickerOperator) DeleteSticker(accountID string, stickerID string) StickerModuleResponse.StickerResponse {
	// 1. Checkout the params is in right format
	stickerObjectID, err := primitive.ObjectIDFromHex(stickerID)
	if err != nil {
		return StickerModuleResponse.CreateParamsErrorResponseWithMessage(
			"StickerID not correct",
		)
	}

	// 2. Check whether this account is exist
	checkResult := AccountModule.GetOperator().IsAccountExist(accountID)
	if !checkResult.IsSuccess() {
		return StickerModuleResponse.CreateParamsErrorResponseWithMessage(
			"Account does not exist",
		)
	}

	// 3. Delete sticker form database
	_, err = ApplicationDB.MongoDB.Collection(ApplicationDB.CollectionSticker).DeleteOne(context.TODO(), bson.M{
		"_id":        stickerObjectID,
		"account_id": accountID,
	})
	if err != nil {
		return StickerModuleResponse.CreateInternalErrorResponseWithMessage(
			"Error occur while deleting sticker",
		)
	}
	return StickerModuleResponse.CreateSuccessResponse()
}

func (operator *StickerOperator) FindSticker (accountID string, page int, pageSize int) StickerModuleResponse.StickerArrayResponse {
	// 1. Check whether this account is exist
	checkResult := AccountModule.GetOperator().IsAccountExist(accountID)
	if !checkResult.IsSuccess() {
		return StickerModuleResponse.StickerArrayResponse{
			StickerResponse : StickerModuleResponse.CreateParamsErrorResponseWithMessage(
				"Account does not exist",
			),
		}
	}

	// 2. Find sticker in database
 	var pSkip int64
 	var pLimit int64
	pSkip = int64(page * pageSize)
	pLimit = int64(pageSize)
	cursor, err := ApplicationDB.MongoDB.Collection(ApplicationDB.CollectionSticker).Find(context.TODO(), bson.M{
		"account_id": accountID,
	}, &options.FindOptions{
		Skip:  &pSkip,
		Limit: &pLimit,
	})
	if err != nil {
		return StickerModuleResponse.StickerArrayResponse{
			StickerResponse : StickerModuleResponse.CreateInternalErrorResponseWithMessage(
				"Error occur while finding stickers",
			),
		}
	}
	var queryResult []StickerV2Model.StickerDatabaseModel
	err = cursor.All(context.TODO(), &queryResult)
	if err != nil {
		return StickerModuleResponse.StickerArrayResponse{
			StickerResponse : StickerModuleResponse.CreateInternalErrorResponseWithMessage(
				"Error occur while parsing stickers",
			),
		}
	}


	var stickerModelArray []interface{}
	for _, databaseModel := range queryResult {
		basicModel := StickerModuleModel.StickerBasicModel{
			ID: databaseModel.ID.Hex(),
			Type: databaseModel.Type,
			AccountID: databaseModel.AccountID,
			Star: databaseModel.Star,
			IsPinned: databaseModel.IsPinned,
			Status: databaseModel.Status,
			Title: databaseModel.Title,
			Background: databaseModel.Background,
			CreateTime: databaseModel.CreateTime,
			UpdateTime: databaseModel.UpdateTime,
		}
		switch databaseModel.Type {
		case StickerModuleConst.StickerTypePlainText:{
			stickerModelArray = append(stickerModelArray, StickerModuleModel.StickerPlainTextModel{
				StickerBasicModel : basicModel,
				Text: databaseModel.PlainText,
			})
			break
		}
		case StickerModuleConst.StickerTypePlainImage:{
			stickerModelArray = append(stickerModelArray, StickerModuleModel.StickerPlainImageModel{
				StickerBasicModel : basicModel,
				Url: databaseModel.PlainImageUrl,
				Description: databaseModel.PlainImageDescription,
			})
			break
		}
		case StickerModuleConst.StickerTypePlainSound:{
			stickerModelArray = append(stickerModelArray, StickerModuleModel.StickerPlainSoundModel{
				StickerBasicModel : basicModel,
				Url: databaseModel.PlainSoundUrl,
				Description: databaseModel.PlainSoundDescription,
				Duration: databaseModel.PlainSoundDuration,
			})
			break
		}
		case StickerModuleConst.StickerTypeTodoList:{
			var todoItemList []StickerModuleModel.StickerTodoListItemModel
			for _, todoAction := range databaseModel.TodoListAction {
				todoItemList = append(todoItemList, StickerModuleModel.StickerTodoListItemModel{
					State: todoAction.State,
					Message: todoAction.Message,
					Description: todoAction.Description,
				})
			}
			stickerModelArray = append(stickerModelArray, StickerModuleModel.StickerTodoListModel{
				StickerBasicModel : basicModel,
				Todos: todoItemList,
				Description: databaseModel.TodoListDescription,
			})
			break
		}
		}
	}
	return StickerModuleResponse.StickerArrayResponse{
		StickerResponse : StickerModuleResponse.CreateSuccessResponse(),
		Stickers:         stickerModelArray,
	}
}

