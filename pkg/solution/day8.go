package solution

import (
	"strconv"
)

func (s Solution) Day8Part1(fn string) int {
	lines := toLines(fn)

	matrix := [][]int{}
	visible := [][]bool{}

	for i, l := range lines {
		matrix = append(matrix, []int{})
		visible = append(visible, []bool{})
		for _, c := range l {
			num, _ := strconv.Atoi(string(c))
			matrix[i] = append(matrix[i], num)
			visible[i] = append(visible[i], false)
		}
	}

	for i := 0; i < len(matrix); i++ {
		curMax := 0
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 || j == 0 || i == len(matrix)-1 || j == len(matrix[0])-1 {
				visible[i][j] = true
			} else if curMax < matrix[i][j] {
				visible[i][j] = true
			}
			curMax = max(curMax, matrix[i][j])
		}

		curMax = 0
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			if i == 0 || j == 0 || i == len(matrix)-1 || j == len(matrix[0])-1 {
				visible[i][j] = true
			} else if curMax < matrix[i][j] {
				visible[i][j] = true
			}
			curMax = max(curMax, matrix[i][j])
		}
	}

	for j := 0; j < len(matrix[0])-1; j++ {
		curMax := 0
		for i := 0; i < len(matrix); i++ {
			if i == 0 || j == 0 || i == len(matrix)-1 || j == len(matrix[0])-1 {
				visible[i][j] = true
			} else if curMax < matrix[i][j] {
				visible[i][j] = true
			}
			curMax = max(curMax, matrix[i][j])
		}

		curMax = 0
		for i := len(matrix) - 1; i >= 0; i-- {
			if i == 0 || j == 0 || i == len(matrix)-1 || j == len(matrix[0])-1 {
				visible[i][j] = true
			} else if curMax < matrix[i][j] {
				visible[i][j] = true
			}
			curMax = max(curMax, matrix[i][j])
		}

	}

	sum := 0
	for i := 0; i < len(visible); i++ {
		for j := 0; j < len(visible[0]); j++ {
			if visible[i][j] {
				sum += 1
			}
		}
	}

	return sum
}

func (s Solution) Day8Part2(fn string) int {
	lines := toLines(fn)
	matrix := [][]int{}
	scores := [][]int{}

	for i, l := range lines {
		matrix = append(matrix, []int{})
		scores = append(scores, []int{})
		for _, c := range l {
			num, _ := strconv.Atoi(string(c))
			matrix[i] = append(matrix[i], num)
			scores[i] = append(scores[i], 1)
		}
	}

	for i := 0; i < len(matrix); i++ {
		heightIdxMap := make(map[int]int)
		// loop through row, left to right
		for j := 0; j < len(matrix[0]); j++ {
			heightIdx, found := heightIdxMap[matrix[i][j]]
			if found {
				scores[i][j] = scores[i][j] * abs(heightIdx-j)
			} else {
				scores[i][j] = scores[i][j] * abs(j-0)
			}

			for k := matrix[i][j]; k >= 0; k-- {
				heightIdxMap[k] = j
			}
		}

		heightIdxMap = make(map[int]int)
		// loop through row, right to left
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			heightIdx, found := heightIdxMap[matrix[i][j]]
			if found {
				scores[i][j] = scores[i][j] * abs(heightIdx-j)
			} else {
				scores[i][j] = scores[i][j] * abs(len(matrix[0])-1-j)
			}

			for k := matrix[i][j]; k >= 0; k-- {
				heightIdxMap[k] = j
			}
		}
	}

	// fmt.Println(scores)

	for j := 0; j < len(matrix[0]); j++ {
		heightIdxMap := make(map[int]int)
		// loop through column, top to down
		for i := 0; i < len(matrix); i++ {
			heightIdx, found := heightIdxMap[matrix[i][j]]
			if found {
				scores[i][j] = scores[i][j] * abs(heightIdx-i)
			} else {
				scores[i][j] = scores[i][j] * abs(i-0)
			}

			for k := matrix[i][j]; k >= 0; k-- {
				heightIdxMap[k] = i
			}
		}

		heightIdxMap = make(map[int]int)
		// loop through column, down to top
		for i := len(matrix) - 1; i >= 0; i-- {
			heightIdx, found := heightIdxMap[matrix[i][j]]
			if found {
				scores[i][j] = scores[i][j] * abs(heightIdx-i)
			} else {
				scores[i][j] = scores[i][j] * abs(len(matrix)-1-i)
			}

			for k := matrix[i][j]; k >= 0; k-- {
				heightIdxMap[k] = i
			}
		}
	}

	ret := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			ret = max(ret, scores[i][j])
		}
	}

	return ret
}
