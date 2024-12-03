package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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
	sum := 0

	pattern := regexp.MustCompile(`mul\((\d*),(\d*)\)`)

	// xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
	matches := pattern.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		a, a_err := strconv.Atoi(match[1])
		b, b_err := strconv.Atoi(match[2])

		if a_err != nil || b_err != nil {
			panic("a or b are not int")
		}
		sum += a * b
	}

	return sum
}

func part2(input string) int64 {
	var sum int64 = 0

	pattern := regexp.MustCompile(`mul\((\d*),(\d*)\)|don\'t\(\)|do\(\)`)

	enabled := true

	matches := pattern.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		cmd := match[0]

		switch cmd {

		case "do()":
			enabled = true

		case "don't()":
			enabled = false

		default:
			a, a_err := strconv.ParseInt(match[1], 10, 64)
			b, b_err := strconv.ParseInt(match[2], 10, 64)

			if a_err != nil || b_err != nil {
				panic("a or b are not int")
			}

			if enabled {
				sum += a * b
			}
		}
	}

	return sum
}
