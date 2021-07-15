package StickerDatabase

const TableNameCategoryModel = "category"

type CategoryModel struct {
	ID         uint   `gorm:"column:_id; primaryKey; autoIncrement; not null"`
	ParentID   uint   `gorm:"column:parent_id;"`
	AccountID  uint   `gorm:"column:parent_id; not null;"`
	Name       string `gorm:"column:name; not null;"`
	CreateTime int    `gorm:"column:create_time; not null; autoCreateTime:milli;"`
	UpdateTime int    `gorm:"column:update_time; not null; autoUpdateTime:milli;"`
	Icon       string `gorm:"column:icon;"`
	Color      int    `gorm:"column:color;"`
	Extra      string `gorm:"column:extra;"`
}

func (CategoryModel) TableName() string {
	return TableNameCategoryModel
}
