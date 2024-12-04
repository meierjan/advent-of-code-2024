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

var directions = [][2]int{
	//{0,0},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
	{0, -1},
	{1, -1},
}

func part1(input string) int {
	runeArray := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		runeArray = append(runeArray, []rune(line))
	}

	count := 0

	for y := 0; y < len(runeArray); y++ {
		for x := 0; x < len(runeArray[y]); x++ {
			for _, direction := range directions {
				if checkDirection(runeArray, y, x, direction[0], direction[1]) {
					count++
				}
			}
		}
	}

	return count
}

func checkDirection(input [][]rune, y int, x int, direction_y int, direction_x int) bool {
	if direction_x == 0 && direction_y == 0 {
		panic("illegal argument")
	}

	// xmas = |4|
	xmas := []rune("XMAS")
	for i, c_rune := range xmas {
		i_y := y + i*direction_y
		i_x := x + i*direction_x

		// check if out of bounds
		if i_x < 0 || i_x == len(input[0]) || i_y < 0 || i_y == len(input) {
			return false
		}

		// check if xmas
		if input[i_y][i_x] != c_rune {
			return false
		}
	}
	return true
}

func part2(input string) int {
	runeArray := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		runeArray = append(runeArray, []rune(line))
	}

	count := 0

	for y := 0; y < len(runeArray); y++ {
		for x := 0; x < len(runeArray[y]); x++ {
			if detectMas(runeArray, y, x) {
				count++
			}
		}
	}

	return count
}

var rotation_coordinated = [][2]int{
	{1, -1},
	{1, 1},
	{-1, 1},
	{-1, -1},
}

func detectMas(input [][]rune, y int, x int) bool {
	count := 0
	for i := 0; i < 4; i++ {
		posM := rotation_coordinated[i]
		posS := rotation_coordinated[(i+2)%4]

		p1 := []int{y + posM[0], x + posM[1]}
		p2 := []int{y + posS[0], x + posS[1]}

		if !checkInBonds(input, p1, p2) {
			return false
		}

		if input[p1[0]][p1[1]] == 'M' && input[y][x] == 'A' && input[p2[0]][p2[1]] == 'S' {
			count++
		}

	}
	return count > 1
}

func checkInBonds(input [][]rune, m []int, s []int) bool {
	if m[1] < 0 || m[1] == len(input[0]) || m[0] < 0 || m[0] == len(input) {
		return false
	}

	if s[1] < 0 || s[1] == len(input[0]) || s[0] < 0 || s[0] == len(input) {
		return false
	}

	return true
}
