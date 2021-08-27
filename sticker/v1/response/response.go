package StickerResponse

import (
	StickerDatabase2 "sticker_board/sticker/v1/model"
)

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
	Data []StickerDatabase2.TagModel
}


type QueryCategoryResponse struct {
	Code int
	Message string
	Data []StickerDatabase2.CategoryModel
}
