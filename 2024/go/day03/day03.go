package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	start_time := time.Now()

	raw_bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	text := string(raw_bytes)

	part1(text)
	part2(text)

	end_time := time.Since(start_time)
	fmt.Printf("Day 01 took %v milliseconds (%v microseconds)\n",
		end_time.Milliseconds(), end_time.Microseconds())
}

func part1(text string) {
	r := regexp.MustCompile("mul\\(([0-9]+),([0-9]+)\\)")
	matches := r.FindAllStringSubmatch(text, -1)

	total := 0
	for _, match := range matches {
		num1, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		total += num1 * num2
	}

	fmt.Println("Part 1 sum: ", total)
}

func is_enabled(match_indices []int, do_indices [][]int, dont_indices [][]int) bool {
	if match_indices[0] < dont_indices[0][0] {
		return true
	}

	// find most recent do
	last_do := do_indices[len(do_indices)-1][0]
	for i := len(do_indices) - 1; i >= 0; i-- {
		if do_indices[i][0] < match_indices[0] {
			last_do = do_indices[i][0]
			break
		}
	}

	// find most recent dont
	last_dont := dont_indices[len(dont_indices)-1][0]
	for i := len(dont_indices) - 1; i >= 0; i-- {
		if dont_indices[i][0] < match_indices[0] {
			last_dont = dont_indices[i][0]
			break
		}
	}

	return last_do > last_dont && match_indices[0] > last_do
}

func part2(text string) {
	r := regexp.MustCompile("mul\\(([0-9]+),([0-9]+)\\)")
	matches := r.FindAllStringSubmatch(text, -1)
	matches_indices := r.FindAllStringSubmatchIndex(text, -1)

	r_do := regexp.MustCompile("do\\(\\)")
	do_indices := r_do.FindAllStringIndex(text, -1)

	r_dont := regexp.MustCompile("don't\\(\\)")
	dont_indices := r_dont.FindAllStringIndex(text, -1)

	total := 0
	for i := 0; i < len(matches); i++ {
		if !is_enabled(matches_indices[i], do_indices, dont_indices) {
			continue
		}

		num1, err := strconv.Atoi(matches[i][1])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(matches[i][2])
		if err != nil {
			panic(err)
		}
		total += num1 * num2
	}

	fmt.Println("Part 2 sum: ", total)
}
