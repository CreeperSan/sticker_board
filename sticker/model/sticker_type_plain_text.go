package StickerDatabase

const TableNameStickerPlainTextModel = "sticker_type_plain_text"

const ColumnStickerPlainTextModelStickerID = "sticker_id"
const ColumnStickerPlainTextModelText = "text"
const ColumnStickerPlainTextModelSearchText = "search_text"

type StickerPlainTextModel struct {
	ID        int    `gorm:"column:_id; primary_key; autoIncrement; not null;"`
	StickerID uint   `gorm:"column:sticker_id; unique; not null;"`
	Text      string `gorm:"column:text; not null;"`
}

func (StickerPlainTextModel) TableName() string {
	return TableNameStickerPlainTextModel
}
