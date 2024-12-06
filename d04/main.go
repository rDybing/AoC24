package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type dataT struct {
	line []string
}

func main() {
	// setup
	var d dataT
	err := d.importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()

	// part 1
	p1, err := d.p1()
	if err != nil {
		fmt.Printf("Failed generating sum of mul(x,y): %v\n", err)
	}
	fmt.Printf("Part1: %d\n", p1)

	// part 2
	p2, err := d.p2()
	if err != nil {
		fmt.Printf("Failed generating sum of mul(x,y) with control codes: %v\n", err)
	}
	fmt.Printf("Part2: %d\n", p2)

	// finish
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func (d dataT) p1() (int, error) {
	var sum int
	maxRow := len(d.line)
	maxCol := len(d.line[0])
	for indexRow := 0; indexRow < maxRow; indexRow++ {
		for indexCol := 0; indexCol < maxCol; indexCol++ {
			// check right
			if (indexCol+3 < maxCol) &&
				(string(d.line[indexRow][indexCol]) == "X" &&
					string(d.line[indexRow][indexCol+1]) == "M" &&
					string(d.line[indexRow][indexCol+2]) == "A" &&
					string(d.line[indexRow][indexCol+3]) == "S") {
				sum++
				fmt.Print("→")
			}
			// check left
			if (indexCol+3 < maxCol) &&
				(string(d.line[indexRow][indexCol]) == "S" &&
					string(d.line[indexRow][indexCol+1]) == "A" &&
					string(d.line[indexRow][indexCol+2]) == "M" &&
					string(d.line[indexRow][indexCol+3]) == "X") {
				sum++
				fmt.Print("←")
			}
			// check down
			if (indexRow+3 < maxRow) &&
				(string(d.line[indexRow][indexCol]) == "X" &&
					string(d.line[indexRow+1][indexCol]) == "M" &&
					string(d.line[indexRow+2][indexCol]) == "A" &&
					string(d.line[indexRow+3][indexCol]) == "S") {
				sum++
				fmt.Print("↓")
			}
			// check up
			if (indexRow+3 < maxRow) &&
				(string(d.line[indexRow][indexCol]) == "S" &&
					string(d.line[indexRow+1][indexCol]) == "A" &&
					string(d.line[indexRow+2][indexCol]) == "M" &&
					string(d.line[indexRow+3][indexCol]) == "X") {
				sum++
				fmt.Print("↑")
			}
			// check down-right
			if (indexRow+3 < maxRow && indexCol+3 < maxCol) &&
				(string(d.line[indexRow][indexCol]) == "X" &&
					string(d.line[indexRow+1][indexCol+1]) == "M" &&
					string(d.line[indexRow+2][indexCol+2]) == "A" &&
					string(d.line[indexRow+3][indexCol+3]) == "S") {
				sum++
				fmt.Print("↘")
			}
			// check up-left
			if (indexRow+3 < maxRow && indexCol+3 < maxCol) &&
				(string(d.line[indexRow][indexCol]) == "S" &&
					string(d.line[indexRow+1][indexCol+1]) == "A" &&
					string(d.line[indexRow+2][indexCol+2]) == "M" &&
					string(d.line[indexRow+3][indexCol+3]) == "X") {
				sum++
				fmt.Print("↖")
			}
			// check up-right
			if (indexRow-3 >= 0 && indexCol+3 < maxCol) &&
				(string(d.line[indexRow][indexCol]) == "S" &&
					string(d.line[indexRow-1][indexCol+1]) == "A" &&
					string(d.line[indexRow-2][indexCol+2]) == "M" &&
					string(d.line[indexRow-3][indexCol+3]) == "X") {
				sum++
				fmt.Print("↗")
			}
			// check down-left
			if (indexRow-3 >= 0 && indexCol+3 < maxCol) &&
				(string(d.line[indexRow][indexCol]) == "X" &&
					string(d.line[indexRow-1][indexCol+1]) == "M" &&
					string(d.line[indexRow-2][indexCol+2]) == "A" &&
					string(d.line[indexRow-3][indexCol+3]) == "S") {
				sum++
				fmt.Print("↙")
			}
		}
	}
	fmt.Println()
	return sum, nil
}

func (d dataT) p2() (int, error) {
	var sum int
	maxRow := len(d.line)
	maxCol := len(d.line[0])
	for indexRow := 0; indexRow < maxRow; indexRow++ {
		for indexCol := 0; indexCol < maxCol; indexCol++ {
			if (indexRow+2 < maxRow && indexCol+2 < maxCol) &&
				(string(d.line[indexRow][indexCol]) == "M" &&
					string(d.line[indexRow+1][indexCol+1]) == "A" &&
					string(d.line[indexRow+2][indexCol+2]) == "S" &&
					string(d.line[indexRow+2][indexCol]) == "M" &&
					string(d.line[indexRow][indexCol+2]) == "S") {
				sum++
			}
			if (indexRow+2 < maxRow && indexCol+2 < maxCol) &&
				(string(d.line[indexRow][indexCol]) == "M" &&
					string(d.line[indexRow+1][indexCol+1]) == "A" &&
					string(d.line[indexRow+2][indexCol+2]) == "S" &&
					string(d.line[indexRow+2][indexCol]) == "S" &&
					string(d.line[indexRow][indexCol+2]) == "M") {
				sum++
			}
			if (indexRow+2 < maxRow && indexCol+2 < maxCol) &&
				(string(d.line[indexRow][indexCol]) == "S" &&
					string(d.line[indexRow+1][indexCol+1]) == "A" &&
					string(d.line[indexRow+2][indexCol+2]) == "M" &&
					string(d.line[indexRow+2][indexCol]) == "M" &&
					string(d.line[indexRow][indexCol+2]) == "S") {
				sum++
			}
			if (indexRow+2 < maxRow && indexCol+2 < maxCol) &&
				(string(d.line[indexRow][indexCol]) == "S" &&
					string(d.line[indexRow+1][indexCol+1]) == "A" &&
					string(d.line[indexRow+2][indexCol+2]) == "M" &&
					string(d.line[indexRow+2][indexCol]) == "S" &&
					string(d.line[indexRow][indexCol+2]) == "M") {
				sum++
			}
		}
	}
	return sum, nil
}

func (d *dataT) importData() error {
	f, err := os.ReadFile("./data.txt")
	if err != nil {
		return err
	}
	line := strings.Split(string(f), "\n")
	d.line = append(d.line, line...)
	fmt.Printf("loaded %d data-points\n", len(d.line))
	return nil
}
