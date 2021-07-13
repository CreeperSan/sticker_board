package StickerBoardAccount

type ActionResponse struct {
	isSuccess bool
	message string
}

func CreateActionSuccessResponse() ActionResponse {
	response := ActionResponse{}
	response.isSuccess = true
	response.message = "Operation successful."
	return response
}

func CreateActionFailResponse(message string) ActionResponse {
	response := ActionResponse{}
	response.isSuccess = false
	response.message = message
	return response
}
