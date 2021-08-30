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

	CreatePlainTextSticker(
		accountID string,
		star int,
		isPinned bool,
		stickerStatus int,
		title string,
		background string,
		tagIDs []string,
		categoryID string,
		text string,
	) StickerModuleResponse.StickerSingleResponse

	CreatePlainImageSticker(
		accountID string,
		star int,
		isPinned bool,
		stickerStatus int,
		title string,
		background string,
		tagIDs []string,
		categoryID string,
		imageUrl string,
		imageDescription string,
	) StickerModuleResponse.StickerSingleResponse

	CreatePlainSoundSticker(
		accountID string,
		star int,
		isPinned bool,
		stickerStatus int,
		title string,
		background string,
		tagIDs []string,
		categoryID string,
		soundUrl string,
		soundDescription string,
		soundDuration int,
	) StickerModuleResponse.StickerSingleResponse

	DeleteSticker(accountID string, stickerID string) StickerModuleResponse.StickerResponse

	FindSticker(accountID string, page int, pageSize int) StickerModuleResponse.StickerArrayResponse

}


