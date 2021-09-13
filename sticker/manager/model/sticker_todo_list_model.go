package StickerModuleModel

type StickerTodoListModel struct {
	StickerBasicModel
	Description string                     `json:"description"`
	Todos       []StickerTodoListItemModel `json:"todos"`
}

type StickerTodoListItemModel struct {
	State       int    `json:"state"` // 0 = not Finish, 1 = Finish
	Message     string `json:"message"`
	Description string `json:"description"`
}
