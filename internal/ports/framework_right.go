package ports

type DBPort interface {
	CloseDBConn()
	AddToHistory(answer int32, operation string) error
}
