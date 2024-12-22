package main

import (
	_ "embed"
	"flag"
	"fmt"
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

type Coord [2]int

type Antinode struct {
	y int
	x int
}

func part1(input string) int {

	coords_by_signal := make(map[string][]Coord)

	// 1   2

	lines := strings.Split(input, "\n")
	max_y := len(lines)
	var max_x int
	for y, line := range lines {
		positions := strings.Split(line, "")
		max_x = len(positions)

		for x, v := range positions {
			if v != "." {
				coords_by_signal[v] = append(coords_by_signal[v], Coord{y, x})
			}
		}
	}

	in_bounds := func(a Antinode) bool {
		if 0 > a.y || a.y >= max_y {
			return false
		}
		if 0 > a.x || a.x >= max_x {
			return false
		}
		return true
	}

	antinodes := make(map[Antinode]bool)

	for _, coords := range coords_by_signal {
		for i, p1 := range coords {
			for j, p2 := range coords {
				if i == j {
					continue
				}
				d := [2]int{p2[0] - p1[0], p2[1] - p1[1]}

				// a1 = p1 - d -> (p1x - dx, p1y - dy)
				a1 := Antinode{y: p1[0] - d[0], x: p1[1] - d[1]}

				if in_bounds(a1) {
					antinodes[a1] = true
				}

				// a2 = p2 + d -> (p2x + dx, p2y + dy)
				a2 := Antinode{y: p2[0] + d[0], x: p2[1] + d[1]}
				if in_bounds(a2) {
					antinodes[a2] = true
				}
			}
		}
	}

	// note that (p1, p2) will generate the same values as (p2, p1). Therfore, we could optimize for that property

	return len(antinodes)
}

func part2(input string) int {
	coords_by_signal := make(map[string][]Coord)

	// 1   2

	lines := strings.Split(input, "\n")
	max_y := len(lines)
	var max_x int
	for y, line := range lines {
		positions := strings.Split(line, "")
		max_x = len(positions)

		for x, v := range positions {
			if v != "." {
				coords_by_signal[v] = append(coords_by_signal[v], Coord{y, x})
			}
		}
	}

	in_bounds := func(a Antinode) bool {
		if 0 > a.y || a.y >= max_y {
			return false
		}
		if 0 > a.x || a.x >= max_x {
			return false
		}
		return true
	}

	antinodes := make(map[Antinode]bool)

	for _, coords := range coords_by_signal {
		for i, p1 := range coords {
			for j, p2 := range coords {
				if i == j {
					continue
				}
				d := [2]int{p2[0] - p1[0], p2[1] - p1[1]}

				// a1 = p1 - d -> (p1x - dx, p1y - dy)
				for n := 0; true; n++ {
					a1 := Antinode{y: p1[0] - (n*d[0]), x: p1[1] - (n*d[1])}
					if in_bounds(a1) {
						antinodes[a1] = true
					} else {
						break
					}
				}

				// a2 = p2 + d -> (p2x + dx, p2y + dy)
				for multi := 0; true; multi++ {
                    a2 := Antinode{y: p2[0] + (multi*d[0]), x: p2[1] + (multi*d[1])}
					if in_bounds(a2) {
						antinodes[a2] = true
					} else {
						break
					}
				}
			}
		}
	}

	// note that (p1, p2) will generate the same values as (p2, p1). Therfore, we could optimize for that property

	return len(antinodes)

}
