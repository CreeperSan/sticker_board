package StickerModuleModel

type StickerPlainImageModel struct {
	StickerBasicModel
	Url         string `json:"url"`
	Description string `json:"description"`
}
