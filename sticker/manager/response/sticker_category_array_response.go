package StickerModuleResponse

import StickerModuleModel "sticker_board/sticker/manager/model"

type StickerCategoryArrayResponse struct {
	StickerResponse
	Categories []StickerModuleModel.CategoryModel
}
