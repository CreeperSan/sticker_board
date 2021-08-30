package StickerModuleResponse

type StickerArrayResponse struct {
	StickerResponse
	Stickers []interface{} // instance of sticker model
}
