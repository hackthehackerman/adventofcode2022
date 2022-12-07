package solution

import (
	"math"
	"strconv"
	"strings"
)

type node struct {
	IsDir   bool
	Name    string
	Size    int
	Content []*node
	Parent  *node
}

func sizeOf(n *node) int {
	if !n.IsDir {
		return n.Size
	} else {
		s := 0
		for _, c := range n.Content {
			s += sizeOf(c)
		}
		return s
	}
}

func (s Solution) Day7Part1(fn string) int {
	lines := toLines(fn)

	ls := false

	root := &node{
		IsDir:   true,
		Name:    "/",
		Size:    0,
		Content: []*node{},
		Parent:  nil,
	}
	cur := root

	for _, l := range lines {
		if l[:4] == "$ cd" {
			ls = false
			dest := strings.Split(l, "$ cd ")[1]
			if dest == ".." {
				cur = cur.Parent
			} else if dest == "/" {
				cur = root
			} else {
				for _, n := range cur.Content {
					if n.Name == dest {
						cur = n
						break
					}
				}
			}
		} else if l[:4] == "$ ls" {
			ls = true
		} else if ls {
			if l[:3] == "dir" {
				name := strings.Split(l, "dir ")[1]
				dirNode := node{
					IsDir:   true,
					Name:    name,
					Size:    0,
					Content: []*node{},
					Parent:  cur,
				}
				cur.Content = append(cur.Content, &dirNode)
			} else {
				split := strings.Split(l, " ")
				size, _ := strconv.Atoi(split[0])
				fileName := split[1]

				fileNode := node{
					IsDir:   false,
					Name:    fileName,
					Size:    size,
					Content: []*node{},
					Parent:  nil,
				}
				cur.Content = append(cur.Content, &fileNode)
			}
		}
	}

	size := 0
	var walk func(n *node)
	walk = func(n *node) {
		if n.IsDir {
			if sizeOf(n) <= 100000 {
				size += sizeOf(n)
			}
		}
		for _, c := range n.Content {
			walk(c)
		}
	}
	walk(root)

	return size
}

func (s Solution) Day7Part2(fn string) int {
	lines := toLines(fn)

	ls := false

	root := &node{
		IsDir:   true,
		Name:    "/",
		Size:    0,
		Content: []*node{},
		Parent:  nil,
	}
	cur := root

	for _, l := range lines {
		if l[:4] == "$ cd" {
			ls = false
			dest := strings.Split(l, "$ cd ")[1]
			if dest == ".." {
				cur = cur.Parent
			} else if dest == "/" {
				cur = root
			} else {
				for _, n := range cur.Content {
					if n.Name == dest {
						cur = n
						break
					}
				}
			}
		} else if l[:4] == "$ ls" {
			ls = true
		} else if ls {
			if l[:3] == "dir" {
				name := strings.Split(l, "dir ")[1]
				dirNode := node{
					IsDir:   true,
					Name:    name,
					Size:    0,
					Content: []*node{},
					Parent:  cur,
				}
				cur.Content = append(cur.Content, &dirNode)
			} else {
				split := strings.Split(l, " ")
				size, _ := strconv.Atoi(split[0])
				fileName := split[1]

				fileNode := node{
					IsDir:   false,
					Name:    fileName,
					Size:    size,
					Content: []*node{},
					Parent:  nil,
				}
				cur.Content = append(cur.Content, &fileNode)
			}
		}
	}

	freeSpace := 70000000 - sizeOf(root)
	size := math.MaxInt
	var walk func(n *node)
	walk = func(n *node) {
		if n.IsDir {
			s := sizeOf(n)
			if freeSpace+s >= 30000000 {
				size = min(size, s)
			}
		}
		for _, c := range n.Content {
			walk(c)
		}
	}
	walk(root)

	return size
}
