package solution

func (s Solution) Day3Part1(fn string) int {
	lines := toLines(fn)

	prioritySum := 0
	for _, l := range lines {
		left := l[0 : len(l)/2]
		right := l[len(l)/2:]

		hit := make(map[rune]bool)
		for _, letter := range left {
			hit[letter] = true
		}
		for _, letter := range right {
			if hit[letter] {
				if letter >= 'a' && letter <= 'z' {
					prioritySum += (int(letter) - int('a') + 1)
				}
				if letter >= 'A' && letter <= 'Z' {
					prioritySum += (int(letter) - int('A') + 27)
				}
				break
			}
		}
	}
	return prioritySum
}

func (s Solution) Day3Part2(fn string) int {
	lines := toLines(fn)

	prioritySum := 0
	ptr := 0
	for {
		if ptr >= len(lines) {
			break
		}
		l1 := lines[ptr]
		l2 := lines[ptr+1]
		l3 := lines[ptr+2]

		cnt := make(map[rune]int)

		for _, lines := range []string{l1, l2, l3} {
			hit := make(map[rune]bool)

			for _, r := range lines {
				hit[r] = true
			}

			for r, _ := range hit {
				cnt[r] += 1
				if cnt[r] == 3 {
					if r >= 'a' && r <= 'z' {
						prioritySum += (int(r) - int('a') + 1)
					}
					if r >= 'A' && r <= 'Z' {
						prioritySum += (int(r) - int('A') + 27)
					}
					break
				}
			}
		}
		ptr += 3
	}
	return prioritySum
}
