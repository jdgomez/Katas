package main

import (
	"testing"
)

func TestCreateGrid(t *testing.T) {

	g := grid{}

	if len(g.lights) != 1000  {
		t.Errorf("Error: the grid doesn`t have the expected size")
	}

	if len(g.lights[0]) != 1000  {
		t.Errorf("Error: the grid doesn`t have the expected size")
	}
}

func TestTurnOn(t *testing.T) {

	g := grid{}

	start := coordinate{0,0}
	finish := coordinate{2,2}
	g.turnOn(start,finish)

	if !g.lights[0][0].on  {
		t.Errorf("Error: the light is off")
	}
}

func TestTurnOff(t *testing.T) {

	g := grid{}

	start := coordinate{0,0}
	finish := coordinate{2,2}
	g.turnOn(start,finish)
	g.turnOff(start,finish)

	if g.lights[0][0].on  {
		t.Errorf("Error: the light is on")
	}
}

func TestToggle(t *testing.T) {

	g := grid{}

	start := coordinate{0,0}
	finish := coordinate{2,2}
	g.toggle(start,finish)
	if !g.lights[0][0].on  {
		t.Errorf("Error: the light is off")
	}

	g.toggle(start,finish)

	if g.lights[0][0].on  {
		t.Errorf("Error: the light is on")
	}
}