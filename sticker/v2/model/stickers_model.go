package StickerV2Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type StickerDatabaseModel struct {
	ID         primitive.ObjectID `bson:"_id, omitempty"`
	Type       int                `bson:"type"`
	AccountID  string             `bson:"account_id"`
	Star       int                `bson:"star"`
	IsPinned   bool               `bson:"is_pinned"`
	Status     int                `bson:"status"`
	Title      string             `bson:"title"`
	Background string             `bson:"background"`
	CreateTime int64              `bson:"create_time"`
	UpdateTime int64              `bson:"update_time"`
	SearchText string             `bson:"search_text"`
	Sort       int                `bson:"sort"`
	TagIDs     []string           `bson:"tags"`
	CategoryID string             `bson:"category"`

	// Type = Plain Text
	PlainText string `bson:"plain_text"`

	// Type = Plain Image
	PlainImageUrl         string `bson:"plain_image_url"`
	PlainImageDescription string `bson:"plain_image_description"`

	// Type = Plain Sound
	PlainSoundUrl         string `bson:"plain_sound_url"`
	PlainSoundDescription string `bson:"plain_sound_description"`
	PlainSoundDuration    int    `bson:"plain_sound_duration"`

	// Type = Todo List
	TodoListDescription string           `bson:"todo_list_description"`
	TodoListAction      []TodoListAction `bson:"todo_list_action"`
}

// For Todo List
type TodoListAction struct {
	State       int    `bson:"state"` // 0 = not Finish, 1 = Finish
	Message     string `bson:"message"`
	Description string `bson:"description"`
}
