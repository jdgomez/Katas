package main

import (
	"testing"
)

func TestCellDiesOfLoneliness(t *testing.T) {

	c := cell{isAlive: true}

	c.evolve(0)
	if c.isAlive {
		t.Errorf("Error: cell doesn't die of lonelyness")
	}

	c.evolve(1)
	if c.isAlive {
		t.Errorf("Error: cell doesn't die of lonelyness")
	}

}

func TestCellLiveWithCompany(t *testing.T) {

	c := cell{isAlive: true}

	c.evolve(2)
	if !c.isAlive {
		t.Errorf("Error: cell doesn't survives with company")
	}

	c.evolve(3)
	if !c.isAlive {
		t.Errorf("Error: cell doesn't survives with company")
	}
}

func TestCellDiesWhenOvercrowded(t *testing.T) {

	c := cell{isAlive: true}

	c.evolve(4)
	if c.isAlive {
		t.Errorf("Error: cell doesn't die overcrowded")
	}

	c.evolve(5)
	if c.isAlive {
		t.Errorf("Error: cell doesn't die overcrowded")
	}
}

func TestCellResurrectsWithThreeNeighbors(t *testing.T) {

	c := cell{isAlive: false}

	c.evolve(2)
	if c.isAlive {
		t.Errorf("Error: cell resurrects with two neighbors")
	}

	c.evolve(3)
	if !c.isAlive {
		t.Errorf("Error: cell doesn't resurrect")
	}
}
