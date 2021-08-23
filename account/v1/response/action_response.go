package StickerBoardAccount

type ActionResponse struct {
	Code    int
	Message string
}

func CreateActionSuccessResponse() ActionResponse {
	response := ActionResponse{}
	response.Code = 200
	response.Message = "Operation succeed."
	return response
}

func CreateActionFailResponse(message string) ActionResponse {
	response := ActionResponse{}
	response.Code = 400
	response.Message = message
	return response
}
