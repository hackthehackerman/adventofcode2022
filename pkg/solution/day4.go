package solution

import (
	"strconv"
	"strings"
)

func parseSeat(s string) (int, int, int, int) {
	split := strings.Split(s, ",")
	left := strings.Split(split[0], "-")
	right := strings.Split(split[1], "-")

	l1, _ := strconv.Atoi(left[0])
	r1, _ := strconv.Atoi(left[1])
	l2, _ := strconv.Atoi(right[0])
	r2, _ := strconv.Atoi(right[1])
	return l1, r1, l2, r2
}

func (s Solution) Day4Part1(fn string) int {

	lines := toLines(fn)
	cnt := 0

	for _, l := range lines {
		l1, r1, l2, r2 := parseSeat(l)
		if l1 <= l2 && r1 >= r2 {
			cnt += 1
		} else if l2 <= l1 && r2 >= r1 {
			cnt += 1
		}

	}

	return cnt
}

func (s Solution) Day4Part2(fn string) int {

	lines := toLines(fn)
	cnt := 0

	for _, l := range lines {
		l1, r1, l2, r2 := parseSeat(l)

		if r1 >= l2 && l1 <= l2 {
			cnt += 1
		} else if r2 >= l1 && l2 <= l1 {
			cnt += 1
		}
	}

	return cnt
}
