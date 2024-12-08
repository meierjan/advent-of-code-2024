package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

type input_line struct {
	result  int64
	numbers []int64
}

func parse_input(input string) []input_line {
	lines := make([]input_line, 0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		result, _ := strconv.ParseInt(parts[0], 10, 64)
		numbers := make([]int64, 0)
		for _, number := range strings.Split(parts[1], " ") {
			i64, _ := strconv.ParseInt(number, 10, 64)
			numbers = append(numbers, i64)


		}
        lines = append(lines, input_line{result: result, numbers: numbers})
	}

	return lines

}

func part1(input string) int64 {
	sum := int64(0)
	input_lines := parse_input(input)
	for _, line := range input_lines {
		if is_valid(line.numbers, 0, line.result, 0) {
			sum += line.result
		}

	}

	return sum
}

func is_valid(list []int64, i int, target int64, current int64) bool {
	if i == len(list) {
		return target == current
	} else {
		return is_valid(list, i+1, target, current+list[i]) || is_valid(list, i+1, target, current*list[i])
	}

}

func part2(input string) int64 {
	sum := int64(0)
	input_lines := parse_input(input)
	for _, line := range input_lines {
		if is_valid2(line.numbers, 0, line.result, 0) {
			sum += line.result
		}
	}

	return sum
}


func is_valid2(list []int64, i int, target int64, current int64) bool {
    if i == len(list) {
        return target == current
    } else {
        c_item := list[i]
        digits := float64(math.Floor(math.Log10(math.Abs(float64(c_item))))) + 1
        concat_result := (current * (int64(math.Pow(10, digits)))) + c_item
        return is_valid2(list, i+1, target, current+list[i]) || is_valid2(list, i+1, target, current*list[i]) || is_valid2(list, i+1, target, concat_result)
    }
}
