package password

import (
	"testing"
)

func TestCanParsePasswordInstructions(t *testing.T) {
	testString := "2-9 c"
	res, _ := ParseInstructions(testString)
	correctInstruction := Instruction{
		Letter:  "c",
		Minimum: 2,
		Maximum: 9,
	}
	if res != correctInstruction {
		t.Errorf("ParseInstructions should return back Instruction type filled in.")
	}
}

func TestCanCheckPassword(t *testing.T) {
	testInstruction := Instruction{
		Letter:  "c",
		Minimum: 2,
		Maximum: 9,
	}
	correctResponse, _ := CheckPassword(testInstruction, "ccccccccc")
	if correctResponse != true {
		t.Errorf("CheckPassword should return true when password matches instructions")
	}
	incorrectResponse, _ := CheckPassword(testInstruction, "aaaaaaaaa")
	if incorrectResponse != false {
		t.Errorf("CheckPassword should return false when password does not match instructions")
	}
}

func TestCanCheckPasswordExactly(t *testing.T) {
	testInstruction := Instruction{
		Letter:  "c",
		Minimum: 2,
		Maximum: 9,
	}
	noCorrectMatches, _ := CheckPasswordExactOne(testInstruction, "aaaaaaaaaaaaaaaaa")
	if noCorrectMatches != false {
		t.Errorf("CheckPasswordExactOne should return false when neither instruction matches")
	}
	bothCorrectMatches, _ := CheckPasswordExactOne(testInstruction, "cccccccccccc")
	if bothCorrectMatches != false {
		t.Errorf("CheckPasswordExactOne should return false when both instructions match")
	}
	oneCorrectMatch, _ := CheckPasswordExactOne(testInstruction, "acaaaaaaaaaaaaaa")
	if oneCorrectMatch != true {
		t.Errorf("CheckPasswordExactOne should return true when only one instruction matches")
	}
}
