package main

type coordinate struct {
	x int
	y int
}
type grid struct {
	lights [1000][1000]light
}

func (g *grid) turnOn(start coordinate, finish coordinate) {
	//TODO: Can extract this loop to a method if the method returns a pointer to the light
	//      and then turning ON the returned ligth, keeping the loop method without ifs statements
	for x := start.x; x < finish.x; x++ {
		for y := start.y; y < finish.y; y++ {
			g.lights[x][y].turnOn()
		}
	}
}

func (g *grid) turnOff(start coordinate, finish coordinate) {
	for x := start.x; x < finish.x; x++ {
		for y := start.y; y < finish.y; y++ {
			g.lights[x][y].turnOff()
		}
	}
}

func (g *grid) toggle(start coordinate, finish coordinate) {
	for x := start.x; x < finish.x; x++ {
		for y := start.y; y < finish.y; y++ {
			g.lights[x][y].toggle()
		}
	}
}