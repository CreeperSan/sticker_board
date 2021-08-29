package StickerModuleResponse

import (
	StickerV2Model "sticker_board/sticker/v2/model"
)

type StickerCategoryArrayResponse struct {
	StickerResponse
	Categories []StickerV2Model.CategoryDatabaseModel
}
