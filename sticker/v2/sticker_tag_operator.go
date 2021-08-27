package StickerV2

import StickerModuleResponse "sticker_board/sticker/manager/response"

func (operator *StickerOperator) CreateTag (accountID string, name string, icon string, color int) StickerModuleResponse.StickerTagResponse {
	return StickerModuleResponse.StickerTagResponse{
		StickerResponse : StickerModuleResponse.CreateTodoResponse(),
	}
}

func (operator *StickerOperator) DeleteTag (accountID string, tagID string) StickerModuleResponse.StickerResponse {
	return StickerModuleResponse.CreateTodoResponse()
}

func (operator *StickerOperator) FindAllTag (accountID string) StickerModuleResponse.StickerTagArrayResponse {
	return StickerModuleResponse.StickerTagArrayResponse{
		StickerResponse : StickerModuleResponse.CreateTodoResponse(),
	}
}
