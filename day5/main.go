package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func FastAtoi(s string) int {
	var val int

	for _, c := range s {
		val = val*10 + int(c-'0')
	}

	return val
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
func GetBoxes(boxes map[int]string, ss []string) map[int]string {
	for idx, s := range ss {
		if idx == 8 || len(s) == 0 {
			break
		}
		if s[0] == '[' || s[0] == ' ' {
			var idx2 = 1
			for i := 1; i < len(s); i += 4 {
				if s[i] != ' ' {
					boxes[idx2] += string(s[i])
				}
				idx2++
			}
		}
	}
	for idx, val := range boxes {
		boxes[idx] = reverse(val)
	}
	return boxes
}

func moveBoxes(boxes map[int]string, ss []string) map[int]string {
	for i := 10; i < len(ss); i++ {
		split := strings.Split(ss[i], " ")
		n := FastAtoi(split[1])
		n1 := FastAtoi(split[3])
		n2 := FastAtoi(split[5])

		for i2 := 0; i2 < n; i2++ {
			boxes[n2] += boxes[n1][len(boxes[n1])-1:]
			boxes[n1] = boxes[n1][:len(boxes[n1])-1]
		}
	}
	return boxes
}

func ex1(ss []string) string {
	boxes := make(map[int]string)
	boxes = GetBoxes(boxes, ss)
	boxes = moveBoxes(boxes, ss)
	var res string
	for idx := 1; idx < 10; idx++ {
		res += string(boxes[idx][len(boxes[idx])-1])
	}
	return res
}

func moveBoxes2(boxes map[int]string, ss []string) map[int]string {
	for i := 10; i < len(ss); i++ {
		split := strings.Split(ss[i], " ")
		n := FastAtoi(split[1])
		n1 := FastAtoi(split[3])
		n2 := FastAtoi(split[5])
		boxes[n2] += boxes[n1][len(boxes[n1])-n:]
		boxes[n1] = boxes[n1][:len(boxes[n1])-n]
	}
	return boxes
}

func ex2(ss []string) string {
	boxes := make(map[int]string)
	boxes = GetBoxes(boxes, ss)
	boxes = moveBoxes2(boxes, ss)
	var res string
	for idx := 1; idx < 10; idx++ {
		res += string(boxes[idx][len(boxes[idx])-1])
	}
	return res
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
