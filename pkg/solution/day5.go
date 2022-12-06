package solution

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Letter string
	Next   *Node
}

func debugNodes(nodes []*Node) {
	for _, n := range nodes {
		l := ""
		next := n
		for {
			if next != nil {
				l = l + " " + next.Letter + " "
				next = next.Next
			} else {
				break
			}
		}
		fmt.Println(l)
	}
}

func (s Solution) Day5Part1(fn string) string {
	lines := toLines(fn)

	numStack := 9
	numStackButtomRow := 7

	nodes := []*Node{}
	for i := 0; i < numStack; i++ {
		nodes = append(nodes, nil)
	}

	// build stack in linked list
	for i := 0; i <= numStackButtomRow; i++ {
		l := lines[i]
		for i := 0; i < numStack; i++ {
			letter := l[i*4+1 : i*4+2]
			if strings.TrimSpace(letter) != "" {
				if nodes[i] == nil {
					nodes[i] = &Node{letter, nil}
				} else {
					next := nodes[i]
					for {
						if next.Next == nil {
							next.Next = &Node{letter, nil}
							break
						} else {
							next = next.Next
						}
					}
				}
			}
		}
	}

	// move crate
	for i := numStackButtomRow + 3; i < len(lines); i++ {
		l := lines[i]
		split := strings.Split(l, " ")
		moves, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])

		from = from - 1
		to = to - 1

		for j := 0; j < moves; j++ {
			tmp := nodes[from].Next
			nodes[from].Next = nodes[to]
			nodes[to] = nodes[from]
			nodes[from] = tmp
		}
	}

	ret := ""
	for _, n := range nodes {
		ret += n.Letter
	}

	return ret
}

func (s Solution) Day5Part2(fn string) string {
	lines := toLines(fn)

	numStack := 9
	numStackButtomRow := 7

	nodes := []*Node{}
	for i := 0; i < numStack; i++ {
		nodes = append(nodes, nil)
	}

	// build stack in linked list
	for i := 0; i <= numStackButtomRow; i++ {
		l := lines[i]
		for i := 0; i < numStack; i++ {
			letter := l[i*4+1 : i*4+2]
			if strings.TrimSpace(letter) != "" {
				if nodes[i] == nil {
					nodes[i] = &Node{letter, nil}
				} else {
					next := nodes[i]
					for {
						if next.Next == nil {
							next.Next = &Node{letter, nil}
							break
						} else {
							next = next.Next
						}
					}
				}
			}
		}
	}

	// move crate
	for i := numStackButtomRow + 3; i < len(lines); i++ {
		l := lines[i]
		split := strings.Split(l, " ")
		moves, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])

		from = from - 1
		to = to - 1

		end := nodes[from]
		for j := 0; j < moves-1; j++ {
			end = end.Next
		}
		tmp := end.Next
		end.Next = nodes[to]
		nodes[to] = nodes[from]
		nodes[from] = tmp
	}

	ret := ""
	for _, n := range nodes {
		ret += n.Letter
	}

	return ret
}
