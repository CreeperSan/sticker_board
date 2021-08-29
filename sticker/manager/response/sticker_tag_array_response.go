package StickerModuleResponse

import (
	StickerV2Model "sticker_board/sticker/v2/model"
)

type StickerTagArrayResponse struct {
	StickerResponse
	Tags []StickerV2Model.TagDatabaseModel
}
