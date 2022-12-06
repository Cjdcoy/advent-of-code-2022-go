package main

import (
	"os"
	"strings"
	"testing"
)

func benchmarkSolutionEx1(b *testing.B) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		b.Fatalf("failed to read input: %s", err)
	}
	tokens := strings.Split(string(input), "\n")
	b.Run("solution", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ex(tokens[0], 4)
		}
	})
}

func benchmarkSolutionEx2(b *testing.B) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		b.Fatalf("failed to read input: %s", err)
	}
	tokens := strings.Split(string(input), "\n")
	b.Run("solution", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ex(tokens[0], 14)
		}
	})
}

func BenchmarkSolutionEx1(b *testing.B) { benchmarkSolutionEx1(b) }

func BenchmarkSolutionEx2(b *testing.B) { benchmarkSolutionEx2(b) }
