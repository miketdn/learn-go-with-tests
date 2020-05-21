package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("/league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Mike", "Wins": 10},
			{"Name": "Bob", "Wins": 24}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()
		want := []Player{
			{"Mike", 10},
			{"Bob", 24},
		}

		assertLeague(t, got, want)
	})
}
