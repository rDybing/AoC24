package main

import (
	"testing"
)

/*
   0 1 2 3 4 5 6 7 8 9
0  M M M S X X M A S M
1  M S A M X M S M S A
2  A M X S X M A A M M
3  M S A M A S M S M X
4  X M A S A M X A M M
5  X X A M M X X A M A
6  S M S M S A S X S S
7  S A X A M A S A A A
8  M A M M M X M M M M
9  M X M X A X M A S X

right: 		ok 	-
left:		ok 	-
down:		ok 	|
up:			ok 	|
down-right: ok 	\
up-left		ok 	\
up-right	ok	/
down-left	ok	/
*/

func TestPart1(t *testing.T) {
	d := dataT{
		line: []string{
			"MMMSXXMASM",
			"MSAMXMSMSA",
			"AMXSXMAAMM",
			"MSAMASMSMX",
			"XMASAMXAMM",
			"XXAMMXXAMA",
			"SMSMSASXSS",
			"SAXAMASAAA",
			"MAMMMXMMMM",
			"MXMXAXMASX",
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
		sumA: 18,
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
			"MMMSXXMASM",
			"MSAMXMSMSA",
			"AMXSXMAAMM",
			"MSAMASMSMX",
			"XMASAMXAMM",
			"XXAMMXXAMA",
			"SMSMSASXSS",
			"SAXAMASAAA",
			"MAMMMXMMMM",
			"MXMXAXMASX",
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
		sumA: 9,
	}
	t.Run(tt.name, func(t *testing.T) {
		tt.sumQ, _ = d.p2()
		if tt.sumA != tt.sumQ {
			t.Fatalf("\n%s: Expected: %d - got: %d\n", tt.name, tt.sumA, tt.sumQ)
		}
	})
}
