package poker_test

import (
	"strings"
	"testing"

	poker "github.com/miketdn/learn-go-with-tests/application"
)

func TestCLI(t *testing.T) {

	t.Run("record Mike win from user input", func(t *testing.T) {
		in := strings.NewReader("Mike wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Mike")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

}
