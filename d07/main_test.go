package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	d := dataT{
		line: []string{
			"190: 10 19",
			"3267: 81 40 27",
			"83: 17 5",
			"156: 15 6",
			"7290: 6 8 6 15",
			"161011: 16 10 13",
			"192: 17 8 14",
			"21037: 9 7 18 13",
			"292: 11 6 16 20",
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
		sumA: 3749,
	}
	t.Run(tt.name, func(t *testing.T) {
		tt.sumQ, _ = d.p1()
		if tt.sumA != tt.sumQ {
			t.Fatalf("\n%s: Expected: %d - got: %d\n", tt.name, tt.sumA, tt.sumQ)
		}
	})
}

func TestPart2(t *testing.T) {
	d := dataT{
		line: []string{
			"",
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
		sumA: 0,
	}
	t.Run(tt.name, func(t *testing.T) {
		tt.sumQ, _ = d.p2()
		if tt.sumA != tt.sumQ {
			t.Fatalf("\n%s: Expected: %d - got: %d\n", tt.name, tt.sumA, tt.sumQ)
		}
	})
}
