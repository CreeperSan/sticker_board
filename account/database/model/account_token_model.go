package StickerBoardAccount

const TableNameAccountTokenModel = "account_token"

const ColumnAccountTokenModelID = "_id"
const ColumnAccountTokenModelToken = "token"
const ColumnAccountTokenModelAccountID = "account_id"
const ColumnAccountTokenModelUpdateTime = "update_time"
const ColumnAccountTokenModelPlatform = "platform"
const ColumnAccountTokenModelBrand = "brand"
const ColumnAccountTokenModelDeviceName = "device_name"
const ColumnAccountTokenModelMachineCode = "machine_code"
const ColumnAccountTokenModelExpireTimeMilliSecond = "expire_time_millisecond"

type AccountTokenModel struct {
	ID 						uint		`gorm:"column:_id; primaryKey; autoIncrement; not null"`
	Token					string		`gorm:"column:token; unique; not null"`
	AccountID				uint		`gorm:"column:account_id; not null"`
	UpdateTime				int64		`gorm:"column:update_time; autoUpdateTime:milli; not null"`
	Platform				int8		`gorm:"column:platform; not null"`
	Brand					string		`gorm:"column:brand; not null"`
	DeviceName				string		`gorm:"column:device_name; not null"`
	MachineCode				string		`gorm:"column:machine_code; not null"`
	ExpireTimeMilliSecond	int64		`gorm:"column:expire_time_millisecond; not null"`
}

func (AccountTokenModel) TableName() string {
	return TableNameAccountTokenModel
}

