package AccountModule

type AccountTokenDatabaseModel struct {
	ID          string
	Token       string
	AccountID   string
	UpdateTime  int64
	Platform    int
	Brand       string
	DeviceName  string
	MachineCode string
	ExpireTime  int64
}