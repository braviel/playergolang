package playerserver

//Player struct to store data from JSON
type Player struct {
	Name string
	Wins int
}

//PlayerStore interface
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}
