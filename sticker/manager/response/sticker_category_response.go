package StickerModuleResponse

import StickerModuleModel "sticker_board/sticker/manager/model"

type StickerCategoryResponse struct {
	StickerResponse
	Category StickerModuleModel.CategoryModel
}
