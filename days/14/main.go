package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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
		ans := part1(input, 103, 101)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

type point struct {
	y int
	x int
}
type robot struct {
	start  point
	change point
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func part1(input string, field_height, field_width int) int {
	lines := strings.Split(input, "\n")
	robots := make([]robot, 0)
	for _, line := range lines {
		robot := robot{}
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.start.x, &robot.start.y, &robot.change.x, &robot.change.y)
		if err != nil {
			panic("Input is not valid")
		}
		robots = append(robots, robot)
	}

	final_coords := make([]point, 0)

	times := 100

	for _, robot := range robots {
		point := point{
			y: mod(robot.start.y+(times*robot.change.y), field_height),
			x: mod(robot.start.x+(times*robot.change.x), field_width),
		}
		final_coords = append(final_coords, point)
	}

	q0, q1, q2, q3 := 0, 0, 0, 0

	y_center := field_height / 2
	x_center := field_width / 2

	for _, coord := range final_coords {
		y, x := coord.y, coord.x
		if y == y_center || x == x_center {
			continue
		}
		if y < y_center {
			if x < x_center {
				q0++
			} else {
				q1++
			}
		} else {
			if x < x_center {
				q2++
			} else {
				q3++
			}
		}

	}

	return q0 * q1 * q2 * q3
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	robots := make([]robot, 0)
	for _, line := range lines {
		robot := robot{}
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.start.x, &robot.start.y, &robot.change.x, &robot.change.y)
		if err != nil {
			panic("Input is not valid")
		}
		robots = append(robots, robot)
	}

	field_height := 103
	field_width := 101

	second := -1
	smallest_d := math.MaxInt

	// NOTE: I took inspiration from redit:
	// Someone posted this image on redit: https://imgur.com/a/KmBmhL0
	// This idea is that if they form a tree all points must be very dense.
	// In thise codw, I calculate an average point (x,y) (areth. mean).
	// Based on this, a delta-to seach point it summed (distrubution).
	// the second with the lowerst distrubtuion is used

	for times := range 100_000 {

		final_coords := make([]point, 0)

		for _, robot := range robots {
			point := point{
				y: mod(robot.start.y+(times*robot.change.y), field_height),
				x: mod(robot.start.x+(times*robot.change.x), field_width),
			}
			final_coords = append(final_coords, point)
		}

		sum_x := 0
		sum_y := 0
		for _, coord := range final_coords {
			sum_x += coord.x
			sum_y += coord.y
		}

		avg_x := sum_x / len(final_coords)
		avg_y := sum_y / len(final_coords)

		d_x := 0
		d_y := 0
		for _, coord := range final_coords {
			d_x += abs(avg_x - coord.x)
			d_y += abs(avg_y - coord.y)
		}

		d := d_x * d_y
		if d < smallest_d {
			smallest_d = d
			second = times
		}

	}

	return second
}
