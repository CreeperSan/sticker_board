package StickerModuleResponse

type StickerResponse struct {
	Code int
	Message string
}

func (response *StickerResponse) IsSuccess() bool {
	return response.Code == ResponseCodeSuccess
}

const ResponseCodeSuccess = 200
const ResponseCodeTodo = 0

func CreateSuccessResponse() StickerResponse {
	return StickerResponse{
		Code: ResponseCodeSuccess,
		Message: "Success",
	}
}


func CreateTodoResponse() StickerResponse {
	return StickerResponse{
		Code: ResponseCodeTodo,
		Message: "Todo",
	}
}
