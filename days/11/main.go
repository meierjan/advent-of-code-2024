package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math/big"
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
func part1(input string) int {
	stones := strings.Split(input, " ")

	rounds := 25

	result := 0
	for _, stone := range stones {
		result += simulate(stone, rounds )
	}
	return result
}

func simulate(stone string, rounds int) int {

	if rounds == 0 {
		return 1
	} else {
		result_stones := make([]string, 0)

		if stone == "0" {
			// If the stone is engraved with the number `0`, it is replaced by a stone engraved with the number `1`.
			result_stones = append(result_stones, "1")
		} else if len(stone)%2 == 0 {
			// If the stone is engraved with a number that has an *even* number of digits, it is replaced by *two stones*. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: `1000` would become stones `10` and `0`.)

			left := stone[:(len(stone) / 2)]
			right := strings.TrimLeft(stone[len(stone)/2:], "0")

            if right == "" {
                right = "0"
            }

			result_stones = append(result_stones, left, right)

		} else {
			// one of the other rules apply, the stone is replaced by a new stone; the old stone's number *multiplied by 2024* is engraved on the new stone.
			value, succ := new(big.Int).SetString(stone, 10)
			if !succ {
				panic("error")
			}
			new_stone := new(big.Int).Mul(value, big.NewInt(2024))
			result_stones = append(result_stones, new_stone.String())
		}
		result := 0

		for _, stone := range result_stones {
			result += simulate(stone, rounds-1)
		}

		return result
	}
}

func part2(input string) int {
    stones := strings.Split(input, " ")

	rounds := 75

	result := 0
	for _, stone := range stones {
		result += simulate(stone, rounds )
	}
	return result
}
