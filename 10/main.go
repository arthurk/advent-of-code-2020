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

func nextValid(start int, bag []int) []int {
	valid := []int{}
	for i := start + 1; i < len(bag); i++ {
		diff := bag[i] - bag[start]
		// fmt.Printf("%d - %d = %d\n", bag[i], bag[start], diff)
		if diff > 3 {
			break
		}
		valid = append(valid, bag[i])
	}
	return valid
}

func main() {
	// part1()
	bag := readInput()
	validMap := [][]int{}
	for i := 0; i < len(bag)-1; i++ {
		validMap = append(validMap, nextValid(i, bag))
	}

	// offset := 1
	// valid := []int{}
	for i := 0; i < len(validMap); i++ {
		for j := 0; j < len(validMap[i]); j++ {
			fmt.Println(validMap[i])
		}
	}
}
