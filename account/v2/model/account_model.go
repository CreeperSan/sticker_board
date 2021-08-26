package AccountV2

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AccountModel Type : Customer type. default is 1
type AccountModel struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Account      string             `bson:"account"`
	Password     string             `bson:"password"`
	Username     string             `bson:"username"`
	RegisterTime int64              `bson:"register_time"`
	Avatar       string             `bson:"avatar"`
	Email        string             `bson:"email"`
	Type         int                `bson:"type"`
}
