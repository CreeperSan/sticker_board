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

	result := db.
		Where(Sticker.ColumnStickerBasicModelAccountID +" = ? and " + Sticker.ColumnStickerBasicModelID + " = ?", accountID, stickerID).
		First(&queryModel)

	if result.Error != nil {
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

		result := tx.Delete(&Sticker.StickerBasicModel{
			ID: stickerID,
			AccountID: accountID,
		})

		result = tx.Where( Sticker.ColumnStickerPlainTextModelStickerID + " = ?", stickerID).Delete(&Sticker.StickerPlainTextModel{})

		result = tx.Where( Sticker.ColumnStickerCategoryModelCategoryID + " = ?", stickerID).Delete(Sticker.StickerCategoryModel{})

		result = tx.Where( Sticker.ColumnStickerTagModelID + " = ?", stickerID).Delete(&Sticker.StickerTagModel{})

		if result.Error != nil {
			return result.Error
		}

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

