package main

import (
	"fmt"
	"sort"
	"math"
	"strconv"
	"strings"
)

func main() {
	min := 284639
	max := 748759

	passwords1 := []int{}
	passwords2 := []int{}

	for i := min; i <= max; i++ {
		nums := []int{}
		for _, n := range strings.Split(strconv.Itoa(i), "") {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		if sort.SliceIsSorted(nums, func(i, j int) bool) { return nums[i] < nums[j] }) {
			if hasSame
		}
	}

	
}

func check(x []int) {
	a := true
	b := false

	for i := 0; i < len(x)-1; i++ {
		if x[i] > x[i+1] {
			a = false
			break
		}

		if x[i] == x[i+1] {
			if x[i-1] != x[i] || x[i+1] != x[i+2] {

				b = true
			}

		}
	}

	if a && b {
		count++
	}
}

func sliceInts(x int) (s []int) {
	for i := len(strconv.Itoa(x)); i >= 0; i-- {
		r := x % int(math.Pow(10, float64(i+1)))
		r /= int(math.Pow(10, float64(i)))

		s = append(s, r)
	}
	return s
}

func deslice(x []int) string {
	S := x

	return (arrayToString(S, ""))
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
