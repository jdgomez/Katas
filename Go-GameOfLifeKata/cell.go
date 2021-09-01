package main

type cell struct {
	isAlive bool
}

func (c *cell) evolve(neighbor int) {

	if neighbor < 2 || neighbor > 3 {
		c.isAlive = false
	} else {
		if (c.isAlive && 2 == neighbor) || 3 == neighbor {
			c.isAlive = true
		}
	}

}
