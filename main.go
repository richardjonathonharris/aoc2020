package main

import (
	"fmt"
	"github.com/richardjonathonharris/aoc2020/day1"
	"github.com/richardjonathonharris/aoc2020/day2"
	"github.com/richardjonathonharris/aoc2020/day3"
	"github.com/richardjonathonharris/aoc2020/day4"
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

func day3() {
	fmt.Println("Day 3!")
	day3rawdata := strings.Split(utils.PlaintextFromFile("./day3/data.txt"), "\n")
	instructions := seeker.Direction{Right: 3, Down: 1}
	trees := seeker.Day3HelperFunction(instructions, day3rawdata, "#")
	fmt.Println("Number of trees hit: ", trees)
	fmt.Println("Part 2!")
	setOfInstructions := []seeker.Direction{
		{Right: 1, Down: 1},
		{Right: 5, Down: 1},
		{Right: 7, Down: 1},
		{Right: 1, Down: 2},
		{Right: 3, Down: 1},
	}
	var treeValues []int
	for _, inst := range setOfInstructions {
		trees := seeker.Day3HelperFunction(inst, day3rawdata, "#")
		treeValues = append(treeValues, trees)
	}
	fmt.Println(treeValues)
	product := 1
	for _, tree := range treeValues {
		product *= tree
	}
	fmt.Println("Product of trees hit", product)
}

func day4() {
	fmt.Println("Day 4!")
	day4rawdata := strings.Split(utils.PlaintextFromFile("./day4/data.txt"), "\n\n")
	countValid := 0
	countValidAndGoodFields := 0
	for _, pass := range day4rawdata {
		if passporter.ValidatePassport(pass, false) {
			countValid += 1
		}
		if passporter.ValidatePassport(pass, true) {
			countValidAndGoodFields += 1
		}
	}
	fmt.Println("Number of valid passports: ", countValid)
	fmt.Println("Number of valid passports with good fields: ", countValidAndGoodFields)
}

func main() {
	day1()
	fmt.Println("\n\n------------------")
	day2()
	fmt.Println("\n\n------------------")
	day3()
	fmt.Println("\n\n------------------")
	day4()
}
