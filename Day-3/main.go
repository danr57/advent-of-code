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

type distances struct {
	manhattan int64
	pathDist  int64
}

func main() {
	wire1 := map[coords]distances{}
	wire2 := map[coords]distances{}

	w, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	s := strings.Split(string(w), "\n")
	wire1directions := strings.Split(string(s[0]), ",")
	wire2directions := strings.Split(string(s[1]), ",")
	mapWire(wire1, wire1directions)
	mapWire(wire2, wire2directions)

	// Solution 1
	solution1, _ := compareMaps(wire1, wire2)
	fmt.Println("Manhattan distance to closest crossed wires:", solution1[0])

	// Solution 2
	_, solution2 := compareMaps(wire1, wire2)
	fmt.Println("Path length to the closest crossed wires:", solution2[0])
}

func mapWire(wire map[coords]distances, route []string) {
	pos := coords{x: 0, y: 0}
	p := int64(0)

	for _, dir := range route {
		d := (string(dir[0]))
		n, _ := strconv.ParseInt(dir[1:], 0, 64)

		switch d {
		case "U":
			for i := int64(0); i < n; i++ {
				pos.y++
				m := (math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
				p++
				wire[pos] = distances{manhattan: int64(m), pathDist: p}
			}
		case "D":
			for i := int64(0); i < n; i++ {
				pos.y--
				m := (math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
				p++
				wire[pos] = distances{manhattan: int64(m), pathDist: p}
			}
		case "L":
			for i := int64(0); i < n; i++ {
				pos.x--
				m := (math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
				p++
				wire[pos] = distances{manhattan: int64(m), pathDist: p}
			}
		case "R":
			for i := int64(0); i < n; i++ {
				pos.x++
				m := (math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
				p++
				wire[pos] = distances{manhattan: int64(m), pathDist: p}
			}
		}
	}
}

func compareMaps(wire1, wire2 map[coords]distances) (mh, pd []int) {

	for a, b := range wire1 {
		if c, p := wire2[a]; p {
			mh = append(mh, int(b.manhattan))
			pd = append(pd, int(b.pathDist)+int(c.pathDist))
		}
	}
	sort.Ints(mh)
	sort.Ints(pd)
	return mh, pd
}
