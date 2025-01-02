package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
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
		fmt.Println("Output: ", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output: ", ans)
	}
}

type point struct {
	x int
	y int
}

type shape struct {
	edges     []point
	points    []point
	category  string
	perimeter int
	area      int
}

func part1(input string) int {
	field := make([][]string, 0)
	for y, line := range strings.Split(input, "\n") {
		field = append(field, make([]string, 0))
		for _, category := range strings.Split(line, "") {
			field[y] = append(field[y], category)
		}
	}

	visited := make(map[point]bool)
	shapes := make([]shape, 0)
	for y := range field {
		for x := range field[y] {
			point := point{y: y, x: x}
			shape, is_new := findShape(field, visited, point)

			if is_new {
				shapes = append(shapes, shape)
			}
		}
	}

	sum := 0
	for _, shape := range shapes {
		sum += shape.area * shape.perimeter
	}

	return sum
}

func isSameCategory(field [][]string, category string, point point) bool {

	if point.y < 0 || len(field) <= point.y {
		return false
	}

	if point.x < 0 || len(field[point.y]) <= point.x {
		return false
	}

	return field[point.y][point.x] == category
}

func getNeighbors(p point) []point {
	return []point{
		{x: p.x - 1, y: p.y},
		{x: p.x + 1, y: p.y},
		{x: p.x, y: p.y - 1},
		{x: p.x, y: p.y + 1},
	}
}

func findShape(field [][]string, visited map[point]bool, start point) (shape, bool) {
	if visited[start] {
		return shape{}, false
	}

	points := make([]point, 0)

	category := field[start.y][start.x]

	perimeter := 0
	area := 0

	candidates := []point{start}
	for true {

		// if no candidates let, we are done
		if len(candidates) == 0 {
			break
		}

		candidate := candidates[0]
		candidates = candidates[1:]

		neighbors := getNeighbors(candidate)

		neighbors_in_shape := 0
		for _, neighbor := range neighbors {
			if slices.Contains(points, neighbor) {
				neighbors_in_shape++
			}
		}

		perimeter += 4 + (neighbors_in_shape * -2)
		area++

		visited[candidate] = true
		points = append(points, candidate)

		// same category neighbors_in_shape
		for _, neighbor := range neighbors {
			// check neighbors if same cat
			if !visited[neighbor] && !slices.Contains(candidates, neighbor) && isSameCategory(field, category, neighbor) {
				candidates = append(candidates, neighbor)
			}
		}
	}

	return shape{
		points:    points,
		category:  category,
		perimeter: perimeter,
		area:      area,
	}, true
}

func part2(input string) int {
	return -1
}
