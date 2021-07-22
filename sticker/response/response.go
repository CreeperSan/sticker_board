package StickerResponse

import StickerDatabase "sticker_board/sticker/model"

type SimpleResponse struct {
	Code    int
	Message string
}

func CreateSuccessSimpleResponse() SimpleResponse {
	return SimpleResponse{
		Code:    200,
		Message: "Operation succeed.",
	}
}


type QueryTagResponse struct {
	Code int
	Message string
	Data []StickerDatabase.TagModel
}


type QueryCategoryResponse struct {
	Code int
	Message string
	Data []StickerDatabase.CategoryModel
}
