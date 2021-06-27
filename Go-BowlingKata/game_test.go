package main

import (
	"fmt"
	"testing"
)

var rollTests = []struct {
	in  []int
	out int
}{
	{[]int{5, 3, 5, 3, 5, 3}, 24},
	{[]int{5, 5, 5}, 20},
	{[]int{10, 5, 5}, 30},
	{[]int{10, 5}, 20},
	{[]int{5, 3, 10, 3, 5, 3}, 37},
	{[]int{5, 3, 10, 5, 5, 3}, 44},
	{[]int{5, 3, 10, 5, 5, 3, 7}, 51},
	{[]int{5, 3, 10, 5, 5, 3, 7, 10}, 71},
	{[]int{5, 3, 10, 5, 5, 3, 7, 10, 10}, 91},
	{[]int{5, 3, 10, 5, 5, 3, 7, 10, 10, 4}, 103},
	{[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, 300},
	{[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 2, 5}, 286},
	{[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 2, 5, 3}, 286},
	{[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 2, 5, 3, 1, 2}, 286},
	{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 20},
	{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}, 20},
}

func TestNewGame(t *testing.T) {

	game := game{}
	if len(game.rolls) != 0 {
		t.Errorf("Error: new games doesn't start empty")
	}
}

func TestRoll(t *testing.T) {

	game := game{}
	if len(game.rolls) != 0 {
		t.Errorf("Error: ball not rolled yet")
	}

	game.roll(5)
	if len(game.rolls) != 1 {
		t.Errorf("Error: number of rolls doesn't match")
	}

	if game.rolls[0].pinsDown != 5 {
		t.Errorf("Error: pins down doen'st match")
	}
}

/*func TestScore(t *testing.T) {

	game := game{}

	if game.score() != 0 {
		t.Errorf("Error: score must be 0 befores first roll")
	}

	game.roll(5)
	if game.score() != 5 {
		t.Errorf("Error: score doesn't match")
	}
}*/

func TestSampleRoles(t *testing.T) {

	for i, sample := range rollTests {
		game := game{}
		for _, pinsDown := range sample.in {
			game.roll(pinsDown)
		}
		if game.score() != sample.out {
			fmt.Println(game.score(), "!==", sample.out)
			t.Errorf("Error: score doesn't match at sample %v ", i)
		}
	}
}
