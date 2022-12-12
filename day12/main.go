package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

type (
	Graph struct {
		nodes []node
	}
	node struct {
		dist     int
		val      byte
		children []int
	}
)

var gS int
var gE int

func buildGraph(ss []string) []node {
	var graph []node
	for y := 0; y < len(ss); y++ {
		var tmp byte
		for x := 0; x < len(ss[y]); x++ {
			var n node
			tmp = ss[y][x]
			if ss[y][x] == 'S' {
				gS = y*(len(ss[y])) + x
				tmp = 'a'
			} else if ss[y][x] == 'E' {
				gE = y*(len(ss[y])) + x
				tmp = 'z'
			}
			n.val = tmp
			if x-1 >= 0 && ss[y][x-1]-1 <= tmp {
				n.children = append(n.children, y*(len(ss[y]))+x-1)
			}
			if x+1 < len(ss[y]) && ss[y][x+1]-1 <= tmp {
				n.children = append(n.children, y*(len(ss[y]))+x+1)
			}
			if y-1 >= 0 && ss[y-1][x]-1 <= tmp {
				n.children = append(n.children, (y-1)*(len(ss[y]))+x)
			}
			if y+1 < len(ss) && ss[y+1][x]-1 <= tmp {
				n.children = append(n.children, (y+1)*(len(ss[y]))+x)
			}
			graph = append(graph, n)
		}
	}
	return graph
}

func (g *Graph) getShortestPath(source, dest int) int {
	var visited = make([]int, len(g.nodes))
	queue := []int{source}
	g.nodes[source].dist = 1
	for ; len(queue) > 0; queue = queue[1:] {
		for _, v := range g.nodes[queue[0]].children {
			if visited[v] == 0 {
				visited[v] = 1
				queue = append(queue, v)
				g.nodes[v].dist = g.nodes[queue[0]].dist + 1
			}
		}
	}
	return g.nodes[dest].dist + 1
}

func ex1(ss []string) int {
	var g Graph
	g.nodes = buildGraph(ss)
	return g.getShortestPath(gS, gE)
}

func (g *Graph) getShortestA(dest int) int {
	smallest := math.MaxInt
	for i := 0; i < len(g.nodes); i++ {
		if g.nodes[i].val == 'a' {
			tmp := g.getShortestPath(i, dest)
			if tmp < smallest {
				smallest = tmp
			}
		}
	}
	return smallest
}
func ex2(ss []string) int {
	var g Graph
	g.nodes = buildGraph(ss)
	return g.getShortestA(gE)
}

func main() {
	f, _ := os.Open("input.txt")
	data, _ := io.ReadAll(f)
	ss := strings.Split(string(data), "\n")

	t := time.Now()
	fmt.Println(ex1(ss))
	fmt.Println(time.Since(t))

	t = time.Now()
	fmt.Println(ex2(ss))
	fmt.Println(time.Since(t))
}
