package solution

import (
	"strconv"
	"strings"
)

type point struct {
	X int
	Y int
}

func (s Solution) Day9Part1(fn string) int {
	ln := toLines(fn)

	head := point{}
	tail := point{}

	visited := make(map[point]bool)

	for _, l := range ln {
		split := strings.Split(l, " ")
		direciton := split[0]
		steps, _ := strconv.Atoi(split[1])
		visited[tail] = true

		for i := 0; i < steps; i++ {
			switch direciton {
			case "R":
				head.X++
			case "L":
				head.X--
			case "U":
				head.Y++
			case "D":
				head.Y--
			}

			// Only when head is in these relative positions, does tail need to move

			// #HHH#
			// H###H
			// H#T#H
			// H###H
			// #HHH#
			delta := point{
				X: head.X - tail.X,
				Y: head.Y - tail.Y,
			}

			switch delta {
			case point{-1, 2}:
				tail.X -= 1
				tail.Y += 1
			case point{0, 2}:
				tail.Y += 1
			case point{1, 2}:
				tail.X += 1
				tail.Y += 1
			case point{-2, 1}:
				tail.X -= 1
				tail.Y += 1
			case point{2, 1}:
				tail.X += 1
				tail.Y += 1
			case point{-2, 0}:
				tail.X--
			case point{2, 0}:
				tail.X++
			case point{-2, -1}:
				tail.X--
				tail.Y--
			case point{2, -1}:
				tail.X++
				tail.Y--
			case point{-1, -2}:
				tail.X--
				tail.Y--
			case point{0, -2}:
				tail.Y--
			case point{1, -2}:
				tail.X++
				tail.Y--
			}

			visited[tail] = true
		}
	}

	return len(visited)
}

func (s Solution) Day9Part2(fn string) int {
	ln := toLines(fn)

	points := []point{}
	for i := 0; i < 10; i++ {
		points = append(points, point{})
	}

	visited := make(map[point]bool)

	for _, l := range ln {
		split := strings.Split(l, " ")
		direciton := split[0]
		steps, _ := strconv.Atoi(split[1])
		visited[points[9]] = true

		for i := 0; i < steps; i++ {
			switch direciton {
			case "R":
				points[0].X++
			case "L":
				points[0].X--
			case "U":
				points[0].Y++
			case "D":
				points[0].Y--
			}

			for idx := 1; idx < 10; idx++ {
				delta := point{
					X: points[idx-1].X - points[idx].X,
					Y: points[idx-1].Y - points[idx].Y,
				}

				// HHHHH
				// H###H
				// H#T#H
				// H###H
				// HHHHH
				switch delta {
				case point{-2, 2}:
					points[idx].X--
					points[idx].Y++
				case point{-1, 2}:
					points[idx].X -= 1
					points[idx].Y += 1
				case point{0, 2}:
					points[idx].Y += 1
				case point{1, 2}:
					points[idx].X += 1
					points[idx].Y += 1
				case point{2, 2}:
					points[idx].X += 1
					points[idx].Y += 1
				case point{-2, 1}:
					points[idx].X -= 1
					points[idx].Y += 1
				case point{2, 1}:
					points[idx].X += 1
					points[idx].Y += 1
				case point{-2, 0}:
					points[idx].X--
				case point{2, 0}:
					points[idx].X++
				case point{-2, -1}:
					points[idx].X--
					points[idx].Y--
				case point{2, -1}:
					points[idx].X++
					points[idx].Y--
				case point{-2, -2}:
					points[idx].X--
					points[idx].Y--
				case point{-1, -2}:
					points[idx].X--
					points[idx].Y--
				case point{0, -2}:
					points[idx].Y--
				case point{1, -2}:
					points[idx].X++
					points[idx].Y--
				case point{2, -2}:
					points[idx].X++
					points[idx].Y--
				}
			}

			visited[points[9]] = true
		}
	}

	return len(visited)
}
