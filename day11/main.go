package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type (
	monkey struct {
		items  []int
		op     string
		weight int
		div    int
		true   int
		false  int

		nbOp int
	}
)

func monkeyParse(ss []string, i int) (monkey, int) {
	var m monkey
	var s string

	s = strings.Replace(ss[i], "  Starting items: ", "", -1)
	split := strings.Split(strings.Replace(s, ",", "", -1), " ")
	for _, v := range split {
		level, _ := strconv.Atoi(v)
		m.items = append(m.items, level)
	}

	s = strings.Replace(ss[i+1], "  Operation: new = old ", "", -1)
	split = strings.Split(s, " ")
	m.op = split[0]
	var err error
	if m.weight, err = strconv.Atoi(split[1]); err != nil {
		m.weight = -1
	}

	s = strings.Replace(ss[i+2], "  Test: divisible by ", "", -1)
	m.div, _ = strconv.Atoi(s)

	s = strings.Replace(ss[i+3], "    If true: throw to monkey ", "", -1)
	m.true, _ = strconv.Atoi(s)

	s = strings.Replace(ss[i+4], "    If false: throw to monkey ", "", -1)
	m.false, _ = strconv.Atoi(s)
	return m, i + 5
}

var mod = 1

func getMonkeys(ss []string) []monkey {
	var monkeys []monkey
	for i := 0; i < len(ss); i++ {
		if len(ss[i]) < 1 {
			continue
		} else {
			var tmp monkey
			tmp, i = monkeyParse(ss, i+1)
			mod *= tmp.div
			monkeys = append(monkeys, tmp)
		}
	}
	return monkeys
}

func monkeyTurn(monkeys []monkey, i int, one bool) []monkey {
	for a := 0; a < len(monkeys[i].items); a++ {
		var worryLevel int
		switch monkeys[i].op {
		case "*":
			if monkeys[i].weight == -1 {
				worryLevel = monkeys[i].items[a] * monkeys[i].items[a]
			} else {
				worryLevel = monkeys[i].items[a] * monkeys[i].weight
			}
		case "+":
			if monkeys[i].weight == -1 {
				worryLevel = monkeys[i].items[a] + monkeys[i].items[a]
			} else {
				worryLevel = monkeys[i].items[a] + monkeys[i].weight
			}
		}
		if one {
			worryLevel = worryLevel / 3
		} else {
			worryLevel %= mod
		}
		if worryLevel%monkeys[i].div == 0 {
			monkeys[monkeys[i].true].items = append(monkeys[monkeys[i].true].items, worryLevel)
		} else {
			monkeys[monkeys[i].false].items = append(monkeys[monkeys[i].false].items, worryLevel)
		}
		monkeys[i].nbOp++
	}
	monkeys[i].items = nil
	return monkeys
}

func monkeyRun(monkeys []monkey, n int, one bool) []monkey {
	for round := 0; round < n; round++ {
		for i := 0; i < len(monkeys); i++ {
			monkeys = monkeyTurn(monkeys, i, one)
		}
	}
	return monkeys
}

func monkeyBusiness(monkeys []monkey) int {
	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].nbOp > monkeys[j].nbOp
	})

	return monkeys[0].nbOp * monkeys[1].nbOp
}

func ex1(ss []string) int {
	monkeys := getMonkeys(ss)
	monkeys = monkeyRun(monkeys, 20, true)
	return monkeyBusiness(monkeys)
}
func ex2(ss []string) int {
	monkeys := getMonkeys(ss)
	monkeys = monkeyRun(monkeys, 10000, false)
	return monkeyBusiness(monkeys)
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
