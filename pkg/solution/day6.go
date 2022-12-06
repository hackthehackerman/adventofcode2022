package solution

func (s Solution) Day6Part1(fn string) int {
	l := toLines(fn)[0]

	hit := make(map[rune]int)
	lPtr := 0
	for i, char := range l {
		if idx, found := hit[char]; found && idx >= lPtr {
			lPtr = idx + 1
		} else if i-lPtr == 3 {
			return i + 1
		}
		hit[char] = i
	}
	return -1
}

func (s Solution) Day6Part2(fn string) int {
	l := toLines(fn)[0]

	hit := make(map[rune]int)
	lPtr := 0
	for i, char := range l {
		if idx, found := hit[char]; found && idx >= lPtr {
			lPtr = idx + 1
		} else if i-lPtr == 13 {
			return i + 1
		}
		hit[char] = i
	}
	return -1
}
