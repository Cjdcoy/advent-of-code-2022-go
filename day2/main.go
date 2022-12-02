package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func ex1(ss []string) int {
	sum := 0
	for _, s := range ss {
		switch s[2] {
		case 'X':
			sum += 1
			switch s[0] {
			case 'A':
				sum += 3
				break
			case 'B':
				break
			case 'C':
				sum += 6
				break
			}
			break
		case 'Y':
			sum += 2
			switch s[0] {
			case 'A':
				sum += 6
				break
			case 'B':
				sum += 3
				break
			case 'C':
				break
			}
			break
		case 'Z':
			sum += 3
			switch s[0] {
			case 'A':
				break
			case 'B':
				sum += 6
				break
			case 'C':
				sum += 3
				break
			}
			break
		}
	}
	return sum
}
func ex2(ss []string) int {
	sum := 0
	for _, s := range ss {
		switch s[2] {
		case 'X':
			switch s[0] {
			case 'A':
				sum += 3
				break
			case 'B':
				sum += 1
				break
			case 'C':
				sum += 2
				break
			}
			break
		case 'Y':
			sum += 3
			switch s[0] {
			case 'A':
				sum += 1
				break
			case 'B':
				sum += 2
				break
			case 'C':
				sum += 3
				break
			}
			break
		case 'Z':
			sum += 6
			switch s[0] {
			case 'A':
				sum += 2
				break
			case 'B':
				sum += 3
				break
			case 'C':
				sum += 1
				break
			}
			break
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
