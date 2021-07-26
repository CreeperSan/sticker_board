package StickerDatabase

const TableNameStickerTagModel = "sticker_tag"

const ColumnStickerTagModelAccountID = "account_id"
const ColumnStickerTagModelTagID = "tag_id"

type StickerTagModel struct {
	ID uint `gorm:"column:_id; primaryKey; autoIncrement; not null;"`
	TagID uint `gorm:"column:tag_id; not null;"`
	StickerID uint `gorm:"column:sticker_id; not null;"`
}

func (StickerTagModel) TableName() string {
	return TableNameStickerTagModel
}
