package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readLists() ([]int, []int) {
	cwd, err := os.Getwd();
	check(err)

	path := filepath.Join(cwd, "inputs", "01.txt")

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	var left []int
	var right []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		left = append(left, a)
		right = append(right, b)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return left, right
}

func part1() int {
	left, right := readLists()

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := range left {
		if left[i] > right[i] {
			sum += left[i] - right[i]
		} else {
			sum += right[i] - left[i]
		}
	}

	return sum
}

func part2() int {
	left, right := readLists()

	scores := make(map[int]int)
	for _, lVal := range left {
		_, exists := scores[lVal]
		if exists {
			fmt.Printf("Exists in map: %d\n", lVal)
			continue
		}

		for _, rVal := range right {
			if lVal == rVal {
				scores[lVal] += 1
			}
		}
	}

	totalScore := 0
	for k, v := range scores {
		totalScore += k * v
	}

	return totalScore
}

func main() {
	fmt.Printf("Part 1: %d\n", part1()) // 1222801
	fmt.Printf("Part 2: %d\n", part2()) // 22545250
}
