package AccountModule

import AccountModule "sticker_board/account/manager/model"

type AccountTokenResponse struct {
	AccountBasicResponse
	AccountModule.AccountTokenDatabaseModel
}
