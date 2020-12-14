package bus

import (
	"fmt"
	"strconv"
)

func FindRemainder(closestTimestamp int, departure int) int {
	return closestTimestamp % departure
}

func FindClosestTimestamp(timestamp int, busId int) int {
	departure := 0
	for departure < timestamp {
		departure += busId
	}
	return departure
}

func SolveDay1(timestamp int, busIds []string) int {
	bestOption := 999
	options := map[int]int{}
	for _, busId := range busIds {
		if busId == "x" {
			continue
		}
		val, err := strconv.Atoi(busId)
		if err != nil {
			fmt.Println("Conversion error")
		}
		closestTimestamp := FindClosestTimestamp(timestamp, val)
		minutesToDeparture := FindRemainder(closestTimestamp, timestamp)
		options[minutesToDeparture] = val
		if minutesToDeparture < bestOption {
			bestOption = minutesToDeparture
		}
	}
	return options[bestOption] * bestOption
}

func GetRequiredOffsets(busIds []string) map[int]int {
	offsets := map[int]int{}
	for idx, busId := range busIds {
		if busId == "x" {
			continue
		}
		val, err := strconv.Atoi(busId)
		if err != nil {
			fmt.Println("Conversion error")
		}
		offsets[idx] = val
	}
	return offsets
}

func generateNextSeq(firstId int, length int) []int {
	seq := []int{}
	for i := 0; i < length; i++ {
		seq = append(seq, firstId+i)
	}
	return seq
}

// Again, shamelessly stolen (https://github.com/daniel-dara/advent-of-code/blob/master/2020/day13/part2.py)
func FindEarliestTimestampsInline(busIds []string) int {
	buses := []int{}
	mods := []int{}
	for idx, busId := range busIds {
		val, err := strconv.Atoi(busId)
		if err == nil {
			buses = append(buses, val)
			mods = append(mods, idx)
		}
	}
	time := 0
	step := 1
	// So basically, if I'm figuring this out correctly
	// Go through each bus and the number of minutes offset of zero that the time stamp needs to be
	// Take whatever time we're at plus that offset and keep moving forward by the step value if it doesn't
	// divide cleanly into the bus id
	// If it does divide cleanly into the bus id, for the next pair, we need to look at the _previous_ bus id
	// as the amount that it steps (because we have to be able to keep the previous values correct as well)
	for idx, busId := range buses {
		for (time+mods[idx])%busId != 0 {
			time += step
		}
		step *= busId
	}
	return time
}
