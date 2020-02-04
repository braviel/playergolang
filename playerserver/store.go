package playerserver

//PlayerStore interface
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}
