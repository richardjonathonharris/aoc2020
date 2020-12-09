package console

import (
	"testing"
)

func TestNoOpAccumulator(t *testing.T) {
	instr := Instruction{Op: NoOp, Value: 0}
	index, accumulator := GetNextIndexAndAccumulator(0, instr)
	if index != 1 {
		t.Errorf("Should increase index by one")
	}
	if accumulator != 0 {
		t.Errorf("Should not increase accumulator")
	}
}

func TestAccAccumulator(t *testing.T) {
	instr := Instruction{Op: Accumulate, Value: 5}
	index, accumulator := GetNextIndexAndAccumulator(0, instr)
	if index != 1 {
		t.Errorf("Acc should increase index by one")
	}
	if accumulator != 5 {
		t.Errorf("Should increase accumulator by value of instruction")
	}
}

func TestJumpAccumulator(t *testing.T) {
	instr := Instruction{Op: Jump, Value: 5}
	index, accumulator := GetNextIndexAndAccumulator(0, instr)
	if index != instr.Value {
		t.Errorf("Acc should increase index by value of instruction")
	}
	if accumulator != 0 {
		t.Errorf("Should not increase accumulator")
	}
}

func TestBuildCodebook(t *testing.T) {
	data := []string{"nop +0", "acc +1", "jmp -4"}
	codebook := BuildCodebook(data)
	if codebook[0].Op != NoOp || codebook[0].Value != 0 {
		t.Errorf("Did not convert first instruction correctly")
	}
	if codebook[1].Op != Accumulate || codebook[1].Value != 1 {
		t.Errorf("Did not convert second instruction correctly")
	}
	if codebook[2].Op != Jump || codebook[2].Value != -4 {
		t.Errorf("Did not convert third instruction correctly")
	}
}

func TestFindRepeatedInstruction(t *testing.T) {
	data := []string{"acc +1", "acc +1", "acc +1", "jmp -1"}
	codebook := BuildCodebook(data)
	resp, _ := FindRepeatedInstruction(codebook)
	if resp != 3 {
		t.Errorf("Could not find repeated instruction")
	}
}

func TestFindsEndOfInstructions(t *testing.T) {
	data := []string{"acc +1", "acc +1"}
	codebook := BuildCodebook(data)
	resp, err := FindRepeatedInstruction(codebook)
	if resp != 2 && err != nil {
		t.Errorf("Did not alert to end of instructions")
	}
}
