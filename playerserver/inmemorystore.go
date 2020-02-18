package playerserver

//InMemoryPlayerStore in memory
type InMemoryPlayerStore struct {
	store map[string]int
}

//NewInMemoryPlayerStore ctor
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

//GetPlayerScore getter
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

//RecordWin setter
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

//GetLeague func
func (i *InMemoryPlayerStore) GetLeague() League {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
