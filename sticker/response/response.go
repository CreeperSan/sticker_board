package StickerResponse

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
