package main

import (
	"os"
	"strings"
	"testing"
)

func benchmarkSolutionOpti(k int, b *testing.B) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		b.Fatalf("failed to read input: %s", err)
	}
	tokens := strings.Split(string(input), "\n")
	b.Run("solution", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			opti(tokens, k)
		}
	})
}

func BenchmarkSolutionTopk3Opti(b *testing.B)  { benchmarkSolutionOpti(3, b) }
func BenchmarkSolutionTopk5Opti(b *testing.B)  { benchmarkSolutionOpti(5, b) }
func BenchmarkSolutionTopk10Opti(b *testing.B) { benchmarkSolutionOpti(10, b) }
func BenchmarkSolutionTopk25Opti(b *testing.B) { benchmarkSolutionOpti(25, b) }
func BenchmarkSolutionTopk50Opti(b *testing.B) { benchmarkSolutionOpti(50, b) }
