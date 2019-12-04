package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type coords struct {
	x int64
	y int64
}

func main() {
	wire1 := map[coords]int64{}
	wire2 := map[coords]int64{}

	w, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	s := strings.Split(string(w), "\n")
	wire1directions := strings.Split(string(s[0]), ",")
	wire2directions := strings.Split(string(s[1]), ",")
	mapWire(wire1, wire1directions)
	mapWire(wire2, wire2directions)

	overlap := compareMaps(wire1, wire2)
	fmt.Println("Manhattan distance to closest crossed wires:", overlap[0])
}

func mapWire(wire map[coords]int64, route []string) {
	pos := coords{x: 0, y: 0}

	for _, dir := range route {
		d := (string(dir[0]))
		n, _ := strconv.ParseInt(dir[1:], 0, 64)

		switch d {
		case "U":
			for i := int64(0); i < n; i++ {
				pos.y++
				m := (math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
				wire[pos] = int64(m)
			}
		case "D":
			for i := int64(0); i < n; i++ {
				pos.y--
				m := (math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
				wire[pos] = int64(m)
			}
		case "L":
			for i := int64(0); i < n; i++ {
				pos.x--
				m := (math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
				wire[pos] = int64(m)
			}
		case "R":
			for i := int64(0); i < n; i++ {
				pos.x++
				m := (math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
				wire[pos] = int64(m)
			}
		}
	}
}

func compareMaps(wire1, wire2 map[coords]int64) (overlaps []int) {

	for a, b := range wire1 {
		if _, p := wire2[a]; p {
			overlaps = append(overlaps, int(b))
		}
	}
	sort.Ints(overlaps)
	return overlaps
}
