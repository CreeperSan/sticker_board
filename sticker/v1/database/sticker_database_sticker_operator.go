package StickerDatabase

import (
	"gorm.io/gorm"
	Application "sticker_board/application/database"
	"sticker_board/sticker/v1/model"
	"sticker_board/sticker/v1/response"
)

func getStickerBasic(accountID uint, stickerID uint) StickerDatabase.StickerBasicModel {
	db := Application.GetDB()

	var queryModel = StickerDatabase.StickerBasicModel{}

	result := db.
		Where(StickerDatabase.ColumnStickerBasicModelAccountID+" = ? and " +StickerDatabase.ColumnStickerBasicModelID+ " = ?", accountID, stickerID).
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

		result := tx.Delete(&StickerDatabase.StickerBasicModel{
			ID: stickerID,
			AccountID: accountID,
		})

		result = tx.Where( StickerDatabase.ColumnStickerPlainTextModelStickerID+ " = ?", stickerID).Delete(&StickerDatabase.StickerPlainTextModel{})

		result = tx.Where( StickerDatabase.ColumnStickerCategoryModelCategoryID+ " = ?", stickerID).Delete(StickerDatabase.StickerCategoryModel{})

		result = tx.Where( StickerDatabase.ColumnStickerTagModelID+ " = ?", stickerID).Delete(&StickerDatabase.StickerTagModel{})

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

