package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type dataT struct {
	code []string
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
	fmt.Printf("Part1: Sum of mul(x,y) instances: %d\n", p1)

	// part 2
	p2, err := d.p2()
	if err != nil {
		fmt.Printf("Failed generating sum of mul(x,y) with control codes: %v\n", err)
	}
	fmt.Printf("Part2: Sum of mul(x,y) instances with control codes: %d\n", p2)

	// finish
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func (d dataT) p1() (int, error) {
	//var realMul []string
	var sum int
	re := regexp.MustCompile(`mul[(][0-9]{1,3}[,][0-9]{1,3}[)]`)
	for _, v := range d.code {
		match := re.FindAllString(v, -1)
		for _, w := range match {
			w = strings.Replace(w, "mul(", "", -1)
			w = strings.Replace(w, ")", "", -1)
			sum += multiplyFromString(w)
		}
	}
	return sum, nil
}

func (d dataT) p2() (int, error) {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	var p2d dataT
	var skip bool
	for _, line := range d.code {
		for _, snippet := range re.FindAllString(line, -1) {
			if strings.Contains(snippet, `don't`) {
				skip = true
				continue
			}
			if strings.Contains(snippet, `do`) {
				skip = false
			}
			if !skip {
				p2d.code = append(p2d.code, snippet)
			}
		}
	}
	sum, err := p2d.p1()
	if err != nil {
		fmt.Printf("Failed generating sum of mul(x,y): %v\n", err)
	}
	return sum, nil
}

func multiplyFromString(s string) int {
	num := strings.Split(s, ",")
	l, _ := strconv.Atoi(num[0])
	r, _ := strconv.Atoi(num[1])
	return l * r
}

func (d *dataT) importData() error {
	f, err := os.ReadFile("./data.txt")
	if err != nil {
		return err
	}
	line := strings.Split(string(f), "\n")
	d.code = append(d.code, line...)
	fmt.Printf("loaded %d data-points\n", len(d.code))
	return nil
}
