package StickerDatabase

const TableNameStickerBasicModel = "sticker_basic"

const ColumnStickerBasicModelID = "_id"
const ColumnStickerBasicModelType = "type"
const ColumnStickerBasicModelAccountID = "account_id"
const ColumnStickerBasicModelStar = "star"
const ColumnStickerBasicModelPinned = "is_pinned"
const ColumnStickerBasicModelStatus = "status"
const ColumnStickerBasicModelTitle = "title"
const ColumnStickerBasicModelBackground = "background"
const ColumnStickerBasicModelCreateTime = "create_time"
const ColumnStickerBasicModelUpdateTime = "update_time"
const ColumnStickerBasicModelSearchText = "search_text"
const ColumnStickerBasicModelExtra = "extra"

type StickerBasicModel struct {
	ID         uint    `gorm:"column:_id; unique; primary_key; autoincrement;"`
	Type       int     `gorm:"column:type; not null;"`
	AccountID  uint    `gorm:"column:account_id; not null;"`
	Star       int     `gorm:"column:star; not nul; default 0;"`
	Pinned     int     `gorm:"column:is_pinned; not null; default 0;"`
	Status     int     `gorm:"column:status; not null; default 0;"`
	Title      string  `gorm:"column:title; not null;"`
	Background string  `gorm:"column:background; not null;"`
	CreateTime int     `gorm:"column:create_time; not null; autoCreateTime:milli;"`
	UpdateTime int     `gorm:"column:update_time; not null; autoUpdateTime:milli;"`
	SearchText string  `gorm:"column:search_text;"`
	Extra      string  `gorm:"column:extra;"`
	Sort       float64 `gorm:"column:sort; default 32768; not null;"`
}

func (StickerBasicModel) TableName() string {
	return TableNameStickerBasicModel
}
