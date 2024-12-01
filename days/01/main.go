package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}
func part1(input string) int {

	var left []int
	var right []int

	// 1   2
	for _, line := range strings.Split(input, "\n") {

		parts := strings.Split(line, "   ")
		l, r := parts[0], parts[1]

		l_i, l_err := strconv.Atoi(l)
		r_i, r_err := strconv.Atoi(r)

		if l_err != nil || r_err != nil {
			panic("invalid format")
		}

		left = append(left, l_i)
		right = append(right, r_i)
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0

	for i, l := range left {
		r := right[i]

		sum += Abs(r - l)
	}

	return sum
}

func part2(input string) int {
	var left []int
	var right []int

	for _, line := range strings.Split(input, "\n") {

		parts := strings.Split(line, "   ")
		l, r := parts[0], parts[1]

		l_i, l_err := strconv.Atoi(l)
		r_i, r_err := strconv.Atoi(r)

		if l_err != nil || r_err != nil {
			panic("invalid format")
		}

		left = append(left, l_i)
		right = append(right, r_i)
	}

	lookup := make(map[int]int)

	for _, r := range right {

		lookup[r] += 1

	}

	total_score := 0

	for _, l := range left {
		count := lookup[l]
		score := l * count

		total_score += score
	}

	return total_score
}
