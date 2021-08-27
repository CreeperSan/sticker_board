package StickerDatabase

const TableNameCategoryModel = "category"

const ColumnCategoryModelID = "_id"
const ColumnCategoryModelParentID = "parent_id"
const ColumnCategoryModelAccountID = "account_id"
const ColumnCategoryModelName = "name"
const ColumnCategoryModelCreateTime = "create_time"
const ColumnCategoryModelUpdateTime = "update_time"
const ColumnCategoryModelIcon = "icon"
const ColumnCategoryModelColor = "color"
const ColumnCategoryModelExtra = "extra"

type CategoryModel struct {
	ID         uint   `gorm:"column:_id; primaryKey; autoIncrement; not null"`
	ParentID   uint   `gorm:"column:parent_id;"`
	AccountID  uint   `gorm:"column:account_id; not null;"`
	Name       string `gorm:"column:name; not null;"`
	CreateTime int    `gorm:"column:create_time; not null; autoCreateTime:milli;"`
	UpdateTime int    `gorm:"column:update_time; not null; autoUpdateTime:milli;"`
	Icon       string `gorm:"column:icon;"`
	Color      int    `gorm:"column:color;"`
	Extra      string `gorm:"column:extra;"`
	Sort       int    `gorm:"column:sort;"`
}

func (CategoryModel) TableName() string {
	return TableNameCategoryModel
}
