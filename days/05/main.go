package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
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
	sum := 0

	before := make(map[int][]int)

	empty_line := -1
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			empty_line = i
			break
		}

		mappings := strings.Split(line, ",")

		for _, rawMapping := range mappings {
			fromTo := strings.Split(rawMapping, "|")
			from, from_err := strconv.Atoi(fromTo[0])
			to, to_err := strconv.Atoi(fromTo[1])

			if from_err != nil || to_err != nil {
				panic("invalid format")
			}

			before[from] = append(before[from], to)
		}

	}

	for i := empty_line + 1; i < len(lines); i++ {
		pages := strings.Split(lines[i], ",")
		pages_i := []int{}
        pageSet := make(map[int]bool)

		for _, page := range pages {
			page_i, err := strconv.Atoi(page)
			if err != nil {
				panic("invalid page format")
			}
			pages_i = append(pages_i, page_i)
            pageSet[page_i] = true
		}

		is_sorted := slices.IsSortedFunc(pages_i, func(a, b int) int {
			
            before_a, has_key := before[a]

            if a == b {
                return 0
            }


            // -1 -> a < b
            // 0 -> a == b 
            // +1 -> a > b
            // p1|p2 -> p1 < p2 -> -1  

            if has_key {
                is_before := slices.Contains(before_a, b)
                if is_before {
                    return -1
                }
            }


			return 1
		})

		if is_sorted {
			middle_index := len(pages_i) / 2
			sum += pages_i[middle_index]
		}

	}

	return sum
}

func part2(input string) int {	sum := 0

	before := make(map[int][]int)

	empty_line := -1
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			empty_line = i
			break
		}

		mappings := strings.Split(line, ",")

		for _, rawMapping := range mappings {
			fromTo := strings.Split(rawMapping, "|")
			from, from_err := strconv.Atoi(fromTo[0])
			to, to_err := strconv.Atoi(fromTo[1])

			if from_err != nil || to_err != nil {
				panic("invalid format")
			}

			before[from] = append(before[from], to)
		}

	}

	for i := empty_line + 1; i < len(lines); i++ {
		pages := strings.Split(lines[i], ",")
		pages_i := []int{}
        pageSet := make(map[int]bool)

		for _, page := range pages {
			page_i, err := strconv.Atoi(page)
			if err != nil {
				panic("invalid page format")
			}
			pages_i = append(pages_i, page_i)
            pageSet[page_i] = true
		}

        sort_func := func(a, b int) int {
			
            before_a, has_key := before[a]

            if a == b {
                return 0
            }


            // -1 -> a < b
            // 0 -> a == b 
            // +1 -> a > b
            // p1|p2 -> p1 < p2 -> -1  

            if has_key {
                is_before := slices.Contains(before_a, b)
                if is_before {
                    return -1
                }
            }


			return 1
		}

		is_sorted := slices.IsSortedFunc(pages_i, sort_func)
        

		if !is_sorted {

            slices.SortFunc(pages_i, sort_func)


			middle_index := len(pages_i) / 2
			sum += pages_i[middle_index]
		}

	}

	return sum
}
