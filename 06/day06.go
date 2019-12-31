package main

import (
	"fmt"
	"github.com/subchord/aoc-2019/util"
	"strings"
)

type Node struct {
	Name     string
	Children []*Node
}

func main() {
	lines := util.ReadLines("inp")
	com := getGraph(lines)
	fmt.Println(part1(com))
	fmt.Println(part2(com))
}

func getGraph(lines []string) *Node {
	nodes := map[string]*Node{}
	for _, line := range lines {
		parts := strings.Split(line, ")")
		a, b := parts[0], parts[1]

		nodea, oka := nodes[a]
		nodeb, okb := nodes[b]

		if !oka {
			nodea = &Node{
				Name:     a,
				Children: []*Node{},
			}
			nodes[a] = nodea
		}

		if !okb {
			nodeb = &Node{
				Name:     b,
				Children: []*Node{},
			}
			nodes[b] = nodeb
		}

		nodes[a].Children = append(nodes[a].Children, nodeb)
	}
	return nodes["COM"]
}

func walkNodes(n *Node, depth int) int {
	s := 0
	for i := range n.Children {
		s += depth + 1
		s += walkNodes(n.Children[i], depth+1)
	}
	return s
}

func part1(n *Node) int {
	return walkNodes(n, 0)
}

func getPath(n *Node, target string) ([]string, bool) {
	path := []string{}
	found := false
	for i := range n.Children {
		if n.Children[i].Name == target {
			found = true
			break
		}
		p, b := getPath(n.Children[i], target)
		if b {
			path = append(path, n.Children[i].Name)
			path = append(path, p...)
			found = b
		}
	}
	return path, found
}

func part2(n *Node) int {
	san, _ := getPath(n, "SAN")
	you, _ := getPath(n, "YOU")

	lastIdx := 0
	for i, s := range you {
		if san[i] != s {
			lastIdx = i
			break
		}
	}

	return len(san[lastIdx:]) + len(you[lastIdx:])
}
