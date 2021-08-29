package StickerModuleResponse

type StickerResponse struct {
	Code    int
	Message string
}

func (response *StickerResponse) IsSuccess() bool {
	return response.Code == ResponseCodeSuccess
}

const ResponseCodeSuccess = 200
const ResponseCodeParamsError = 400
const ResponseCodeInternalError = 502
const ResponseCodeTodo = 0

func CreateSuccessResponse() StickerResponse {
	return StickerResponse{
		Code:    ResponseCodeSuccess,
		Message: "Success",
	}
}

func CreateTodoResponse() StickerResponse {
	return StickerResponse{
		Code:    ResponseCodeTodo,
		Message: "Todo",
	}
}

func CreateParamsErrorResponseWithMessage(message string) StickerResponse {
	return StickerResponse{
		Code:    ResponseCodeParamsError,
		Message: message,
	}
}

func CreateInternalErrorResponseWithMessage(message string) StickerResponse {
	return StickerResponse{
		Code:    ResponseCodeInternalError,
		Message: message,
	}
}
