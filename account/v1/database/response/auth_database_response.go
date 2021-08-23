package Account

type AuthDatabaseResponse struct {
	Code                  int
	Message               string
	UpdateTime            int64
	ExpireTimeMilliSecond int64
	AccountID             uint
}
