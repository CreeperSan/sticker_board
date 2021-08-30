package StickerModuleModel

type StickerBasicModel struct {
	ID         string
	Type       int
	AccountID  string
	Star       int
	IsPinned   bool
	Status     int
	Title      string
	Background string
	CreateTime int64
	UpdateTime int64
}
