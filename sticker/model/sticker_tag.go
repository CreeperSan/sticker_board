package StickerDatabase

const TableNameTagModel = "tag"

const ColumnTagModelID = "_id"
const ColumnTagModelAccountID = "account_id"
const ColumnTagModelCreateTime = "create_time"
const ColumnTagModelUpdateTime = "update_time"
const ColumnTagModelName = "name"
const ColumnTagModelIcon = "icon"
const ColumnTagModelColor = "color"
const ColumnTagModelExtra = "extra"

type TagModel struct {
	ID         uint   `gorm:"column:_id; primaryKey; autoIncrement; not null;"`
	AccountID  uint   `gorm:"column:account_id; not null;"`
	CreateTime int    `gorm:"column:create_time; not null; autoCreateTime:milli;"`
	UpdateTime int    `gorm:"column:update_time; not null; autoUpdateTime:milli;"`
	Name       string `gorm:"column:name;"`
	Icon       string `gorm:"column:icon;"`
	Color      int    `gorm:"column:color;"`
	Extra      string `gorm:"column:extra;"`
	Sort       int    `gorm:"column:sort;"`
}

func (TagModel) TableName() string {
	return TableNameTagModel
}