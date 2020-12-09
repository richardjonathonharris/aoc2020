package console

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	Op        Operation
	Value     int
	NumCalled int
}

type Operation string

const (
	Accumulate Operation = "acc"
	Jump                 = "jmp"
	NoOp                 = "nop"
)

func GetNextIndexAndAccumulator(currAcc int, instr Instruction) (newIndex int, newAcc int) {
	switch instr.Op {
	case NoOp:
		return 1, currAcc
	case Accumulate:
		return 1, currAcc + instr.Value
	case Jump:
		return instr.Value, currAcc
	default:
		// some kind of invalid operation, move on to next index
		return 1, currAcc
	}
}

func BuildCodebook(data []string) map[int]*Instruction {
	codebook := map[int]*Instruction{}
	for idx, inst := range data {
		splitVals := strings.Split(inst, " ")
		value, _ := strconv.Atoi(splitVals[1])
		codebook[idx] = &Instruction{Op: Operation(splitVals[0]), Value: value}
	}
	return codebook
}

func FindRepeatedInstruction(codebook map[int]*Instruction) (int, error) {
	notRepeated := false
	acc := 0
	idx := 0
	codebook[0].NumCalled += 1
	for notRepeated != true {
		newIdx, newAcc := GetNextIndexAndAccumulator(acc, *codebook[idx])
		if idx+newIdx == len(codebook) {
			// this is the last instruction, so break instruction loop via error
			fmt.Println(idx, newIdx, acc, newAcc)
			return newAcc, errors.New("Attention: would execute next instruction after end of instructions!")
		} else if codebook[idx+newIdx].NumCalled == 1 {
			notRepeated = true
		} else {
			codebook[idx+newIdx].NumCalled += 1
		}
		idx += newIdx
		acc = newAcc
	}
	return acc, nil
}
