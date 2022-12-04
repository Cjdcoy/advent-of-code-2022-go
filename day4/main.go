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

func ex1(ss []string) int {
	sum := 0
	for i := 0; i < len(ss); i++ {
		s := strings.Split(ss[i], ",")
		s2 := strings.Split(s[0], "-")
		s3 := strings.Split(s[1], "-")
		r1 := [2]int{0}
		r2 := [2]int{0}
		r1[0] = FastAtoi(s2[0])
		r1[1] = FastAtoi(s2[1])
		r2[0] = FastAtoi(s3[0])
		r2[1] = FastAtoi(s3[1])
		if r1[0] <= r2[0] && r1[1] >= r2[1] {
			sum++
		} else if r2[0] <= r1[0] && r2[1] >= r1[1] {
			sum++
		}
	}
	return sum
}

func ex2(ss []string) int {
	sum := 0
	for i := 0; i < len(ss); i++ {
		s := strings.Split(ss[i], ",")
		s2 := strings.Split(s[0], "-")
		s3 := strings.Split(s[1], "-")
		r1 := [2]int{0}
		r2 := [2]int{0}
		r1[0] = FastAtoi(s2[0])
		r1[1] = FastAtoi(s2[1])
		r2[0] = FastAtoi(s3[0])
		r2[1] = FastAtoi(s3[1])
		if r1[0] <= r2[0] && r1[1] >= r2[0] || r1[1] >= r2[1] && r1[0] <= r2[1] {
			sum++
		} else if r2[0] <= r1[0] && r2[1] >= r1[0] || r2[1] >= r1[1] && r2[0] <= r2[1] {
			sum++
		}
	}
	return sum
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
