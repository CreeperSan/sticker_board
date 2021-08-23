package AccountAPI

type AccountInitializeResponse struct {
	IsSuccess bool
	Message   string
}

type AccountRegisterResponse struct {
	IsSuccess bool
	Message   string
}

type AccountLoginResponse struct {
	IsSuccess     bool
	Message       string
	Token         string
	EffectiveTime int
	UID           int
}

type AccountAuthTokenResponse struct {
	IsSuccess     bool
	Message       string
	Token         string
	EffectiveTime int
	UID           int
}

type AccountIsExistResponse struct {
	IsSuccess bool
	IsExist   bool
}

type AccountApi interface {
	Initialize() AccountInitializeResponse

	RegisterAccount(
		account string,
		password string,
		username string,
		email string,
	) AccountRegisterResponse

	LoginAccount(
		account string,
		password string,
		platform int8,
		brand string,
		deviceName string,
		machineCode string,
	) AccountLoginResponse

	AuthToken(
		accountID uint,
		token string,
		platform int,
		brand string,
		deviceName string,
		machineCode string,
	) AccountAuthTokenResponse

	IsAccountExist(
		accountID uint,
	) bool
}
