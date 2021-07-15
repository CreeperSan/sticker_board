package StickerDatabase

const TableNameStickerPlainTextModel = "sticker_type_plain_text"

type StickerPlainTextModel struct {
	StickerID uint   `gorm:"column:sticker_id; unique; primary_key;"`
	Text      string `gorm:"column:text; not null;"`
	SearchText string `gorm:"column:text; not null;"`
}

func (StickerPlainTextModel) TableName() string {
	return TableNameStickerPlainTextModel
}
