package StickerBoardAccount

const TableNameAccountModel = "account"

const ColumnAccountModelID = "_id"
const ColumnAccountModelAccount = "account"
const ColumnAccountModelPassword = "password"
const ColumnAccountModelUserName = "username"
const ColumnAccountModelRegisterTime = "register_time"
const ColumnAccountModelAvatar = "avatar"
const ColumnAccountModelEmail = "email"

// AccountModel
// Used to save user's account info in database
type AccountModel struct {
	//gorm.Model
	ID 				uint		`gorm:"column:_id; primaryKey; autoIncrement; not null"`
	Account 		string		`gorm:"column:account; unique; not null"`
	Password 		string		`gorm:"column:password; not null"`
	UserName 		string		`gorm:"column:username; not null"`
	RegisterTime 	int64		`gorm:"column:register_time; autoCreateTime:milli"`
	Avatar 			string		`gorm:"column:avatar; "`
	Email 			string		`gorm:"column:email; unique; not null"`
}

func (AccountModel) TableName() string {
	return TableNameAccountModel
}
