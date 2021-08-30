package StickerModuleModel

type StickerPlainSoundModel struct {
	StickerBasicModel
	Url    string
	Duration    int // Millisecond
	Description string
}
