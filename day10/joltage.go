package joltage

import (
	"errors"
	"fmt"
	"strings"
	"strconv"
	"sync"
	"time"
)

type JoltDiff int

const (
	One   JoltDiff = 1
	Two            = 2
	Three          = 3
)

type JoltRecord map[string]int

func GetNextJoltDiff(currJolt int, nextJolt int) (JoltDiff, error) {
	if nextJolt - currJolt < 1 || nextJolt - currJolt > 3 {
		return One, errors.New("Jolt Difference is too big")
	}
	return JoltDiff(nextJolt - currJolt), nil
}

func GetPossibleNextJoltDiffs(currJolt int, data[]int) []int {
	goodJolts := []int{}
	for _, nextJolt := range data {
		if nextJolt - currJolt < 1 || nextJolt - currJolt > 3 {
			continue
		} else {
			goodJolts = append(goodJolts, nextJolt)
		}
	}
	return goodJolts
}

func DetermineRoute(records *[]string, currentJump string, currJolt int, data[]int) {
	var recordedJump string
	if len(currentJump) == 0 {
		recordedJump = "0"
	} else {
		recordedJump = fmt.Sprintf("%s-%d", currentJump, currJolt)
	}
	nextJolts := GetPossibleNextJoltDiffs(currJolt, data)
	if len(nextJolts) == 0 {
		// if there are no possible next jumps
		*records = append(*records, recordedJump)
	}
	for _, jump := range nextJolts {
		fmt.Println(recordedJump)
		DetermineRoute(records, recordedJump, jump, data)
	}
}

func generateNextRecords(ch chan string, wg *sync.WaitGroup, choice string, currentChoices []string, data []int) {
	defer wg.Done()
	var splitChoice string
	if len(currentChoices) == 1 {
		// we're just starting out
		splitChoice = "0"
	} else {
		splitChoices := strings.Split(choice, "-")
		splitChoice = splitChoices[len(splitChoices) - 1]
	}
	currentJump, err := strconv.Atoi(splitChoice)
	if err != nil {
		fmt.Println("Couldn't parse", err, splitChoice)
	}
	nextJumps := GetPossibleNextJoltDiffs(currentJump, data)
	for _, jump := range nextJumps {
		ch <- fmt.Sprintf("%s-%d", choice, jump)
	}
}

func generateRecordsToSend(ch chan string, writech chan string, wg *sync.WaitGroup, rec string, maxVal int) {
	defer wg.Done()
	splitChoices := strings.Split(rec, "-")
	splitChoice, _ := strconv.Atoi(splitChoices[len(splitChoices) - 1])
	if splitChoice !=  maxVal {
		ch <- rec
	} else {
		writech <- rec
	}
}

func DetermineRoutes(records *[]string, currentChoices []string, data[]int, maxVal int) {
	// records -> all entries added to array
	// currentChoices -> all leaps made last iteration
	// data -> array of data
	fmt.Println("Called DetermineRoutes, current choices to look at ", len(currentChoices))
	currTime := time.Now()
	nextRecords := []string{}
	nextRecordsChannel := make(chan string)
	nextRecordWaitGroup := &sync.WaitGroup{}

	nextRecordWaitGroup.Add(len(currentChoices))
	for _, choice := range currentChoices {
		go generateNextRecords(nextRecordsChannel, nextRecordWaitGroup, choice, currentChoices, data)
	}

	go func() {
		for val := range nextRecordsChannel {
			nextRecords = append(nextRecords, val)
		}
	}()

	nextRecordWaitGroup.Wait()
	close(nextRecordsChannel)

	fmt.Println("Finished getting next hops, elapsed ", time.Now().Sub(currTime))
	currTime = time.Now()

	recordsToSend := []string{}
	recordsToRecord := []string{}

	sendChannel := make(chan string)
	recordChannel := make(chan string)
	recordsWG := &sync.WaitGroup{}


	if len(nextRecords) > 0 {
		// still more work to be done
		recordsWG.Add(len(nextRecords))
		for _, rec := range nextRecords {
			go generateRecordsToSend(sendChannel, recordChannel, recordsWG, rec, maxVal)
		}
		// Do the next level of work
		go func() {
			for val := range sendChannel {
				recordsToSend = append(recordsToSend, val)
			}
		}()

		go func() {
			for val := range recordChannel {
				recordsToRecord = append(recordsToRecord, val)
			}
		}()

		recordsWG.Wait()
		close(sendChannel)
		close(recordChannel)

		fmt.Println("Finished getting what to pass next and what to save", time.Now().Sub(currTime))
		for _, val := range recordsToRecord {
			fmt.Println("appending to records", val)
			*records = append(*records, val)
		}

		if len(recordsToSend) > 0 {
			DetermineRoutes(records, nextRecords, data, maxVal)
		}
	}
}

func GetMapOfJoltDiffs(data []int) (JoltRecord, error) {
	jolts := JoltRecord{"1": 0, "2": 0, "3": 0}
	for idx, jolt := range data {
		if idx == len(data) -1 {
			break
		}
		joltDiff, err := GetNextJoltDiff(jolt, data[idx + 1])
		if err != nil {
			return jolts, errors.New("Couldn't create map because invalid jolt diff")
		}
		switch joltDiff {
		case One:
			jolts["1"] += 1
		case Two:
			jolts["2"] += 1
		case Three:
			jolts["3"] += 1
 		}
	}
	return jolts, nil
}