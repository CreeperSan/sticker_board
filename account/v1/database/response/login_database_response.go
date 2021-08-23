package Account

type LoginDatabaseResponse struct {
	Code          int
	Message       string
	Token         string
	EffectiveTime int64
	UID           uint
}

