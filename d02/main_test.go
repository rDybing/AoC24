package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	d := dataT{
		report: [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9},
		},
	}
	tt := struct {
		name string
		data dataT
		sumQ int
		sumA int
	}{
		name: "part1",
		data: d,
		sumA: 2,
	}
	t.Run(tt.name, func(t *testing.T) {
		tt.sumQ, _, _ = d.safeReports()
		if tt.sumA != tt.sumQ {
			t.Fatalf("\n%s: Expected: %d - got: %d\n", tt.name, tt.sumA, tt.sumQ)
		}
	})
}

func TestPart2(t *testing.T) {
	d := dataT{
		report: [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9},
		},
	}
	tt := struct {
		name string
		data dataT
		sumQ int
		sumA int
	}{
		name: "part2",
		data: d,
		sumA: 4,
	}
	t.Run(tt.name, func(t *testing.T) {
		safe, unsafe, _ := d.safeReports()
		dampened, _ := unsafe.dampenedReports()
		tt.sumQ = safe + dampened
		if tt.sumA != tt.sumQ {
			t.Fatalf("\n%s: Expected: %d - got: %d\n", tt.name, tt.sumA, tt.sumQ)
		}
	})
}
