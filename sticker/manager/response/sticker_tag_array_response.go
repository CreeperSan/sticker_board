package StickerModuleResponse

import StickerModuleModel "sticker_board/sticker/manager/model"

type StickerTagArrayResponse struct {
	StickerResponse
	Tags []StickerModuleModel.TagModel
}
