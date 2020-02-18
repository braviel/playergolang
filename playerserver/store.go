package playerserver

//Player struct to store data from JSON
type Player struct {
	Name string
	Wins int
}

//League is a list of Player
type League []Player

//Find find a player in league
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

//PlayerStore interface
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}
