package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkSurroundings(ys, xs, y, x int) (int, int) {
	if ys == y && xs+2 == x {
		return ys, xs + 1
	}
	if ys == y && xs-2 == x {
		return ys, xs - 1
	}
	if ys+2 == y && xs == x {
		return ys + 1, xs
	}
	if ys-2 == y && xs == x {
		return ys - 1, xs
	}

	if ys-2 == y && xs+1 == x || ys-2 == y && xs+2 == x || ys-1 == y && xs+2 == x {
		return ys - 1, xs + 1
	}
	if ys-2 == y && xs-1 == x || ys-2 == y && xs-2 == x || ys-1 == y && xs-2 == x {
		return ys - 1, xs - 1
	}
	if ys+2 == y && xs-1 == x || ys+2 == y && xs-2 == x || ys+1 == y && xs-2 == x {
		return ys + 1, xs - 1
	}
	if ys+2 == y && xs+1 == x || ys+2 == y && xs+2 == x || ys+1 == y && xs+2 == x {
		return ys + 1, xs + 1
	}
	return ys, xs
}

/*
#######
##OOO##
#O###O#
#O#T#O#
#O###O#
##OOO##
#######

*/

func ex1(ss []string) int {
	m := make(map[point]bool)
	var x, y, xs, ys int

	m[point{xs, ys}] = true
	for _, s := range ss {
		split := strings.Split(s, " ")
		nbMoves, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "U":
			for i := 0; i < nbMoves; i++ {
				y--
				ys, xs = checkSurroundings(ys, xs, y, x)
				m[point{xs, ys}] = true

			}
			break
		case "D":
			for i := 0; i < nbMoves; i++ {
				y++
				ys, xs = checkSurroundings(ys, xs, y, x)
				m[point{xs, ys}] = true
			}
			break
		case "L":
			for i := 0; i < nbMoves; i++ {
				x--
				ys, xs = checkSurroundings(ys, xs, y, x)
				m[point{xs, ys}] = true
			}
			break
		case "R":
			for i := 0; i < nbMoves; i++ {
				x++
				ys, xs = checkSurroundings(ys, xs, y, x)
				m[point{xs, ys}] = true
			}
			break
		}
	}
	tot := len(m)
	return tot
}

type point struct {
	y, x int
}

func moveRope(m map[point]bool, rope []point, y, x int) (map[point]bool, []point) {
	for i := 0; i < len(rope); i++ {
		if i == 0 {
			rope[i].y, rope[i].x = checkSurroundings(rope[i].y, rope[i].x, y, x)
		} else {
			rope[i].y, rope[i].x = checkSurroundings(rope[i].y, rope[i].x, rope[i-1].y, rope[i-1].x)
		}
		if i+1 == len(rope) {
			m[rope[i]] = true
		}
	}
	return m, rope
}

func ex2(ss []string) int {
	m := make(map[point]bool)
	var rope []point
	var x, y int
	//create rope
	for i := 0; i < 9; i++ {
		rope = append(rope, point{y, x})
	}
	m[point{x, y}] = true
	for _, s := range ss {
		split := strings.Split(s, " ")
		nbMoves, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "U":
			for i := 0; i < nbMoves; i++ {
				y--
				m, rope = moveRope(m, rope, y, x)
			}
			break
		case "D":
			for i := 0; i < nbMoves; i++ {
				y++
				m, rope = moveRope(m, rope, y, x)
			}
			break
		case "L":
			for i := 0; i < nbMoves; i++ {
				x--
				m, rope = moveRope(m, rope, y, x)
			}
			break
		case "R":
			for i := 0; i < nbMoves; i++ {
				x++
				m, rope = moveRope(m, rope, y, x)
			}
			break
		}
	}
	tot := len(m)
	return tot
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
