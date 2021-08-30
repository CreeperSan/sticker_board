package StickerModuleModel

type StickerPlainSoundModel struct {
	StickerBasicModel
	Url         string `json:"url"`
	Duration    int    `json:"duration"` // Millisecond
	Description string `json:"description"`
}
