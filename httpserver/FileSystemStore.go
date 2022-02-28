package main

import (
	"encoding/json"
	"io"
)

type FileSystemStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	f.rewindStore()
	league, _ := NewLeague(f.database)
	for _, player := range league {
		if player.Name == name {
			return player.Wins
		}
	}
	return 0
}

func (f *FileSystemStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)
	if player != nil {
		player.Wins++
	}
	f.rewindStore()
	_ = json.NewEncoder(f.database).Encode(league)
}

func (f *FileSystemStore) GetLeague() League {
	f.rewindStore()
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemStore) rewindStore() {
	_, _ = f.database.Seek(0, 0)
}
