package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var lines []int
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	replacer := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
	for scanner.Scan() {
		output := replacer.Replace(scanner.Text())
		seatid, _ := strconv.ParseInt(output, 2, 0)
		lines = append(lines, int(seatid))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Ints(lines)
	fmt.Println("Part 1:", lines[len(lines)-1])
	for i, v := range lines {
		if lines[i+1]-v > 1 {
			fmt.Println("Part 2:", v+1)
			break
		}
	}
}
