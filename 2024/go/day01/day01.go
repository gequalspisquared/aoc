package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

	left_ids := make([]int, 0)
	right_ids := make([]int, 0)
	for scanner.Scan() {
		ids := strings.Fields(scanner.Text())

		left_id, err := strconv.Atoi(ids[0])
		if err != nil {
			panic(err)
		}
		right_id, err := strconv.Atoi(ids[1])
		if err != nil {
			panic(err)
		}

		left_ids = append(left_ids, left_id)
		right_ids = append(right_ids, right_id)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// sorting the lists is needed for part 1 and helps speed up part 2
	sort.Ints(left_ids)
	sort.Ints(right_ids)

	part1(left_ids, right_ids)
	part2(left_ids, right_ids)

	end_time := time.Since(start_time)

	fmt.Printf("Day 01 took %v milliseconds (%v microseconds)\n",
		end_time.Milliseconds(), end_time.Microseconds())
}

func part1(left_ids []int, right_ids []int) {
	total_distance := 0
	for i := 0; i < len(left_ids); i++ {
		distance := left_ids[i] - right_ids[i]
		if distance < 0 {
			distance *= -1
		}
		total_distance += distance
	}

	fmt.Println("Total distance: ", total_distance)
}

func part2(left_ids []int, right_ids []int) {
	total_similarity := 0
	for _, left_id := range left_ids {
		for _, right_id := range right_ids {
			if left_id == right_id {
				total_similarity += left_id
			} else if right_id > left_id {
				break
			}
		}
	}

	fmt.Println("Total similarity: ", total_similarity)
}
