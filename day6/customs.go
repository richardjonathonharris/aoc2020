package customs

import (
	"regexp"
	"strings"
)

func CountUniqueLetters(questions string) int {
	spaces := regexp.MustCompile(`\s+`)
	joinedQuestions := spaces.ReplaceAllString(strings.ReplaceAll(questions, "\n", ""), "")
	questHash := make(map[string]int)
	for _, q := range joinedQuestions {
		questHash[string(q)]++
	}
	return len(questHash)
}

func CountLettersAllAnswer(questions string) int {
	numRespondants := len(strings.Split(questions, "\n"))
	spaces := regexp.MustCompile(`\s+`)
	joinedQuestions := spaces.ReplaceAllString(strings.ReplaceAll(questions, "\n", ""), "")
	questHash := make(map[string]int)
	for _, q := range joinedQuestions {
		questHash[string(q)]++
	}
	questionsEveryoneAnswered := 0
	for _, value := range questHash {
		if value == numRespondants {
			questionsEveryoneAnswered += 1
		}
	}
	return questionsEveryoneAnswered
}
