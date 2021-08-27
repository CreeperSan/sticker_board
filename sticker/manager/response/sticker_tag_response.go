package StickerModuleResponse

import StickerModuleModel "sticker_board/sticker/manager/model"

type StickerTagResponse struct {
	StickerResponse
	Tag StickerModuleModel.TagModel
}
