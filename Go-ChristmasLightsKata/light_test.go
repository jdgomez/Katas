package main

import "testing"

func TestTurnOnLight(t *testing.T) {

	l := light{}
	l.turnOn()

	if !l.on  {
		t.Errorf("Error: light is off")
	}
}

func TestTurnOffLight(t *testing.T) {

	l := light{}
	l.turnOff()

	if l.on  {
		t.Errorf("Error: light is on")
	}
}

func TestToggleLight(t *testing.T) {

	l := light{}
	l.toggle()

	if !l.on  {
		t.Errorf("Error: light is off")
	}

	l.toggle()

	if l.on  {
		t.Errorf("Error: light is on")
	}
}