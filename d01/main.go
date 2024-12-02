package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type dataT struct {
	leftColumn  []int
	rightColumn []int
}

func main() {
	// setup
	var d dataT
	err := d.importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()
	err = d.sortData()
	if err != nil {
		log.Panicf("Failed to sort data: %v\n", err)
	}

	// part 1
	var distance int
	distance, err = d.addDistance()
	if err != nil {
		log.Panicf("Failed to add distance: %v\n", err)
	}
	fmt.Printf("Sum part 1 distance = %d\n", distance)

	// part 2
	var similar int
	similar, err = d.addSimilar()
	if err != nil {
		log.Panicf("Failed to add similar: %v\n", err)
	}
	fmt.Printf("Sum part 2 similarity: %d\n", similar)

	// finish
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}
func (d dataT) addSimilar() (int, error) {
	var similar int
	for _, lc := range d.leftColumn {
		var instances int
		for _, rc := range d.rightColumn {
			if lc == rc {
				instances++
			}
		}
		similar += lc * instances
	}
	return similar, nil
}

func (d dataT) addDistance() (int, error) {
	var distance int
	for i, v := range d.leftColumn {
		if v > d.rightColumn[i] {
			distance += v - d.rightColumn[i]
		}
		if v < d.rightColumn[i] {
			distance += d.rightColumn[i] - v
		}
	}
	return distance, nil
}

func (d *dataT) sortData() error {
	sort.Slice(d.leftColumn, func(i, j int) bool {
		return d.leftColumn[i] < d.leftColumn[j]
	})
	sort.Slice(d.rightColumn, func(i, j int) bool {
		return d.rightColumn[i] < d.rightColumn[j]
	})
	return nil
}

func (d *dataT) importData() error {
	f, err := os.ReadFile("./data.txt")
	if err != nil {
		return err
	}
	arr := strings.Split(string(f), "\n")
	for _, v := range arr {
		var lc, rc int
		row := strings.Split(string(v), "   ")
		if len(row) == 2 {
			lc, _ = strconv.Atoi(row[0])
			rc, _ = strconv.Atoi(row[1])
		}
		if lc != 0 && rc != 0 {
			d.leftColumn = append(d.leftColumn, lc)
			d.rightColumn = append(d.rightColumn, rc)
		}

	}
	fmt.Printf("loaded %d data-points\n", len(d.leftColumn))
	return nil
}
