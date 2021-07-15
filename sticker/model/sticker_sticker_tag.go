package StickerDatabase

const TableNameStickerTagModel = "sticker_tag"

type StickerTagModel struct {
	AccountID uint `gorm:"column:account_id; not null;"`
	TagID uint `gorm:"column:tag_id; not null;"`
}

func (StickerTagModel) TableName() string {
	return TableNameStickerTagModel
}
