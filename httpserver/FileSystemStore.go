package main

import (
	"encoding/json"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(database ReadWriteSeekTruncate) *FileSystemPlayerStore {
	_, _ = database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{json.NewEncoder(&Tape{database}), league}
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	_ = f.database.Encode(f.league)
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}
