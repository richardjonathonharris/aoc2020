package bits

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func IntToBit(value int) string {
	return fmt.Sprintf("%036b", value)
}

func BitToInt(value string) int {
	conversion, err := strconv.ParseInt(value, 2, 64)
	if err != nil {
		fmt.Println("Could not convert value: ", value)
		panic(err)
	}
	return int(conversion)
}

func ApplyMaskToBit(bit string, mask string) string {
	bits := strings.Split(bit, "")
	for idx, val := range mask {
		if string(val) != "X" {
			bits[idx] = string(val)
		}
	}
	return strings.Join(bits, "")
}

func ApplyMaskToBitFloater(bit string, mask string) string {
	bits := strings.Split(bit, "")
	for idx, val := range mask {
		if string(val) == "0" {
			continue
		} else {
			bits[idx] = string(val)
		}
	}
	return strings.Join(bits, "")
}

func GenerateMaskOption(vals [][]bool) [][]bool {
	var options [][]bool
	for _, val := range vals {
		for _, opt := range []bool{true, false} {
			expandedOption := []bool{}
			for _, v := range val {
				expandedOption = append(expandedOption, v)
			}
			expandedOption = append(expandedOption, opt)
			options = append(options, expandedOption)
		}
	}
	return options
}

func GenerateMaskOptions(optionsAtLevel map[int][][]bool, totalNumLevels int) {
	if len(optionsAtLevel) < totalNumLevels {
		maxKey := 0
		for k := range optionsAtLevel {
			if k > maxKey {
				maxKey = k
			}
		}
		optionsAtLevel[maxKey+1] = GenerateMaskOption(optionsAtLevel[maxKey])
		GenerateMaskOptions(optionsAtLevel, totalNumLevels)
	}
}

func GenerateMemoryAddressDecoder(options [][]bool, val string) []int {
	// find indices of "X"
	indices := []int{}
	for i, charVar := range val {
		if string(charVar) == "X" {
			indices = append(indices, i)
		}
	}
	bitAddresses := []string{}
	for _, opt := range options {
		newString := strings.Split(val, "")
		for idx, boolVal := range opt {
			if boolVal == true {
				newString[indices[idx]] = "1"
			} else {
				newString[indices[idx]] = "0"
			}
		}
		bitAddresses = append(bitAddresses, strings.Join(newString, ""))
	}
	intAddresses := []int{}
	for _, bitAddress := range bitAddresses {
		intAddresses = append(intAddresses, BitToInt(bitAddress))
	}
	return intAddresses
}

func ParseInstructions(instructions []string) map[string]int {
	memAddress := map[string]int{}
	maskCommand := regexp.MustCompile(`mask = (?P<mask>[0,1,X]*)`)
	memCommand := regexp.MustCompile(`mem\[(?P<address>[0-9]*)\] = (?P<value>[0-9]*)`)
	mask := ""
	for _, instr := range instructions {
		if maskCommand.MatchString(instr) {
			mask = maskCommand.FindStringSubmatch(instr)[1]
		} else if memCommand.MatchString(instr) {
			address := memCommand.FindStringSubmatch(instr)[1]
			value, err := strconv.Atoi(memCommand.FindStringSubmatch(instr)[2])
			if err != nil {
				fmt.Println("COULD NOT CONVERT")
			}
			valueBit := BitToInt(ApplyMaskToBit(IntToBit(value), mask))
			memAddress[address] = valueBit
		}
	}
	return memAddress
}

func ParseInstructionsMemoryAddresses(instructions []string) map[int]int {
	memAddress := map[int]int{}
	maskCommand := regexp.MustCompile(`mask = (?P<mask>[0,1,X]*)`)
	memCommand := regexp.MustCompile(`mem\[(?P<address>[0-9]*)\] = (?P<value>[0-9]*)`)
	mask := ""
	fmt.Println("Generating 25 levels of options, hopefully that's enough")
	optionsAtLevel := map[int][][]bool{}
	optionsAtLevel[1] = [][]bool{{true}, {false}}
	GenerateMaskOptions(optionsAtLevel, 25)
	for _, instr := range instructions {
		if maskCommand.MatchString(instr) {
			mask = maskCommand.FindStringSubmatch(instr)[1]
		} else if memCommand.MatchString(instr) {
			// get memory address in bit
			memoryAddressInt := memCommand.FindStringSubmatch(instr)[1]
			memAddInt, err := strconv.Atoi(memoryAddressInt)
			if err != nil {
				panic(err)
			}
			memoryAddressBit := IntToBit(memAddInt)
			// get value to put at those addresses
			value, err := strconv.Atoi(memCommand.FindStringSubmatch(instr)[2])
			if err != nil {
				fmt.Println("COULD NOT CONVERT")
			}

			// apply the mask to the bit
			maskedBit := ApplyMaskToBitFloater(memoryAddressBit, mask)

			// generate the new addresses to update
			countXs := strings.Count(maskedBit, "X")
			if len(optionsAtLevel) < countXs {
				panic("We don't have enough options!")
			}
			memoryAddresses := GenerateMemoryAddressDecoder(optionsAtLevel[countXs], maskedBit)

			for _, memoryAddress := range memoryAddresses {
				memAddress[memoryAddress] = value
			}
		}
	}
	return memAddress
}
