package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func insertPop(table []int, n int, k int) []int {
	tableLen := len(table)
	for idx := 0; idx < tableLen; idx++ {
		if n > table[idx] {
			if tableLen >= k {
				table = append(table[:idx+1], table[idx:tableLen-1]...)
				table[idx] = n
				break
			} else {
				table = append(table[:idx+1], table[idx:]...)
				table[idx] = n
				break
			}
		}
	}
	return table
}

func opti(ss []string, k int) {
	var table = []int{0}
	tmp := 0
	for idx := 0; idx < len(ss); idx++ {
		if ss[idx] == "" {
			table = insertPop(table, tmp, k)
			tmp = 0
		} else {
			n := FastAtoi(ss[idx])
			tmp += n
		}
	}
	fmt.Println(table[0] + table[1] + table[2])
}

func FastAtoi(s string) int {
	var val int

	for _, c := range s {
		val = val*10 + int(c-'0')
	}

	return val
}

func main() {
	f, _ := os.Open("input.txt")
	data, _ := io.ReadAll(f)
	ss := strings.Split(string(data), "\n")
	k := 3
	t := time.Now()
	opti(ss, k)
	fmt.Println(time.Since(t))
}
