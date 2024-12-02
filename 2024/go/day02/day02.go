package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "sort"
	"strconv"
	"time"
)

func main() {
	start_time := time.Now()

	input_file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input_file.Close()

	// read the file line by line
	scanner := bufio.NewScanner(input_file)

	reports := make([][]int, 0)
	for scanner.Scan() {
		report_text := strings.Split(scanner.Text(), " ")

		report := make([]int, len(report_text))
		for i, level_text := range report_text {
			level, err := strconv.Atoi(level_text)
			if err != nil {
				panic(err)
			}
			report[i] = level
		}

		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1(reports)
	part2(reports)

	end_time := time.Since(start_time)
	fmt.Printf("Day 01 took %v milliseconds (%v microseconds)\n",
		end_time.Milliseconds(), end_time.Microseconds())
}

func is_monotonic_banded(report []int, max_diff int) bool {
	if report[0] == report[1] {
		return false
	}

	is_increasing := report[1] > report[0]
	prev_level := report[0]

	for _, level := range report[1:] {
		if prev_level == level {
			return false
		}

		if is_increasing {
			if level < prev_level {
				return false
			} else if level-prev_level > max_diff {
				return false
			}
		} else {
			if level > prev_level {
				return false
			} else if prev_level-level > max_diff {
				return false
			}
		}

		prev_level = level
	}

	return true
}

func part1(reports [][]int) {
	num_safe := 0
	for _, report := range reports {
		if is_monotonic_banded(report, 3) {
			num_safe++
		}
	}

	fmt.Println("There are ", num_safe, " safe reports.")
}

func remove_index(report []int, index int) []int {
	new_report := make([]int, len(report)-1)
	j := 0
	for i := 0; i < len(report); i++ {
		if i == index {
			continue
		}

		new_report[j] = report[i]
		j++
	}

	return new_report
}

func is_safe_dampened(report []int, max_diff int) bool {
	for i := range report {
		if is_monotonic_banded(remove_index(report, i), max_diff) {
			return true
		}
	}

	return false
}

func part2(reports [][]int) {
	num_safe := 0
	for _, report := range reports {
		if is_safe_dampened(report, 3) {
			num_safe++
		}
	}

	fmt.Println("There are ", num_safe, " safe dampened reports.")
}
