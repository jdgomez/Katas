package main

type UbGame struct {
	currentRoll int
	rolls       [21]int
}

func (g *UbGame) roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll++
}

func (g UbGame) score() int {
	score := 0
	frameIndex := 0
	for frame := 0; frame < 10; frame++ {
		if frameIsStrike(g, frameIndex) {
			score += 10 + strikeBonus(g,frameIndex)
			frameIndex++
		} else if frameIsSpare(g, frameIndex) {
			score += 10 + spareBonus(g,frameIndex)
			frameIndex += 2
		} else {
			score += g.rolls[frameIndex] + g.rolls[frameIndex+1]
			frameIndex += 2
		}
	}
	return score
}

func frameIsSpare(g UbGame, frameIndex int) bool {
	return g.rolls[frameIndex]+g.rolls[frameIndex+1] == 10
}

func frameIsStrike(g UbGame, frameIndex int) bool {
	return g.rolls[frameIndex] == 10
}

func strikeBonus(g UbGame, frameIndex int) int {
	return g.rolls[frameIndex+1] + g.rolls[frameIndex+2]
}

func spareBonus(g UbGame, frameIndex int) int {
	return g.rolls[frameIndex+2]
}
