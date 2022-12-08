package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func checkLeft(ss []string, y, x int) bool {
	for x2 := x - 1; x2 >= 0; x2-- {
		if ss[y][x2] >= ss[y][x] {
			return false
		}
	}
	return true
}
func checkRight(ss []string, y, x int) bool {
	for x2 := x + 1; x2 < len(ss[y]); x2++ {
		if ss[y][x2] >= ss[y][x] {
			return checkLeft(ss, y, x)
		}
	}
	return true
}
func checkBottom(ss []string, y, x int) bool {
	for y2 := y + 1; y2 < len(ss); y2++ {
		if ss[y2][x] >= ss[y][x] {
			return checkRight(ss, y, x)
		}
	}
	return true
}
func checkTop(ss []string, y, x int) bool {
	for y2 := y - 1; y2 >= 0; y2-- {
		if ss[y2][x] >= ss[y][x] {
			return checkBottom(ss, y, x)
		}
	}
	return true
}
func isVisible(ss []string, y, x int) bool {
	return checkTop(ss, y, x)
}

// Speed oriented algorithm
func ex1(ss []string) int {
	sum := 0
	for y := 1; y < len(ss)-1; y++ {
		for x := 1; x < len(ss[y])-1; x++ {
			if isVisible(ss, y, x) {
				sum++
			}
		}
	}
	return sum + len(ss)*2 + len(ss[0])*2 - 4
}

func scoreLeft(ss []string, y, x, sum int) int {
	for x2 := x - 1; x2 >= 0; x2-- {
		if ss[y][x2] >= ss[y][x] {
			return sum * (x - x2)
		}
	}
	return sum * x
}
func scoreRight(ss []string, y, x, sum int) int {
	for x2 := x + 1; x2 < len(ss[y]); x2++ {
		if ss[y][x2] >= ss[y][x] {
			return scoreLeft(ss, y, x, sum*(x2-x))
		}
	}
	return scoreLeft(ss, y, x, sum*(len(ss[0])-1-x))
}
func scoreBottom(ss []string, y, x, sum int) int {
	for y2 := y + 1; y2 < len(ss); y2++ {
		if ss[y2][x] >= ss[y][x] {
			return scoreRight(ss, y, x, sum*(y2-y))
		}
	}
	return scoreRight(ss, y, x, sum*(len(ss)-1-y))
}
func scoreTop(ss []string, y, x int) int {
	for y2 := y - 1; y2 >= 0; y2-- {
		if ss[y2][x] >= ss[y][x] {
			return scoreBottom(ss, y, x, y-y2)
		}
	}
	return scoreBottom(ss, y, x, y-1)
}

func getScore(ss []string, y, x int) int {
	return scoreTop(ss, y, x)
}

func ex2(ss []string) int {
	best := 0
	for y := 0; y < len(ss); y++ {
		for x := 0; x < len(ss[y]); x++ {
			if score := getScore(ss, y, x); score > best {
				best = score
			}
		}
	}
	return best
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
