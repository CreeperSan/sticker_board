package AccountModule

import AccountModule "sticker_board/account/manager/model"

type AccountResponse struct {
	AccountBasicResponse
	Account AccountModule.AccountModel
}
