package AccountModule

import Code "sticker_board/account/manager/const"

type AccountBasicResponse struct {
	Code int
	Message string
}

func (response AccountBasicResponse) IsSuccess() bool {
	return 200 == response.Code
}

func CreateBasicResponseTodo() AccountBasicResponse {
	return AccountBasicResponse{
		Code: Code.ResponseCodeTodo,
		Message: "Function is still in developing",
	}
}

func CreateBasicResponseParamsError() AccountBasicResponse {
	return AccountBasicResponse{
		Code: Code.ResponseCodeParamsError,
		Message: "Params error",
	}
}

func CreateBasicResponseParamsErrorWithMessage(message string) AccountBasicResponse {
	return AccountBasicResponse{
		Code: Code.ResponseCodeParamsError,
		Message: message,
	}
}

func CreateBasicResponseSuccess() AccountBasicResponse {
	return AccountBasicResponse{
		Code: Code.ResponseCodeSuccess,
		Message: "Success",
	}
}

func CreateBasicResponseDatabaseDisconnected() AccountBasicResponse {
	return AccountBasicResponse{
		Code: Code.ResponseCodeDatabaseDisconnected,
		Message: "Database disconnected",
	}
}

func CreateBasicResponseUnhandledError() AccountBasicResponse {
	return AccountBasicResponse{
		Code: Code.ResponseCodeDatabaseUnhandledError,
		Message: "Unhandled error occurred",
	}
}
