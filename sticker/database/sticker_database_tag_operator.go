package StickerDatabase

import (
	Account "sticker_board/account/database"
	Application "sticker_board/application/database"
	StickerDatabase "sticker_board/sticker/model"
	StickerResponse "sticker_board/sticker/response"
)

func CreateTag(accountID uint, name string) StickerResponse.SimpleResponse {
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
		Extra: "",
		Color: -1,
		Icon: "",
	}

	db.Create(&tagModel)

	return StickerResponse.CreateSuccessSimpleResponse()
}

func DeleteTag(accountID uint, tagID uint) StickerResponse.SimpleResponse {
	//db := Application.GetDB()

	//db.Where(+" = ? and "++" = ?", accountID, tagID)

	return StickerResponse.SimpleResponse{
		Code: 400,
		Message: "Todo",
	}
}

