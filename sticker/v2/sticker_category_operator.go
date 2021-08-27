package StickerV2

import StickerModuleResponse "sticker_board/sticker/manager/response"

func (operator StickerOperator) CreateCategory (
	accountID string,
	parentCategoryID string,
	name string,
	icon string,
	color int,
) StickerModuleResponse.StickerCategoryResponse {
	return StickerModuleResponse.StickerCategoryResponse{
		StickerResponse : StickerModuleResponse.CreateTodoResponse(),
	}
}

func (operator StickerOperator) DeleteCategory (
	accountID string,
	categoryID string,
) StickerModuleResponse.StickerResponse {
	return StickerModuleResponse.CreateTodoResponse()
}

func (operator StickerOperator) FindAllCategory (
	accountID string,
) StickerModuleResponse.StickerCategoryArrayResponse {
	return StickerModuleResponse.StickerCategoryArrayResponse{
		StickerResponse : StickerModuleResponse.CreateTodoResponse(),
	}
}
