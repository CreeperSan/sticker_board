package StickerDatabase

import (
	"gorm.io/gorm"
	Application "sticker_board/application/database"
	StickerConst "sticker_board/sticker/const"
	Sticker "sticker_board/sticker/model"
	StickerResponse "sticker_board/sticker/response"
)

func CreateStickerPlainText(accountID uint, star int, pinned int, title string, content string, categoryID uint, tagIDArray []uint) StickerResponse.SimpleResponse {
	db := Application.GetDB()

	errTransaction := db.Transaction(func(tx *gorm.DB) error {
		stickerBasicModel := Sticker.StickerBasicModel{
			Type: StickerConst.TypePlainText,
			AccountID: accountID,
			Star: star,
			Pinned: pinned,
			Status: StickerConst.StatusNew,
			Title: title,
			Background: "color:default",
			SearchText: content,
			Extra: "",
		}

		dbResult := tx.Create(&stickerBasicModel)
		if dbResult.Error != nil {
			return dbResult.Error
		}

		dbResult = tx.Create(&Sticker.StickerPlainTextModel{
			StickerID: stickerBasicModel.ID,
			Text: content,
		})
		if dbResult.Error != nil {
			return dbResult.Error
		}

		// Bind categoryID with sticker
		if categoryID > 0 {
			// find out whether the category is exist and belong to user
			var categoryModel = Sticker.CategoryModel{}
			dbResult = tx.
				Where(Sticker.ColumnCategoryModelID+ " = ? and " + Sticker.ColumnCategoryModelAccountID + " = ?", categoryID, accountID).
				First(&categoryModel)
			if dbResult.Error != nil {
				return dbResult.Error
			}
			// database operation
			dbResult = tx.Create(&Sticker.StickerCategoryModel{
				StickerID: stickerBasicModel.ID,
				CategoryID: categoryID,
			})
			if dbResult.Error != nil {
				return dbResult.Error
			}
		}

		// Bind tag with tag
		if len(tagIDArray) > 0 {
			for i := 0; i < len(tagIDArray); i++ {
				var tagID = tagIDArray[i]
				// find out whether the tag is exist and belong to user
				var tagModel = Sticker.TagModel{}
				dbResult = tx.
					Where(Sticker.ColumnTagModelID+ " = ? and " + Sticker.ColumnTagModelAccountID + " = ?", tagID, accountID).
					First(&tagModel)
				if dbResult.Error != nil {
					return dbResult.Error
				}
				// database operation
				dbResult = tx.Create(&Sticker.StickerTagModel{
					StickerID: stickerBasicModel.ID,
					TagID: tagID,
				})
				if dbResult.Error != nil {
					return dbResult.Error
				}
			}
		}

		return nil
	})

	if errTransaction != nil {
		return StickerResponse.SimpleResponse{
			Code: 400,
			Message: "Params Error",
		}
	}

	return StickerResponse.CreateSuccessSimpleResponse()
}

