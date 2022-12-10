package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func incSum(sum, v, cycle, x int, offsets []int) (int, int, []int) {
	if len(offsets) > 0 && cycle == offsets[0] {
		sum += x * offsets[0]
		x += v
		offsets = offsets[1:]
	} else if len(offsets) > 0 && cycle > offsets[0] {
		sum += x * offsets[0]
		x += v
		offsets = offsets[1:]
	} else {
		x += v
	}
	return sum, x, offsets
}

func ex1(ss []string) int {
	sum := 0
	cycle := 0
	x := 1
	offsets := []int{20, 60, 100, 140, 180, 220}
	for i := 0; i < len(ss); i++ {
		if ss[i][:4] == "noop" {
			cycle++
			sum, x, offsets = incSum(sum, 0, cycle, x, offsets)
		} else {
			split := strings.Split(ss[i], " ")
			v, _ := strconv.Atoi(split[1])

			cycle += 2
			sum, x, offsets = incSum(sum, v, cycle, x, offsets)
		}
	}
	return sum
}

func newLine(cycle int, offsets []int, line []byte) ([]int, []byte) {
	if len(offsets) > 0 && cycle >= offsets[0] {
		offsets = offsets[1:]
		picture = append(picture, line)
		line = nil
	}
	return offsets, line
}

func putPixel(x, cycle int, line []byte) []byte {
	if x >= cycle%40 && x <= cycle%40+2 {
		line = append(line, '#')
	} else {
		line = append(line, ' ')
	}
	return line
}

func draw(v, cycle, x int, offsets []int, line []byte, n int) (int, []int, []byte, int) {
	for i := 0; i < n; i++ {
		line = putPixel(x+2, cycle, line)
		offsets, line = newLine(cycle, offsets, line)
		cycle++
	}
	x += v
	return x, offsets, line, cycle
}

var picture [][]byte

func ex2(ss []string) [][]byte {
	picture = nil
	cycle := 1
	x := 1
	offsets := []int{40, 80, 120, 160, 200, 240}
	var line []byte
	for i := 0; i < len(ss); i++ {
		if ss[i][:4] == "noop" {
			x, offsets, line, cycle = draw(0, cycle, x, offsets, line, 1)
		} else {
			split := strings.Split(ss[i], " ")
			v, _ := strconv.Atoi(split[1])
			x, offsets, line, cycle = draw(v, cycle, x, offsets, line, 2)
		}
	}
	return picture
}

func main() {
	f, _ := os.Open("input.txt")
	data, _ := io.ReadAll(f)
	ss := strings.Split(string(data), "\n")

	t := time.Now()
	fmt.Println(ex1(ss))
	fmt.Println(time.Since(t))

	t = time.Now()
	for _, s := range ex2(ss) {
		fmt.Println(string(s))
	}
	fmt.Println(time.Since(t))
}
