package StickerDatabase

const TableNameStickerPlainTextModel = "sticker_type_plain_text"

const ColumnStickerPlainTextModelStickerID = "sticker_id"
const ColumnStickerPlainTextModelText = "text"
const ColumnStickerPlainTextModelSearchText = "search_text"

type StickerPlainTextModel struct {
	StickerID uint   `gorm:"column:sticker_id; unique; primary_key;"`
	Text      string `gorm:"column:text; not null;"`
	SearchText string `gorm:"column:search_text; not null;"`
}

func (StickerPlainTextModel) TableName() string {
	return TableNameStickerPlainTextModel
}
