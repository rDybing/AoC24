package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
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
		name: "part1",
		data: d,
		sumA: 0,
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
