package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instr struct {
	Op  string
	Arg int
}
type instrSet []instr

func run(lines instrSet) (int, error) {
	acc := 0
	seen := make(map[int]bool)
	for i := 0; i < len(lines); i++ {
		_, ran := seen[i]
		if ran {
			return acc, errors.New("already ran")
		}
		seen[i] = true

		op := lines[i].Op
		arg := lines[i].Arg
		switch op {
		case "jmp":
			i += arg - 1
		case "acc":
			acc += arg
		}
	}
	return acc, nil
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	// parse input to bootcode
	bootcode := make(instrSet, 0)
	for i := 0; i < len(lines)-1; i++ {
		parts := strings.Fields(lines[i])
		arg, _ := strconv.Atoi(parts[1])
		bootcode = append(bootcode, instr{Op: parts[0], Arg: arg})
	}

	acc, err := run(bootcode)
	if err != nil {
		fmt.Println("Part 1:", acc)
	}

	// change jmp to nop
	for i := 0; i <= len(bootcode); i++ {
		op := bootcode[i].Op
		switch op {
		case "jmp":
			bootcode[i].Op = "nop"
		case "nop":
			bootcode[i].Op = "jmp"
		}
		acc, err := run(bootcode)
		bootcode[i].Op = op
		if err == nil {
			fmt.Println("Part 2:", acc)
			break
		}
	}
}
