package customs

import (
	"fmt"
	"testing"
)

func TestCountsAllUniqueLetters(t *testing.T) {
	testString := `
		a
		abc
		kmno
	`
	resp := CountUniqueLetters(testString)
	if resp != 7 {
		t.Errorf("CountUniqueLetters should return 7 unique questions.")
	}
}

func TestCountsAllTotalQuestionsAnswered(t *testing.T) {
	testString := `a
		abc
		kmno`
	resp := CountLettersAllAnswer(testString)
	if resp != 1 {
		t.Errorf(fmt.Sprintf("Received %d expected 1", resp))
	}
}
