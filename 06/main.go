package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part1(gg []string) {
	total := 0
	for _, a := range gg {
		gc := map[rune]bool{}
		for _, c := range a {
			_, ok := gc[c]
			if !ok {
				gc[c] = true
				total++
			}
		}

	}
	fmt.Println("part 1:", total)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	// each entry are answers from one group
	gg := []string{}
	count := 0
	buf := ""

	charCounts := []map[rune]int{}
	charCount := map[rune]int{}

	p2total := 0

	for _, line := range lines {
		// fmt.Printf("new line (len %d) (pcount %d): %s\n", len(line), count, line)

		// number of person in each group
		count++

		for _, c := range line {
			charCount[c] += 1
		}

		if len(line) == 0 {
			gg = append(gg, buf)
			buf = ""
			charCounts = append(charCounts, charCount)
			// part 2
			for _, z := range charCount {
				res := 0.0
				res = float64(z) / float64(count-1)
				if res == 1.0 {
					p2total++
				}
			}

			// reset
			charCount = map[rune]int{}
			count = 0
		}

		buf += line
	}

	// fmt.Println("\n", charCounts)
	fmt.Printf("part 2: %d\n", p2total)

	// iterate over each char and mark as seen
	part1(gg)
}
