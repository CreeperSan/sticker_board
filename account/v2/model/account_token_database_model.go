package AccountV2

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccountTokenDatabaseModel struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	Token       string             `bson:"token"`
	AccountID   string             `bson:"account_id"`
	UpdateTime  int64              `bson:"update_time"`
	Platform    int                `bson:"platform"`
	Brand       string             `bson:"brand"`
	DeviceName  string             `bson:"device_name"`
	MachineCode string             `bson:"machine_code"`
	ExpireTime  int64              `bson:"expire_time"`
}
