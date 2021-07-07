package main

type light struct {
	on bool
}

func (l *light) turnOn() {
	l.on = true
}

func (l *light) turnOff() {
	l.on = false
}

func (l *light) toggle() {
	l.on = !l.on
}