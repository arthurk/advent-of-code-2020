package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1() int {
	lines := getNumbers()
	pre := 25
	for i := pre; i < len(lines)-1; i++ {
		iv := lines[i]
		// fmt.Println(iv)

		prediff := i - pre
		ok := false
		for j, k := range lines[i-pre : i] {
			for x := j; x <= pre-1; x++ {
				if x == j {
					continue
				}
				n1 := k
				n2 := lines[(x + prediff)]
				sum := n1 + n2
				if sum == iv {
					// fmt.Println("ok")
					ok = true
				}
				// fmt.Printf("i=%d j=%d prediff=%d x=%d | %s+%s=%d\n", i, j, prediff, x+prediff, k, lines[x+prediff], sum)
			}
		}

		if !ok {
			fmt.Println("not ok", iv)
			return iv
		}
	}
	return 0
}

func getNumbers() []int {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	numbers := []int{}
	for i := 0; i < len(lines)-1; i++ {
		n, _ := strconv.Atoi(lines[i])
		numbers = append(numbers, n)
	}
	return numbers
}

func part2(target int) {
	fmt.Println("target", target)
	numbers := getNumbers()
	for i := 0; i <= len(numbers); i++ {

		for j := i + 1; j < len(numbers); j++ {
			sum := numbers[i]
			smallest := numbers[i]
			largest := numbers[j]
			for k := i + 1; k <= j; k++ {
				//res := numbers[i] + numbers[k]
				if numbers[k] < smallest {
					smallest = numbers[k]
				}
				if numbers[k] > largest {
					largest = numbers[k]
				}
				sum += numbers[k]
				// fmt.Printf("%d + %d = %d | sum %d\n", numbers[i], numbers[k], res, sum)
				if sum == target {
					// fmt.Println("OK", numbers[i], numbers[k], numbers[j]+numbers[k])
					fmt.Println("smallest, largest, sum", smallest, largest, smallest+largest)
					return
				}
			}
		}
	}
}

func main() {
	res := part1()
	part2(res)
}
