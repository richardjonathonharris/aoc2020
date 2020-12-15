package bits

import (
	"testing"
)

func TestIntToBit(t *testing.T) {
	resp := IntToBit(73)
	if resp != "000000000000000000000000000001001001" {
		t.Errorf("Should convert correctly, received %s", resp)
	}
}

func TestBitToInt(t *testing.T) {
	resp := BitToInt("000000000000000000000000000001001001")
	if resp != 73 {
		t.Errorf("Should convert correctly, received %d", resp)
	}
}

func TestApplyMaskToBit(t *testing.T) {
	bit := "000000000000000000000000000000001011"
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	resp := ApplyMaskToBit(bit, mask)
	if resp != "000000000000000000000000000001001001" {
		t.Errorf("Should convert correctly, received %s", resp)
	}
}

func TestApplyMaskToBitFloater(t *testing.T) {
	bit := "000000000000000000000000000000101010"
	mask := "000000000000000000000000000000X1001X"
	resp := ApplyMaskToBitFloater(bit, mask)
	if resp != "000000000000000000000000000000X1101X" {
		t.Errorf("Should convert correctly, received %s", resp)

	}
}

func TestGenerateMaskOption(t *testing.T) {
	currentRun := [][]bool{
		{true, true, true},
		{true, true, false},
		{true, false, false},
		{false, true, true},
		{false, true, false},
		{false, false, true},
		{true, false, true},
		{false, false, false},
	}
	options := GenerateMaskOption(currentRun)
	if len(options) != 16 {
		t.Errorf("Did not generate correct number of options")
	}
}

func TestGenerateMaskOptions(t *testing.T) {
	optionsAtLevel := map[int][][]bool{}
	optionsAtLevel[1] = [][]bool{{true}, {false}}
	GenerateMaskOptions(optionsAtLevel, 5)
	if len(optionsAtLevel) != 5 {
		t.Errorf("Did not create all specified options")
	}
}

func TestApplyMemoryDecoderMaskToBit(t *testing.T) {
	maskedValue := "00000000000000000000000000000001X0XX"
	optionsAtLevel := map[int][][]bool{}
	optionsAtLevel[1] = [][]bool{{true}, {false}}
	GenerateMaskOptions(optionsAtLevel, 3)
	resp := GenerateMemoryAddressDecoder(optionsAtLevel[3], maskedValue)
	if len(resp) != 8 {
		t.Errorf("Should convert to array of 8 memory addresses, %+v", resp)
	}
}

func TestInstructionParse(t *testing.T) {
	instructions := []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}

	memAddress := ParseInstructions(instructions)

	if memAddress["8"] != 64 {
		t.Errorf("MemAddress not created correctly, received %d", memAddress["8"])
	}
	if memAddress["7"] != 101 {
		t.Errorf("MemAddress not created correctly, received %d", memAddress["7"])
	}
}

func TestInstructionParseMemoryAddress(t *testing.T) {
	instructions := []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}
	memAddress := ParseInstructionsMemoryAddresses(instructions)

	sum := 0
	for _, val := range memAddress {
		sum += val
	}
	if sum != 208 {
		t.Errorf("Did not sum up correctly, memAddress is %+v", memAddress)
	}
}
