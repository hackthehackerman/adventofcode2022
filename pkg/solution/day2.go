package solution

import "strings"

func (s Solution) Day2Part1(fn string) int {
	lines := toLines(fn)

	score := 0
	for _, l := range lines {
		split := strings.Split(l, " ")

		opp := int(split[0][0]) - int('A') + 1
		me := int(split[1][0]) - int('X') + 1

		score += me
		if opp == me {
			score += 3
		} else if me-opp == 1 || me-opp == -2 {
			score += 6
		}
	}
	return score
}

func (s Solution) Day2Part2(fn string) int {
	lines := toLines(fn)

	score := 0
	for _, l := range lines {
		split := strings.Split(l, " ")

		opp := int(split[0][0]) - int('A') + 1
		result := int(split[1][0]) - int('X') + 1

		if result == 1 {
			score += 0
			if opp == 1 {
				score += 3
			} else {
				score += (opp - 1)
			}
		} else if result == 2 {
			score += 3
			score += opp
		} else if result == 3 {
			score += 6
			if opp == 3 {
				score += 1
			} else {
				score += (opp + 1)
			}
		}
	}
	return score
}
