package StickerDatabase

const TableNameStickerCategoryModel = "sticker_category"


type StickerCategoryModel struct {
	AccountID uint `gorm:"column:account_id; not null;"`
	CategoryID uint `gorm:"column:category_id; not null;unique;"`
}

func (StickerCategoryModel) TableName() string {
	return TableNameStickerCategoryModel
}
