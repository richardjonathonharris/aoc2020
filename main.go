package main

import (
	"fmt"
	"github.com/richardjonathonharris/aoc2020/day1"
	"github.com/richardjonathonharris/aoc2020/day2"
	"github.com/richardjonathonharris/aoc2020/day3"
	"github.com/richardjonathonharris/aoc2020/day4"
	"github.com/richardjonathonharris/aoc2020/day5"
	"github.com/richardjonathonharris/aoc2020/day6"
	"github.com/richardjonathonharris/aoc2020/day7"
	"github.com/richardjonathonharris/aoc2020/day8"
	"github.com/richardjonathonharris/aoc2020/day9"
	"github.com/richardjonathonharris/aoc2020/day10"
	"github.com/richardjonathonharris/aoc2020/utils"
	"sort"
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

func day5() {
	fmt.Println("Day 5!")
	day5rawdata := strings.Split(utils.PlaintextFromFile("./day5/data.txt"), "\n")
	biggestId := 0
	var seatIds []int
	for _, direction := range day5rawdata {
		estimatedPosition := seats.FindSeats(direction, [2]int{0, 127}, [2]int{0, 7})
		row, column, seatId, err := estimatedPosition.GetSeatInformation()
		fmt.Sprintf("Row %d, Seat %d", row, column)
		if err != nil {
			fmt.Println("We could not verify the seat!")
			break
		}
		if seatId > biggestId {
			biggestId = seatId
		}
		seatIds = append(seatIds, seatId)
	}
	fmt.Println("Biggest seat id ", biggestId)
	sort.Ints(seatIds)
	for idx, seatVal := range seatIds {
		if idx == 0 || idx == len(seatIds)-1 {
			continue
		}
		if seatIds[idx+1] == seatVal+2 {
			fmt.Println(fmt.Sprintf("Possible seat val found! Idx %d = seat id %d and Idx %d = seat id %d!", idx, seatVal, idx+1, seatIds[idx+1]))
		}
	}
}

func day6() {
	fmt.Println("Day 6!")
	day6rawdata := strings.Split(utils.PlaintextFromFile("./day6/data.txt"), "\n\n")
	sumQuestions := 0
	sumAllAnsweredYes := 0
	for _, form := range day6rawdata {
		sumQuestions += customs.CountUniqueLetters(form)
		sumAllAnsweredYes += customs.CountLettersAllAnswer(form)
	}
	fmt.Println("Sum of questions answered yes by groups: ", sumQuestions)
	fmt.Println("Sum of questions answered yes by each person in groups: ", sumAllAnsweredYes)
}

func day7() {
	fmt.Println("Day 7!")
	day7rawdata := strings.Split(utils.PlaintextFromFile("./day7/data.txt"), "\n")
	bagMap := bags.BagMap{}
	for _, bagText := range day7rawdata {
		newBag := bags.CreateNewBag(bagText)
		bagMap[newBag.Color] = newBag
	}
	bagMapWithContainedBy := bags.BagMap{}
	for color, bag := range bagMap {
		containedBy := bags.AddAllContainingBags(color, bagMap)
		bagMapWithContainedBy[color] = bags.Bag{Color: color, Contains: bag.Contains, ContainedBy: containedBy}
	}

	bagsContainingShowyBag := map[string]int{}
	bags.ListAllContainingBags("shiny gold", bagMapWithContainedBy, bagsContainingShowyBag)
	fmt.Println("Number of bags that could ultimately hold a shiny gold bag ", len(bagsContainingShowyBag))
	showyBagContains := map[string]int{}
	bags.CountAllBagsThatBagContains("shiny gold", bagMapWithContainedBy, showyBagContains)
	counter := 0
	for _, val := range showyBagContains {
		counter += val
	}
	fmt.Println(fmt.Sprintf("Number of bags that a shiny gold bag holds %+v", counter))
}

func day8() {
	fmt.Println("Day 8!")
	day8rawdata := strings.Split(utils.PlaintextFromFile("./day8/data.txt"), "\n")
	codebook := console.BuildCodebook(day8rawdata)
	accValue, _ := console.FindRepeatedInstruction(codebook)
	fmt.Println("Accumulator value right before repeated instruction", accValue)
	// Brute forcing this but you know? I'm fine with that right now
	for i := 0; i < len(codebook); i++ {
		// Specifically create a new codebook
		newCodebook := console.BuildCodebook(day8rawdata)
		for k, v := range newCodebook {
			if k == i && codebook[i].Op == console.Jump {
				newCodebook[k] = &console.Instruction{Op: console.NoOp, Value: codebook[i].Value}
			} else if k == i && codebook[i].Op == console.NoOp {
				newCodebook[k] = &console.Instruction{Op: console.Jump, Value: codebook[i].Value}
			} else {
				newCodebook[k] = v
			}
		}
		accValue, err := console.FindRepeatedInstruction(newCodebook)
		if err != nil {
			fmt.Println("Accumulator ", accValue, "by changing value at idx ", i)
		}
	}
}

func day9() {
	fmt.Println("Day 9!")
	day9rawdata := strings.Split(utils.PlaintextFromFile("./day9/data.txt"), "\n")
	var day9data []int
	for _, item := range day9rawdata {
		newVal, err := strconv.Atoi(item)
		if err != nil {
			panic("AHHHH IT DID NOT CONVERT")
		}
		day9data = append(day9data, newVal)
	}
	firstInvalid, err := xmas.FindFirstInvalidIndex(day9data, 25, 26)
	if err != nil {
		fmt.Println("Boy howdy, that's an error!")
	}
	fmt.Println("First invalid value is ", day9data[firstInvalid])
	firstBad, lastBad := xmas.FindSmallLargeSumValues(day9data, firstInvalid)
	fmt.Println("First bad value of sum ", firstBad, "last bad value of sum ", lastBad)
	fmt.Println("Their sum is: ", firstBad+lastBad)
}

func day10() {
	fmt.Println("Day 10!")
	day10rawdata := strings.Split(utils.PlaintextFromFile("./day10/data.txt"), "\n")
	day10data := []int{0}
	maxJoltage := 0 // original plug
	for _, item := range day10rawdata {
		newVal, err := strconv.Atoi(item)
		if err != nil {
			panic("AHHHH IT DID NOT CONVERT")
		}
		if newVal > maxJoltage {
			maxJoltage = newVal
		}
		day10data = append(day10data, newVal)
	}
	day10data = append(day10data, maxJoltage + 3) // final plug
	sort.Ints(day10data)
	mapJoltDiffs, _ := joltage.GetMapOfJoltDiffs(day10data)
	fmt.Printf("JoltDiffs %+v\n", mapJoltDiffs)
	product :=  mapJoltDiffs["1"] * mapJoltDiffs["3"]
	fmt.Println("Product between one and three would be", product)
	possibleRoutes := []string{}
	joltage.DetermineRoutes(&possibleRoutes, []string{"0"}, day10data, maxJoltage)
	fmt.Printf("Found %d possible routes\n", possibleRoutes)
}

func main() {
	day1()
	fmt.Println("\n\n------------------")
	day2()
	fmt.Println("\n\n------------------")
	day3()
	fmt.Println("\n\n------------------")
	day4()
	fmt.Println("\n\n------------------")
	day5()
	fmt.Println("\n\n------------------")
	day6()
	fmt.Println("\n\n------------------")
	day7()
	fmt.Println("\n\n------------------")
	day8()
	fmt.Println("\n\n------------------")
	day9()
	fmt.Println("\n\n------------------")
	day10()
}
