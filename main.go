package main

import (
	"fmt"
	"github.com/richardjonathonharris/aoc2020/day1"
	"github.com/richardjonathonharris/aoc2020/day2"
	"github.com/richardjonathonharris/aoc2020/utils"
	"strconv"
	"strings"
)

func day1() {
	fmt.Println("Day 1!")
	day1rawdata := strings.Split(utils.PlaintextFromFile("./day1/data.txt"), "\n")
	var day1data []int
	for _, val := range day1rawdata {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Your conversion did not work!")
		}
		day1data = append(day1data, intVal)
	}
	vals, err := find2020Sum.FindFirst2ValuesEqualValue(day1data, 2020)
	if err != nil {
		fmt.Println("Could not find summation!")
	}
	fmt.Println("Product: ", vals[0]*vals[1])
	fmt.Println("Finding product for 3 entries")
	product, err := find2020Sum.FindMultipleEqualValue(day1data, 2020, 3)
	if err != nil {
		fmt.Println("Couldn't find a match!")
	}
	fmt.Println("Product for the second part:", product)
}

func day2() {
	fmt.Println("Day 2!")
	day2rawdata := strings.Split(utils.PlaintextFromFile("./day2/data.txt"), "\n")
	var day2data [][]string
	for _, val := range day2rawdata {
		day2data = append(day2data, strings.Split(val, ":"))
	}
	countSuccessfulPassword := 0
	countFailedPassword := 0
	countPasswords := 0
	for _, val := range day2data {
		instr, err := password.ParseInstructions(val[0])
		if err != nil {
			fmt.Println("Could not parse password")
		}
		result, err := password.CheckPassword(instr, val[1])
		if err != nil {
			fmt.Println("Could not check password")
		}
		if result == true {
			countSuccessfulPassword += 1
		} else {
			countFailedPassword += 1
		}
		countPasswords += 1
	}
	fmt.Println("Successful passwords:", countSuccessfulPassword)
	fmt.Println("Unsuccessful passwords:", countFailedPassword)
	fmt.Println("Total passwords:", countPasswords)

	fmt.Println("Part 2!")
	countExactMatch := 0
	countFailedMatch := 0
	countExactPasswords := 0
	for _, val := range day2data {
		instr, err := password.ParseInstructions(val[0])
		if err != nil {
			fmt.Println("Could not parse password")
		}
		result, err := password.CheckPasswordExactOne(instr, strings.TrimSpace(val[1]))
		if err != nil {
			fmt.Println("Could not check password")
		}
		if result == true {
			countExactMatch += 1
		} else {
			countFailedMatch += 1
		}
		countExactPasswords += 1
	}
	fmt.Println("Successful passwords:", countExactMatch)
	fmt.Println("Unsuccessful passwords:", countFailedMatch)
	fmt.Println("Total passwords:", countExactPasswords)
}

func main() {
	day1()
	fmt.Println("\n\n------------------")
	day2()
	fmt.Println("\n\n------------------")
}
