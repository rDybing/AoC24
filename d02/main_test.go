package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	d := dataT{
		leftColumn:  []int{3, 4, 2, 1, 3, 3},
		rightColumn: []int{4, 3, 5, 3, 9, 3},
	}
	tt := struct {
		name string
		data dataT
		sumQ int
		sumA int
	}{
		name: "part1",
		data: d,
		sumA: 11,
	}
	t.Run(tt.name, func(t *testing.T) {
		d.sortData()
		tt.sumQ, _ = d.addDistance()
		if tt.sumA != tt.sumQ {
			t.Fatalf("\n%s: Expected: %d - got: %d\n", tt.name, tt.sumA, tt.sumQ)
		}
	})
}

func TestPart2(t *testing.T) {
	d := dataT{
		leftColumn:  []int{3, 4, 2, 1, 3, 3},
		rightColumn: []int{4, 3, 5, 3, 9, 3},
	}
	tt := struct {
		name string
		data dataT
		sumQ int
		sumA int
	}{
		name: "part2",
		data: d,
		sumA: 31,
	}
	t.Run(tt.name, func(t *testing.T) {
		d.sortData()
		tt.sumQ, _ = d.addSimilar()
		if tt.sumA != tt.sumQ {
			t.Fatalf("\n%s: Expected: %d - got: %d\n", tt.name, tt.sumA, tt.sumQ)
		}
	})
}
