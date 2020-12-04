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

type passport map[string]string

func readFile() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	entries := []string{}
	scanner := bufio.NewScanner(file)
	line := ""
	for scanner.Scan() {
		l := scanner.Text()
		//fmt.Printf("read \"%s\" len %d\n", l, len(l))
		if len(l) != 0 {
			//fmt.Printf("append %s to %s\n", l, line)
			line += l + " "
			continue
		}

		// empty line is start of new passport
		if len(l) == 0 {
			// trim empty space from prev line
			line = strings.TrimSpace(line)

			// append line to all entries
			entries = append(entries, line)

			// reset line
			line = ""
			continue
		}
	}
	line = strings.TrimSpace(line)
	entries = append(entries, line)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return entries
}

func isValidByr(v string) bool {
	// (Birth Year)
	// four digits; at least 1920 and at most 2002
	iv, _ := strconv.Atoi(v)
	if len(v) == 4 && iv >= 1920 && iv <= 2002 {
		//fmt.Println("valid", v)
		return true
	}
	// fmt.Println("invld", v)
	return false
}

func isValidIyr(v string) bool {
	// (Issue Year)
	// four digits; at least 2010 and at most 2020.
	iv, _ := strconv.Atoi(v)
	if len(v) == 4 && iv >= 2010 && iv <= 2020 {
		// fmt.Println("valid", v)
		return true
	}
	// fmt.Println("invld", v)
	return false
}

func isValidEyr(v string) bool {
	// (Expiration Year)
	// four digits; at least 2020 and at most 2030.
	iv, _ := strconv.Atoi(v)
	if len(v) == 4 && iv >= 2020 && iv <= 2030 {
		// fmt.Println("valid", v)
		return true
	}
	// fmt.Println("invld", v)
	return false
}

func isValidHgt(v string) bool {
	// (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	cm := strings.Split(v, "cm")
	if len(cm) == 2 {
		iv, err := strconv.Atoi(cm[0])
		if err != nil {
			log.Fatal(err)
		}
		if iv >= 150 && iv <= 193 {
			//fmt.Println("cm OK", iv)
			return true
		}
	}
	in := strings.Split(v, "in")
	if len(in) == 2 {
		iv, err := strconv.Atoi(in[0])
		if err != nil {
			log.Fatal(err)
		}
		if iv >= 59 && iv <= 76 {
			return true
		}
	}
	// fmt.Println("Invalid", v)
	return false
}

func isValidHcl(v string) bool {
	// (Hair Color)
	// a # followed by exactly six characters 0-9 or a-f.
	matched, err := regexp.Match(`^#[a-f0-9]{6}$`, []byte(v))
	if err != nil {
		log.Fatal(err)
	}
	if matched {
		// fmt.Println("valid", v)
		return true
	}
	// fmt.Println("invalid", v)
	return false
}

func isValidEcl(v string) bool {
	// (Eye Color)
	// exactly one of: amb blu brn gry grn hzl oth.
	matched, err := regexp.Match(`^(amb|blu|brn|gry|grn|hzl|oth)$`, []byte(v))
	if err != nil {
		log.Fatal(err)
	}
	if matched {
		//fmt.Println("valid", v)
		return true
	}
	// fmt.Println("invalid", v)
	return false
}

func isValidPid(v string) bool {
	// (Passport ID)
	// a nine-digit number, including leading zeroes.
	matched, err := regexp.Match(`^[0-9]{9}$`, []byte(v))
	if err != nil {
		log.Fatal(err)
	}
	if matched {
		// fmt.Println("valid", v)
		return true
	}
	// fmt.Println("invld", v)
	return false
}

func checkRequiredFields(p passport) bool {
	reqFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, f := range reqFields {
		if _, ok := p[f]; !ok {
			return false
		}
	}
	return true
}

func main() {
	// read and parse data
	entries := readFile()
	var passports []passport
	for _, e := range entries {
		//fmt.Printf("entry %d: %s\n\n", i, e)
		p := passport{}
		parts := strings.Split(e, " ")
		for _, part := range parts {
			kv := strings.Split(part, ":")
			p[kv[0]] = kv[1]
		}
		passports = append(passports, p)
	}

	// validate data
	part1Valid := 0
	part2Valid := 0
	for _, p := range passports {
		// Part 1: validate if required fields exist
		if !checkRequiredFields(p) {
			continue
		}
		part1Valid++

		// Part 2: validate field values
		if !isValidByr(p["byr"]) {
			continue
		}
		if !isValidIyr(p["iyr"]) {
			continue
		}
		if !isValidEyr(p["eyr"]) {
			continue
		}
		if !isValidHgt(p["hgt"]) {
			continue
		}
		if !isValidHcl(p["hcl"]) {
			continue
		}
		if !isValidEcl(p["ecl"]) {
			continue
		}
		if !isValidPid(p["pid"]) {
			continue
		}

		part2Valid++
	}

	fmt.Println("Part 1 valid passports:", part1Valid)
	fmt.Println("Part 2 valid passports:", part2Valid)
}
