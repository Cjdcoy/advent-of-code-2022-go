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
	var size, localSize, subDirs int
	var subSizes []int
	for i := idx; i < len(ss); i++ {
		if ss[i][2:4] == "cd" {
			if len(ss[i]) > 6 && ss[i][5:7] == ".." {
				//ex2
				sizes = append(sizes, localSize)
				return localSize, i
			}
			size, i = cd(ss, i+1)
			subSizes = append(subSizes, size)
			subDirs--
			localSize += size
			//ex1
			if localSize <= 100000 && subDirs == 0 {
				tot += localSize
			} else if localSize > 100000 && subDirs == 0 {
				for _, n := range subSizes {
					if n < 100000 {
						tot += n
					}
				}
			}
		} else if ss[i][0:3] == "dir" {
			subDirs++
		} else if val, err := strconv.Atoi(strings.Split(ss[i], " ")[0]); err == nil {
			localSize += val
		}
	}
	//ex2
	sizes = append(sizes, localSize)
	return localSize, len(ss)
}

func ex1(ss []string) int {
	cd(ss, 0)
	return tot
}

func ex2(ss []string) int {
	best := math.MaxInt
	totSize, _ := cd(ss, 0)
	for _, n := range sizes {
		tmp := 70000000 - (totSize) + n
		if tmp > 30000000 && n < best {
			best = n
		}
	}
	sizes = nil
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
