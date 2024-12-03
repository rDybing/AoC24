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
	report [][]int
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
	safe, unsafe, err := d.safeReports()
	if err != nil {
		fmt.Printf("Failed generating number of safe reports: %v\n", err)
	}
	fmt.Printf("Part1: Safe reports sum: %d\n", safe)

	// part 2
	var dampened int
	dampened, err = unsafe.dampenedReports()
	if err != nil {
		fmt.Printf("Failed generating number of safe and dampened reports: %v\n", err)
	}
	fmt.Printf("Part2: Safe + dampened reports sum: %d\n", safe+dampened)

	// finish
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func (d dataT) dampenedReports() (int, error) {
	var sum int
	for _, row := range d.report {
		var dampSafe bool
		for i := range row {
			newRow := removeIndex(row, i)
			safe := getSafe(newRow)
			if safe {
				dampSafe = true
			}
		}
		if dampSafe {
			sum++
		}
	}
	return sum, nil
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func (d dataT) safeReports() (int, dataT, error) {
	var sum int
	var unsafe dataT
	for _, row := range d.report {
		safe := getSafe(row)
		if safe {
			sum++
		} else {
			// for part 2
			unsafe.report = append(unsafe.report, row)
		}
	}
	return sum, unsafe, nil
}

func getSafe(row []int) bool {
	var prev, next int
	var safe, ascending, dirSet bool
	for _, level := range row {
		if prev == 0 {
			prev = level
			continue
		} else {
			next = level
		}
		// break if equal
		if next == prev {
			safe = false
			break
		}
		// ascending
		if next > prev {
			if !dirSet {
				ascending = true
				dirSet = true
			}
			diff := next - prev
			if diff <= 3 && ascending {
				safe = true
			} else {
				safe = false
				break
			}
		}
		// descending
		if next < prev {
			if !dirSet {
				ascending = false
				dirSet = true
			}
			diff := prev - next
			if diff <= 3 && !ascending {
				safe = true
			} else {
				safe = false
				break
			}
		}
		prev = next
	}
	return safe
}

func (d *dataT) importData() error {
	f, err := os.ReadFile("./data.txt")
	if err != nil {
		return err
	}
	arr := strings.Split(string(f), "\n")
	for _, v := range arr {
		var levels []int
		row := strings.Split(string(v), " ")
		for _, w := range row {
			temp, _ := strconv.Atoi(w)
			levels = append(levels, temp)
		}
		d.report = append(d.report, levels)
	}
	fmt.Printf("loaded %d data-points\n", len(d.report))
	return nil
}
