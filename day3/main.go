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
		var tmp []byte
		for idx := 0; idx < len(s)/2; idx++ {
			for idx2 := len(s) / 2; idx2 < len(s); idx2++ {
				if s[idx] == s[idx2] {
					if contains(string(tmp), s[idx]) {
						continue
					} else {
						if s[idx] >= 'a' {
							sum += int(s[idx]-'a') + 1
						} else {
							sum += int(s[idx]-'A') + 1 + 26
						}
						tmp = append(tmp, s[idx])
					}
				}
			}
		}
	}
	return sum
}

func contains(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

func getMatches(s1, s2 string) []byte {
	var matches []byte

	for idx := 0; idx < len(s1); idx++ {
		for idx2 := 0; idx2 < len(s2); idx2++ {
			if s1[idx] == s2[idx2] {
				if contains(string(matches), s2[idx2]) {
					continue
				} else {
					matches = append(matches, s2[idx2])
				}
			}
		}
	}
	return matches
}

func match(s1, s2, s3 string) int {
	m1 := getMatches(s1, s2)
	m2 := getMatches(s1, s3)
	m3 := getMatches(string(m1), string(m2))
	if m3[0] >= 'a' {
		return int(m3[0]-'a') + 1
	} else {
		return int(m3[0]-'A') + 1 + 26
	}
}

func ex2(ss []string) int {
	sum := 0
	for idx0 := 0; idx0 < len(ss); idx0 += 3 {
		//small speed optimization
		if len(ss[idx0]) <= len(ss[idx0+1]) && len(ss[idx0]) <= len(ss[idx0+2]) {
			sum += match(ss[idx0], ss[idx0+1], ss[idx0+2])
		} else if len(ss[idx0+1]) <= len(ss[idx0]) && len(ss[idx0+1]) <= len(ss[idx0+2]) {
			sum += match(ss[idx0+1], ss[idx0], ss[idx0+2])
		} else {
			sum += match(ss[idx0+2], ss[idx0], ss[idx0+1])
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
