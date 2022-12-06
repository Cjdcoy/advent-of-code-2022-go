package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func contains(c uint8, s string) bool {
	for i := 0; i < len(s); i++ {
		if c == s[i] {
			return true
		}
	}
	return false
}

// Speed oriented algorithm
func ex(ss string, k int) int {
	for i := 0; i < len(ss); i++ {
		var n = 0
		var c = 0
		for ; n < k-1; n++ {
			if contains(ss[i+n], ss[i+n+1:i+k]) {
				break
			} else {
				c++
			}
		}
		if c == k-1 {
			return i + k
		}
	}
	return 0
}

func main() {
	f, _ := os.Open("input.txt")
	data, _ := io.ReadAll(f)
	ss := strings.Split(string(data), "\n")

	t := time.Now()
	fmt.Println(ex(ss[0], 4))
	fmt.Println(time.Since(t))

	t = time.Now()
	fmt.Println(ex(ss[0], 14))
	fmt.Println(time.Since(t))
}
