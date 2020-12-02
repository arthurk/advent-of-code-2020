package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// FindSumTo2020 returns a slice of ints that contains two numbers that add up to 2020 (Part 1)
func FindSumTo2020(lines []int) []int {
	for x, i := range lines {
		for _, j := range lines[x:] {
			if i+j == 2020 {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// FindProduct2020 returns a slice of ints that contains three numbers that add up to 2020 (Part 2)
func FindProduct2020(lines []int) []int {
	for x, i := range lines {
		for y, j := range lines[x:] {
			for _, k := range lines[y:] {
				if i+j+k == 2020 {
					return []int{i, j, k}
				}
			}
		}
	}
	return []int{}
}

func main() {
	// read input from file
	numbers := []int{}
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, num)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1:
	sumNumbers := FindSumTo2020(numbers)
	fmt.Printf("Part 1: Numbers: %d Sum: %d\n", sumNumbers, sumNumbers[0]*sumNumbers[1])

	// Part 2:
	productNumbers := FindProduct2020(numbers)
	fmt.Printf("Part 2: Numbers: %d Product: %d\n", productNumbers, productNumbers[0]*productNumbers[1]*productNumbers[2])
}
