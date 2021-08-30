package StickerModuleModel

type StickerBasicModel struct {
	ID         string `json:"id"`
	Type       int    `json:"type"`
	AccountID  string `json:"-"`
	Star       int    `json:"star"`
	IsPinned   bool   `json:"is_pinned"`
	Status     int    `json:"status"`
	Title      string `json:"title"`
	Background string `json:"background"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}
