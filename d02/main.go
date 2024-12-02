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

	// part 1

	// finish
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
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
