package StickerV2Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CategoryDatabaseModel struct {
	ID         primitive.ObjectID `bson:"_id, omitempty"`
	AccountID  string             `bson:"account_id"`
	Name       string             `bson:"name"`
	CreateTime int64              `bson:"create_time"`
	UpdateTime int64              `bson:"update_time"`
	Icon       string             `bson:"icon"`
	Color      int                `bson:"color"`
	Sort       int                `bson:"sort"`
}
