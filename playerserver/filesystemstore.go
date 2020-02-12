package playerserver

import "io"

//FileSystemPlayerStore struct store player in file system
type FileSystemPlayerStore struct {
	database io.Reader
}

//GetLeague func
func (f *FileSystemPlayerStore) GetLeague() []Player {
	league, _ := NewLeague(f.database)
	return league
}
