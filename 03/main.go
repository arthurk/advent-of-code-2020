package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput() []string {
	var lines []string
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func countTrees(lines []string, right, down int) int {
	treeCount := 0
	charCount := 1
	for i := down; i <= len(lines)-1; i += down {
		charIdx := (charCount * right) % len(lines[i])
		char := string(lines[i][charIdx])
		//fmt.Printf("lines[%d][%d] = %s\n", i, charIdx, char)
		if char == "#" {
			treeCount++
		}
		charCount++
	}
	return treeCount
}

func main() {
	// slope format: {right, down}
	// challenge part one is slope {3, 1}
	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	lines := readInput()
	mul := 1
	for _, s := range slopes {
		treeCount := countTrees(lines, s[0], s[1])
		mul *= treeCount
		fmt.Printf("Right: %d, Down: %d, Trees: %d\n", s[0], s[1], treeCount)
	}

	fmt.Println("Multiplied number of all trees:", mul)
}
