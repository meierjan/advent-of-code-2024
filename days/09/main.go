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
func part1(input string) int64 {
	id := int64(0)

	data := make([]int64, 0)

	for i, length_str := range strings.Split(input, "") {
		is_file := i%2 == 0

		length, err := strconv.ParseInt(length_str, 10, 64)

		if err != nil {
			fmt.Println(length_str)
			panic("error")
		}

		if is_file {
			for range length {
				data = append(data, id)
			}
			id++
		} else {
			for range length {
				data = append(data, -1)
			}
		}
	}

	i_end := len(data) - 1
	for i, id := range data {
		if i >= i_end {
			break
		}
		if id == -1 {
			// find data from end to write here

			new_id := int64(-1)
			for j := i_end; j > i; j-- {
				// data found
				if data[j] != -1 {
					new_id = data[j]
					data[j] = -1
					i_end = j - 1
					break
				}
			}

			if new_id != -1 {
				data[i] = new_id
			}

		}

	}

	checksum := int64(0)

	for i, id := range data {
		if id == -1 {
			break
		}

		checksum += int64(i) * id
	}

	return checksum
}

func part2(input string) int64 {
	return -1
}
