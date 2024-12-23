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
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

var directions = [][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func is_in_bounds(field [][]bool, y int, x int) bool {
	return 0 <= y && 0 <= x && y < len(field) && x < len(field[y])
}

func count_visits(field [][]bool) int {
	count := 0
	for _, line := range field {
		for _, is_visited := range line {
			if is_visited {
				count++
			}
		}
	}
	return count

}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	// create fields
	obstacle_field := make([][]bool, len(lines))
	visited_field := make([][]bool, len(lines))

	var start_position [2]int

	// initialize fields
	for y, line := range lines {
		for x, marking := range strings.Split(line, "") {
			is_obstacle := false
			if marking == "#" {
				is_obstacle = true
			}
			obstacle_field[y] = append(obstacle_field[y], is_obstacle)

			is_guard := false
			if marking == "^" {
				is_guard = true
				start_position = [2]int{y, x}
			}
			visited_field[y] = append(visited_field[y], is_guard)
		}
	}

	// calculate guard path
	current_direction := 0
	current_position := start_position

	in_bounds := true
	for in_bounds {

		visited_field[current_position[0]][current_position[1]] = true
		var next_position [2]int

		for i := 0; i < 4; i++ {
			next_position = [2]int{current_position[0] + directions[current_direction][0], current_position[1] + directions[current_direction][1]}

			if !is_in_bounds(obstacle_field, next_position[0], next_position[1]) {
				in_bounds = false
				break
			}

			// true if next position is obstacle
			if obstacle_field[next_position[0]][next_position[1]] {
				// turn
				current_direction = (current_direction + 1) % 4
			} else {
				break
			}

		}

		current_position = next_position

	}

	return count_visits(visited_field)
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	// create fields
	obstacle_field := make([][]bool, len(lines))
	visited_field := make([][][]int, len(lines))

	var start_position [2]int

	// initialize fields
	for y, line := range lines {
		for x, marking := range strings.Split(line, "") {
			is_obstacle := false
			if marking == "#" {
				is_obstacle = true
			}
			obstacle_field[y] = append(obstacle_field[y], is_obstacle)

			if marking == "^" {
				start_position = [2]int{y, x}
			}
			visited_field[y] = append(visited_field[y], make([]int, 0))
		}
	}

	// calculate guard path
	current_direction := 0
	current_position := start_position
	plantable_obstacles := 0

	in_bounds := true
	for in_bounds {

		visited_field[current_position[0]][current_position[1]] = append(visited_field[current_position[0]][current_position[1]], current_direction)

		var next_position [2]int

		direction_for_circle := (current_direction + 1) % 4
		i_sidecheck := 0
		for true {

			to_check_coords := []int{current_position[0] + directions[direction_for_circle][0] * i_sidecheck, current_position[1] + directions[direction_for_circle][1] * i_sidecheck}
            if !is_in_bounds(obstacle_field, to_check_coords[0], to_check_coords[1]) {
                break
            }
            to_check := visited_field[to_check_coords[0]][to_check_coords[1]]
			if slices.Contains(to_check, direction_for_circle) {
				plantable_obstacles++
                break
			}

			i_sidecheck++
		}

		for i := 0; i < 4; i++ {
			next_position = [2]int{current_position[0] + directions[current_direction][0], current_position[1] + directions[current_direction][1]}

			if !is_in_bounds(obstacle_field, next_position[0], next_position[1]) {
				in_bounds = false
				break
			}

			// true if next position is obstacle
			if obstacle_field[next_position[0]][next_position[1]] {
				// turn
				current_direction = (current_direction + 1) % 4
			} else {
				break
			}

		}

		current_position = next_position

	}

	return plantable_obstacles
}
