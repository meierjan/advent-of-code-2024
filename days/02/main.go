package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	valid_lines := 0
	for _, line := range strings.Split(input, "\n") {
		items := get_items(line)

		if is_safe_report(items) {
			valid_lines++
		}
	}

	return valid_lines
}

func remove_item_at(list []int, index int) []int {
	var new_list []int

	for i, item := range list {
		if i != index {
			new_list = append(new_list, item)
		}
	}

	return new_list

}

func part2(input string) int {
	valid_lines := 0
	for _, line := range strings.Split(input, "\n") {
		items := get_items(line)

		for i := 0; i < len(items); i++ {
			sub_items := remove_item_at(items, i)
			if is_safe_report(sub_items) {
				valid_lines++
				break
			}

		}
	}

	return valid_lines
}

func get_items(line string) []int {
	items_raw := strings.Split(line, " ")
	items := make([]int, len(items_raw))

	for i, item := range items_raw {
		i_item, err := strconv.Atoi(item)
		if err != nil {
			panic("invalid format")
		}
		items[i] = i_item
	}

	return items
}

func is_safe_report(items []int) bool {

	direction := (items[1] - items[0]) >= 0

	for i := 1; i < len(items); i++ {
		d := items[i] - items[i-1]
		d_abs := Abs(d)
		i_dir := d >= 0

		if i_dir != direction || d_abs > 3 || d_abs < 1 {
			return false
		}
	}

	return true
}
