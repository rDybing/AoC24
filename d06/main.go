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

/*
	mapT.pos:
	0: clear
	1: obstacle
	2: visited
*/

type mapT struct {
	pos     [][]int
	maxPosX int
	maxPosY int
	guard   guardT
}

/*
	guardT.dir:
	0: up
	1: right
	2: down
	3: right
*/

type guardT struct {
	posX int
	posY int
	dir  int
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
	var bounds, step bool
	m, err := d.parseData()
	if err != nil {
		log.Panicf("Parsing data failed: %v\n", err)
	}
	for !bounds {
		bounds, step = m.moveGuard()
		if step {
			sum++
		}
		m.drawMap()
		fmt.Printf("pos: %+v - step: %d\n", m.guard, sum)
	}
	return sum, nil
}

func (d dataT) p2() (int, error) {
	var sum int
	return sum, nil
}

func (m *mapT) moveGuard() (bool, bool) {
	var bounds, step bool
	switch m.guard.dir {
	case 0:
		m.guard.posY--
		fmt.Printf("%d\n", m.pos[m.guard.posX][m.guard.posY])
		if m.guard.posY < 0 {
			bounds = true
			step = false
		}
		if m.pos[m.guard.posY][m.guard.posX] == 1 {
			m.guard.posY++
			m.guard.dir++
			bounds = false
			step = false
		}
		if m.pos[m.guard.posY][m.guard.posX] == 2 {
			bounds = false
			step = false
		}
		if m.pos[m.guard.posY][m.guard.posX] == 0 {
			m.pos[m.guard.posY][m.guard.posX] = 2
			bounds = false
			step = true
		}
	case 1:
		m.guard.posX++
		if m.guard.posX > m.maxPosX {
			bounds = true
			step = false
		}
		if m.pos[m.maxPosY][m.guard.posX] == 1 {
			m.guard.posX--
			m.guard.dir++
			bounds = false
			step = false
		}
		if m.pos[m.guard.posY][m.guard.posX] == 2 {
			bounds = false
			step = false
		}
		if m.pos[m.guard.posY][m.guard.posX] == 0 {
			m.pos[m.guard.posY][m.guard.posX] = 2
			bounds = false
			step = true
		}
	case 3:
		m.guard.posY++
		if m.guard.posY > m.maxPosY {
			bounds = true
			step = false
		}
		if m.pos[m.maxPosY][m.guard.posX] == 1 {
			m.guard.posY--
			m.guard.dir++
			bounds = false
			step = false
		}
		if m.pos[m.guard.posY][m.guard.posX] == 2 {
			bounds = false
			step = false
		}
		if m.pos[m.guard.posY][m.guard.posX] == 0 {
			m.pos[m.guard.posY][m.guard.posX] = 2
			bounds = false
			step = true
		}
	case 4:
		m.guard.posX--
		if m.guard.posX < 0 {
			bounds = true
			step = false
		}
		if m.pos[m.maxPosY][m.guard.posX] == 1 {
			m.guard.posX++
			m.guard.dir = 0
			bounds = false
			step = false
		}
		if m.pos[m.guard.posY][m.guard.posX] == 2 {
			bounds = false
			step = false
		}
		if m.pos[m.guard.posY][m.guard.posX] == 0 {
			m.pos[m.guard.posY][m.guard.posX] = 2
			bounds = false
			step = true
		}
	}
	return bounds, step
}

func (m mapT) drawMap() {
	for y, _ := range m.pos {
		for x, v := range m.pos[y] {
			if v == 1 {
				fmt.Print("#")
			} else {
				if m.guard.posX == x && m.guard.posY == y {
					switch m.guard.dir {
					case 0:
						fmt.Print("^")
					case 1:
						fmt.Print(">")
					case 2:
						fmt.Print("v")
					case 3:
						fmt.Print("<")
					}
				} else {
					if v == 2 {
						fmt.Print("X")
					}
					if v == 0 {
						fmt.Print(".")
					}
				}
			}
		}
		fmt.Println()
	}
}

func (d dataT) parseData() (mapT, error) {
	var m mapT
	m.maxPosX = len(d.line[0]) - 1
	m.maxPosY = len(d.line) - 1
	for y, v := range d.line {
		var tempX []int
		for x, w := range v {
			switch w {
			case '#':
				tempX = append(tempX, 1)
			case '^':
				tempX = append(tempX, 0)
				m.guard.posX = x
				m.guard.posY = y
				m.guard.dir = 0
			default:
				tempX = append(tempX, 0)
			}
		}
		m.pos = append(m.pos, tempX)
	}
	return m, nil
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
