package StickerDatabase

const TableNameStickerCategoryModel = "sticker_category"

const ColumnStickerCategoryModelID = "_id"
const ColumnStickerCategoryModelAccountID = "account_id"
const ColumnStickerCategoryModelCategoryID = "category_id"

type StickerCategoryModel struct {
	ID uint `gorm:"column:_id; primaryKey; autoIncrement; not null;"`
	CategoryID uint `gorm:"column:category_id; not null;unique;"`
	StickerID uint `gorm:"column:sticker_id; not null;"`
}

func (StickerCategoryModel) TableName() string {
	return TableNameStickerCategoryModel
}
