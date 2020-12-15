package game

type Memory struct {
	LastSpoken int
}

func NumberGame(inputs []int, iters int) int {
	// key will be the number, value will be the last turn it was spoken
	numberMap := map[int]Memory{}
	// prepopulate number map with the seed inputs
	for idx, val := range inputs[0 : len(inputs)-1] {
		numberMap[val] = Memory{LastSpoken: idx + 1}
	}
	lastValueSaid := inputs[len(inputs)-1]
	for i := len(inputs) + 1; i < iters+1; i++ {
		_, ok := numberMap[lastValueSaid]
		if ok == false {
			// adding in the new record to state that it's been said
			numberMap[lastValueSaid] = Memory{LastSpoken: i - 1}
			lastValueSaid = 0
		} else {
			// last time it spoken was not its first time, so find difference and update
			newLastValue := (i - 1) - numberMap[lastValueSaid].LastSpoken
			numberMap[lastValueSaid] = Memory{LastSpoken: i - 1}
			lastValueSaid = newLastValue
		}
	}
	return lastValueSaid
}
