package StickerDatabase

import (
	"sticker_board/account/v1/database"
	Application "sticker_board/application/database"
	StickerDatabase "sticker_board/sticker/model"
	StickerResponse "sticker_board/sticker/response"
)

func CreateCategory(accountID uint, parentID uint, name string, icon string, color int, extra string) StickerResponse.SimpleResponse {
	db := Application.GetDB()

	// check if account exists
	if !Account.IsAccountExist(accountID) {
		return StickerResponse.SimpleResponse{
			Code: 400,
			Message: "Account not exist.",
		}
	}

	var categoryModel = StickerDatabase.CategoryModel{
		AccountID: accountID,
		Name: name,
		Extra: extra,
		Color: color,
		Icon: icon,
		ParentID: parentID,
	}

	db.Create(&categoryModel)

	return StickerResponse.CreateSuccessSimpleResponse()
}

func DeleteCategory(accountID uint, categoryID uint) StickerResponse.SimpleResponse {
	db := Application.GetDB()

	// check if account exists
	if !Account.IsAccountExist(accountID) {
		return StickerResponse.SimpleResponse{
			Code: 400,
			Message: "Account not exist.",
		}
	}

	db.
		Where(StickerDatabase.ColumnCategoryModelAccountID+" = ? and "+StickerDatabase.ColumnCategoryModelID+" = ?", accountID, categoryID).
		Delete(&StickerDatabase.CategoryModel{})

	return StickerResponse.SimpleResponse{
		Code: 200,
		Message: "Success",
	}
}

func QueryAllCategory(accountID uint) StickerResponse.QueryCategoryResponse {
	db := Application.GetDB()

	// check if account exists
	if !Account.IsAccountExist(accountID) {
		return StickerResponse.QueryCategoryResponse{
			Code: 400,
			Message: "Account not exist.",
		}
	}

	var queryResult []StickerDatabase.CategoryModel

	db.
		Where(StickerDatabase.ColumnCategoryModelAccountID+" = ?", accountID).
		Find(&queryResult)


	return StickerResponse.QueryCategoryResponse{
		Code: 200,
		Message: "Success",
		Data: queryResult,
	}
}

