package main

import "fmt"

type roll struct {
	throw    int
	pinsDown int
}

type game struct {
	rolls []roll
}

func (g *game) roll(pinsDown int) {
	t := len(g.rolls) + 1
	g.rolls = append(g.rolls, roll{t, pinsDown})
}

func (g game) score() int {

	var totalScore int
	firstRollInFrame := true
	var lastFrame int
	var totalRolls = len(g.rolls)
	for i, r := range g.rolls {
		if firstRollInFrame {
			lastFrame += 1
		}
		if lastFrame == 12 {
			break
		}
		if lastFrame == 11 {
			if g.rolls[i-1].pinsDown != 10 {
				break
			}
		}
		totalScore += r.pinsDown
		if !firstRollInFrame {
			if i < totalRolls-1 {
				if g.rolls[i-1].pinsDown+r.pinsDown == 10 {
					totalScore += g.rolls[i+1].pinsDown
				}
			}
		}

		if firstRollInFrame {
			if r.pinsDown == 10 {

				firstRollInFrame = !firstRollInFrame
				if i < totalRolls-1 {
					totalScore += g.rolls[i+1].pinsDown
				}
				if i < totalRolls-2 {
					totalScore += g.rolls[i+2].pinsDown
				}

			}
		}

		firstRollInFrame = !firstRollInFrame

		fmt.Println("frame:", i, "-", lastFrame)
	}

	return totalScore
}
