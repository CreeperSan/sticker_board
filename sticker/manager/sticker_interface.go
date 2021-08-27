package StickerModule

import StickerModuleResponse "sticker_board/sticker/manager/response"

type StickerInterface interface {

	Initialize()


	// Sticker Tag

	CreateTag(accountID string, name string, icon string, color int) StickerModuleResponse.StickerTagResponse

	DeleteTag(accountID string, tagID string) StickerModuleResponse.StickerResponse

	FindAllTag(accountID string) StickerModuleResponse.StickerTagArrayResponse


	// Sticker Category

	CreateCategory(accountID string, parentCategoryID string, name string, icon string, color int) StickerModuleResponse.StickerCategoryResponse

	DeleteCategory(accountID string, categoryID string) StickerModuleResponse.StickerResponse

	FindAllCategory(accountID string) StickerModuleResponse.StickerCategoryArrayResponse


	// Sticker

	CreatePlainTextSticker()

	DeleteSticker()

	FindSticker()

}


