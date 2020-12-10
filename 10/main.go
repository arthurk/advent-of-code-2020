package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readInput() []int {
	var lines []int
	file, _ := os.Open("input-test-1.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, a)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Ints(lines)
	return lines
}

func part1() {
	bag := readInput()
	diffCount := make(map[int]int)
	for i := range bag {
		if i == len(bag)-1 {
			break
		}

		diff := bag[i+1] - bag[i]
		diffCount[diff]++
	}
	for k, v := range diffCount {
		fmt.Printf("%d differences of %d jolt\n", v+1, k)
	}

	fmt.Println("Part 1:", (diffCount[1]+1)*(diffCount[3]+1))
}

func main() {
	part1()
}
