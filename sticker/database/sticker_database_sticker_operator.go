package StickerDatabase

import (
	"gorm.io/gorm"
	Application "sticker_board/application/database"
	Sticker "sticker_board/sticker/model"
	StickerResponse "sticker_board/sticker/response"
)

func getStickerBasic(accountID uint, stickerID uint) Sticker.StickerBasicModel {
	db := Application.GetDB()

	var queryModel = Sticker.StickerBasicModel{}

	err := db.
		Where(Sticker.ColumnStickerBasicModelAccountID +" = ? and " + Sticker.ColumnStickerBasicModelID + " = ?").
		First(&queryModel)

	if err != nil {
		return queryModel
	}
	return queryModel
}

func DeleteSticker(accountID uint, stickerID uint) StickerResponse.SimpleResponse {

	// check whether the sticker belong to this user
	queryResult := getStickerBasic(accountID, stickerID)

	if queryResult.ID <= 0 {
		return StickerResponse.SimpleResponse{
			Code: 403,
			Message: "Sticker does not exists",
		}
	}

	db := Application.GetDB()

	err := db.Transaction(func(tx *gorm.DB) error {

		tx.Delete(&Sticker.StickerBasicModel{
			ID: stickerID,
			AccountID: accountID,
		})

		tx.Delete(&Sticker.StickerPlainTextModel{
			StickerID: stickerID,
		})

		tx.Delete(&Sticker.StickerCategoryModel{
			StickerID: stickerID,
		})

		tx.Delete(&Sticker.StickerTagModel{
			StickerID: stickerID,
		})

		return nil
	})

	if err != nil {
		return StickerResponse.SimpleResponse{
			Code: 500,
			Message: "Server internal error",
		}
	}

	return StickerResponse.CreateSuccessSimpleResponse()

}

