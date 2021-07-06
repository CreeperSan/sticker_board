package StickerBoardAccount

// AccountModel
// Used to save user's account info in database
type AccountModel struct {
	ID string
	Account string
	Password string
	UserName string
	RegisterTime int64
	Avatar string
	Email string
}
