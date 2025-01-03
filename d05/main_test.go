package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	d := dataT{
		line: []string{
			"47|53",
			"97|13",
			"97|61",
			"97|47",
			"75|29",
			"61|13",
			"75|53",
			"29|13",
			"97|29",
			"53|29",
			"61|53",
			"97|53",
			"61|29",
			"47|13",
			"75|47",
			"97|75",
			"47|61",
			"75|61",
			"47|29",
			"75|13",
			"53|13",
			"",
			"75,47,61,53,29",
			"97,61,53,29,13",
			"75,29,13",
			"75,97,47,61,53",
			"61,13,29",
			"97,13,75,29,47"},
	}
	tt := struct {
		name string
		data dataT
		sumQ int
		sumA int
	}{
		name: "part1",
		data: d,
		sumA: 143,
	}
	t.Run(tt.name, func(t *testing.T) {
		d.splitData()
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
