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
	rule []string
}

func main() {
	// setup
	var d dataT
	if err := d.importData(); err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	// split input
	if err := d.splitData(); err != nil {
		log.Panicf("Failed splitting data: %v\n", err)
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
	var sum int
	return sum, nil
}

func (d dataT) p2() (int, error) {
	var sum int
	return sum, nil
}

func (d *dataT) splitData() error {
	var split bool
	var index int
	for i, v := range d.line {
		if split {
			d.rule = append(d.rule, v)
		}
		if v == "" {
			split = true
			index = i
		}
	}
	temp := d.line[0:index]
	d.line = nil
	d.line = temp
	//fmt.Printf("%v\n\n%v\n", d.line, d.rule)
	return nil
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
