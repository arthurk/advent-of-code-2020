package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type listItem struct {
	num1     int
	num2     int
	letter   string
	password string
}

func (l listItem) IsValidPart2() bool {
	p1 := string(l.password[l.num1-1])
	p2 := string(l.password[l.num2-1])
	return (p1 != p2) && (p1 == l.letter || p2 == l.letter)
}

func (l listItem) IsValidPart1() bool {
	count := strings.Count(l.password, l.letter)
	return count >= l.num1 && count <= l.num2
}

func newListItem(line string) *listItem {
	r := regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)$`)
	matches := r.FindStringSubmatch(line)
	minCount, _ := strconv.Atoi(matches[1])
	maxCount, _ := strconv.Atoi(matches[2])
	return &listItem{
		num1:     minCount,
		num2:     maxCount,
		letter:   matches[3],
		password: matches[4],
	}
}

func main() {
	p1ValidCount := 0
	p2ValidCount := 0

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		item := newListItem(scanner.Text())
		if item.IsValidPart1() {
			p1ValidCount++
		}
		if item.IsValidPart2() {
			p2ValidCount++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1 Valid passwords:", p1ValidCount)
	fmt.Println("Part 2 Valid passwords:", p2ValidCount)
}
