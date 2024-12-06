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
		fmt.Printf("Part 1 failed: %v\n", err)
	}
	fmt.Printf("Part1: %d\n", p1)

	// part 2
	p2, err := d.p2()
	if err != nil {
		fmt.Printf("Part 2 failed: %v\n", err)
	}
	fmt.Printf("Part2: %d\n", p2)

	// finish
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func (d dataT) p1() (int, error) {
	//var realMul []string
	var sum int
	return sum, nil
}

func (d dataT) p2() (int, error) {
	var sum int
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
