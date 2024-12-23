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
	field := make([][]int, 0)

	for _, line := range strings.Split(input, "\n") {
		list := make([]int, len(line))
		for x, elevation := range strings.Split(line, "") {
			list[x], _ = strconv.Atoi(elevation)
		}
		field = append(field, list)
	}

	sum := 0
	for y, line := range field {
		for x, elevation := range line {
			if elevation == 0 {
                // map is local to coordiante
				sum += findRek(field, y, x, elevation, make(map[string] bool))
			}

		}
	}

	return sum

}

func findRek(field [][]int, y, x int, elevation int, visited map[string]bool) int {

	if x < 0 || len(field[0]) <= x {
		return 0
	}

	if y < 0 || len(field) <= y {
		return 0
	}

	if field[y][x] == elevation {
		if elevation == 9 {
			key := fmt.Sprintf("%d,%d", y, x)
            // NOTE:
            // this is hack, but I wouldn't find another way to reuse the function for both 
            // parts. If visited is nill, caching is disabled
            if visited == nil {
                return 1
            } else 	if !visited[key] {
                visited[key] = true 
				return 1
			} else {
				return 0
			}
		} else {

			return findRek(field, y, x-1, elevation+1, visited) +
				findRek(field, y, x+1, elevation+1, visited) +
				findRek(field, y-1, x, elevation+1, visited) +
				findRek(field, y+1, x, elevation+1, visited)

		}
	} else {
		return 0
	}

}

func part2(input string) int {
    field := make([][]int, 0)

	for _, line := range strings.Split(input, "\n") {
		list := make([]int, len(line))
		for x, elevation := range strings.Split(line, "") {
			list[x], _ = strconv.Atoi(elevation)
		}
		field = append(field, list)
	}

	sum := 0
	for y, line := range field {
		for x, elevation := range line {
			if elevation == 0 {
				sum += findRek(field, y, x, elevation, nil)
			}

		}
	}

	return sum
}
