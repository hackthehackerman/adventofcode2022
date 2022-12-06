package solution

import (
	"bufio"
	"os"
)

func toLines(fn string) (r []string) {
	file, _ := os.Open(fn)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}

	return
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
