package solution

import (
	"sort"
	"strconv"
)

type Solution struct {
}

func (s Solution) Day1Part1(fn string) int {
	lines := toLines(fn)

	curMax := 0
	tmp := 0
	for i, v := range lines {
		if v == "" || i == len(lines)-1 {
			curMax = max(curMax, tmp)
			tmp = 0
			continue
		}
		v, _ := strconv.Atoi(v)
		tmp += v
	}
	return curMax
}

func (s Solution) Day1Part2(fn string) int {
	lines := toLines(fn)

	calories := []int{}
	sumCalorie := 0
	for i, v := range lines {
		if v == "" || i == len(lines)-1 {
			calories = append(calories, sumCalorie)
			sumCalorie = 0
			continue
		}
		v, _ := strconv.Atoi(v)
		sumCalorie += v
	}
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	return calories[0] + calories[1] + calories[2]
}
