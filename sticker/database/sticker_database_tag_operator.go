package StickerDatabase

import (
	"sticker_board/account/v1/database"
	Application "sticker_board/application/database"
	StickerDatabase "sticker_board/sticker/model"
	StickerResponse "sticker_board/sticker/response"
)

func CreateTag(accountID uint, name string, icon string, color int, extra string) StickerResponse.SimpleResponse {
	db := Application.GetDB()

	// check if account exists
	if !Account.IsAccountExist(accountID) {
		return StickerResponse.SimpleResponse{
			Code: 400,
			Message: "Account not exist.",
		}
	}

	var tagModel = StickerDatabase.TagModel{
		AccountID: accountID,
		Name: name,
		Extra: extra,
		Color: color,
		Icon: icon,
	}

	db.Create(&tagModel)

	return StickerResponse.CreateSuccessSimpleResponse()
}

func DeleteTag(accountID uint, tagID uint) StickerResponse.SimpleResponse {
	db := Application.GetDB()

	// check if account exists
	if !Account.IsAccountExist(accountID) {
		return StickerResponse.SimpleResponse{
			Code: 400,
			Message: "Account not exist.",
		}
	}

	db.
		Where(StickerDatabase.ColumnTagModelAccountID+" = ? and "+StickerDatabase.ColumnTagModelID+" = ?", accountID, tagID).
		Delete(&StickerDatabase.TagModel{})

	return StickerResponse.SimpleResponse{
		Code: 200,
		Message: "Success",
	}
}

func QueryAllTag(accountID uint) StickerResponse.QueryTagResponse {
	db := Application.GetDB()

	// check if account exists
	if !Account.IsAccountExist(accountID) {
		return StickerResponse.QueryTagResponse{
			Code: 400,
			Message: "Account not exist.",
		}
	}

	var queryResult []StickerDatabase.TagModel

	db.
		Where(StickerDatabase.ColumnTagModelAccountID+" = ?", accountID).
		Find(&queryResult)


	return StickerResponse.QueryTagResponse{
		Code: 200,
		Message: "Success",
		Data: queryResult,
	}
}

