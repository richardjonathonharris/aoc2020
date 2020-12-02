package password

import (
	"errors"
	"strconv"
	"strings"
)

type Instruction struct {
	Letter  string
	Minimum int
	Maximum int
}

func ParseInstructions(instr string) (Instruction, error) {
	initialSplit := strings.Split(instr, " ")
	sliceSplit := strings.Split(initialSplit[0], "-")
	minValue, err := strconv.Atoi(sliceSplit[0])
	if err != nil {
		return Instruction{}, errors.New("Failed to convert string to int")
	}
	maxValue, err := strconv.Atoi(sliceSplit[1])
	if err != nil {
		return Instruction{}, errors.New("Failed to convert string to int")
	}
	return Instruction{
		Letter:  initialSplit[1],
		Minimum: minValue,
		Maximum: maxValue,
	}, nil
}

func CheckPassword(instr Instruction, password string) (bool, error) {
	charCount := 0
	for _, c := range password {
		if instr.Letter == string(c) {
			charCount += 1
		}
	}
	greaterThanMinimum := instr.Minimum <= charCount
	lessThanMaximum := charCount <= instr.Maximum
	return greaterThanMinimum && lessThanMaximum, nil
}

func CheckPasswordExactOne(instr Instruction, password string) (bool, error) {
	if instr.Maximum-1 > len(password) {
		return false, errors.New("Password is shorter than maximum index")
	}
	minimumCharacterCorrect := string(password[instr.Minimum-1]) == instr.Letter
	maximumCharacterCorrect := string(password[instr.Maximum-1]) == instr.Letter
	if (minimumCharacterCorrect && maximumCharacterCorrect) || (!minimumCharacterCorrect && !maximumCharacterCorrect) {
		return false, nil
	} else {
		return true, nil
	}
}
