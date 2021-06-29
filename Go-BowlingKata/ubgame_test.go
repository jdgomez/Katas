package main

import (
	"fmt"
	"testing"
)

func rollMany(game *UbGame, n int, pins int) {
	for i := 0; i < n; i++ {
		game.roll(pins)
	}
}
func rollSpare(game *UbGame) {
	game.roll(5)
	game.roll(5)
}

func rollStrike(game *UbGame) {
	game.roll(10)
}

func TestDummyUbGame(t *testing.T) {
	game := UbGame{}
	rollMany(&game, 20, 0)
	if game.score() != 0 {
		t.Errorf("Error: score doesn't match")
	}
}

func TestAllOnes(t *testing.T) {
	game := UbGame{}
	rollMany(&game, 20, 1)
	if game.score() != 20 {
		t.Errorf("Error: score doesn't match")
	}
}

func TestOneSpare(t *testing.T) {
	game := UbGame{}
	rollSpare(&game)
	game.roll(3)
	rollMany(&game, 17, 0)
	if game.score() != 16 {
		t.Errorf("Error: score doesn't match")
	}
}

func TestOneStrike(t *testing.T) {
	game := UbGame{}
	rollStrike(&game)
	game.roll(3)
	game.roll(4)
	rollMany(&game, 16, 0)
	if game.score() != 24 {
		t.Errorf("Error: score doesn't match")
	}
}

func TestPerfectFame(t *testing.T) {
	game := UbGame{}
	rollMany(&game, 12, 10)
	if game.score() != 300 {
		t.Errorf("Error: score doesn't match")
	}
}

func TestEdgeCaseLastSpare(t *testing.T) {
	game := UbGame{}
	rollMany(&game, 18, 0)
	rollSpare(&game)
	game.roll(3)
	if game.score() != 16 {
		fmt.Println(game.score())
		t.Errorf("Error: score doesn't match")
	}
}

func TestEdgeCaseLastStrike(t *testing.T) {
	game := UbGame{}
	rollMany(&game, 18, 0)
	rollStrike(&game)
	game.roll(3)
	game.roll(4)
	if game.score() != 24 {
		fmt.Println(game.score())
		t.Errorf("Error: score doesn't match")
	}
}