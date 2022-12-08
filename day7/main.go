package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var tot = 0
var sizes = []int{0}

func cd(ss []string, idx int) (int, int) {
	var s, localSize, subDirs int
	var subSizes []int
	for i := idx; i < len(ss); i++ {
		cmd := strings.Split(ss[i], " ")
		if val, err := strconv.Atoi(cmd[0]); err == nil {
			localSize += val
		} else if string(cmd[1]) == "cd" {
			if cmd[2] == ".." {
				return localSize, i
			}
			s, i = cd(ss, i+1)
			subSizes = append(subSizes, s)
			subDirs--
			localSize += s
			if localSize <= 100000 && subDirs == 0 {
				tot += localSize
			} else if localSize > 100000 && subDirs == 0 {
				for _, n := range subSizes {
					if n < 100000 {
						tot += n
					}
				}
			}
		} else if cmd[0] == "dir" {
			subDirs++
		}
	}
	return localSize, len(ss)
}

func ex1(ss []string) int {
	cd(ss, 0)
	return tot
}

func cd2(ss []string, idx int) (int, int) {
	var localSize, subDirs, s int
	var subSizes []int
	for i := idx; i < len(ss); i++ {
		cmd := strings.Split(ss[i], " ")
		if val, err := strconv.Atoi(cmd[0]); err == nil {
			localSize += val
		} else if string(cmd[1]) == "cd" {
			if cmd[2] == ".." {
				sizes = append(sizes, localSize)
				return localSize, i
			}
			s, i = cd2(ss, i+1)
			subSizes = append(subSizes, s)
			subDirs--
			localSize += s
		} else if cmd[0] == "dir" {
			subDirs++
		}
	}
	sizes = append(sizes, localSize)
	return localSize, len(ss)
}

func ex2(ss []string) int {
	best := math.MaxInt
	totSize, _ := cd2(ss, 0)
	for _, n := range sizes {
		tmp := 70000000 - (totSize) + n
		if tmp > 30000000 && n < best {
			best = n
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
