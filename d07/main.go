package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type dataT struct {
	line []string
}

type equationT struct {
	sum int
	arg []int
}

func main() {
	// setup
	var d dataT
	if err := d.importData(); err != nil {
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
	var sum int
	for _, v := range d.line {
		var eq equationT
		eq.splitLine(v)
		fmt.Printf("%+v\n", eq)
		if eq.testArgs() {
			sum += eq.sum
		}
	}
	return sum, nil
}

func (eq equationT) testArgs() bool {
}

func (eq *equationT) splitLine(line string) {
	temp := strings.Split(line, ": ")
	eq.sum, _ = strconv.Atoi(temp[0])
	temp2 := strings.Split(temp[1], " ")
	for _, v := range temp2 {
		temp3, _ := strconv.Atoi(v)
		eq.arg = append(eq.arg, temp3)
	}
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
