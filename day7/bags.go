package bags

import (
	"fmt"
	"strconv"
	"strings"
)

type Contains struct {
	Number string
	Color  string
}

type Bag struct {
	Color       string
	Contains    []Contains
	ContainedBy []string
}

type BagMap map[string]Bag

// ParseDirections returns the color of the bag and array of what it contains
func ParseDirections(direction string) (string, []Contains) {
	splitDirection := strings.Split(direction, " contain ")
	color := strings.ReplaceAll(splitDirection[0], " bags", "")
	bagContains := []Contains{}
	for _, info := range strings.Split(strings.ReplaceAll(splitDirection[1], ".", ""), ", ") {
		if info == "no other bags" {
			break
		} else {
			chunks := strings.Split(strings.TrimSpace(info), " ")
			bagInfo := Contains{
				Number: chunks[0],
				Color:  strings.Join(chunks[1:len(chunks)-1], " "),
			}
			bagContains = append(bagContains, bagInfo)
		}
	}
	return color, bagContains
}

func CreateNewBag(direction string) Bag {
	color, contains := ParseDirections(direction)
	return Bag{Color: color, Contains: contains}
}

func AddAllContainingBags(color string, bm BagMap) []string {
	var containingColors []string
	for bagColor, bag := range bm {
		for _, containedBag := range bag.Contains {
			if containedBag.Color == color {
				containingColors = append(containingColors, bagColor)
			}
		}
	}
	return containingColors
}

func ListAllContainingBags(color string, bm BagMap, containingBags map[string]int) {
	for _, bagColor := range bm[color].ContainedBy {
		containingBags[bagColor]++
		ListAllContainingBags(bagColor, bm, containingBags)
	}
}

func CountAllBagsThatBagContains(color string, bm BagMap, containingBags map[string]int) {
	for _, bag := range bm[color].Contains {
		num, err := strconv.Atoi(bag.Number)
		if err != nil {
			fmt.Println("ERROR COULD NOT CONVERT")
		}
		for i := 0; i < num; i++ {
			containingBags[bag.Color] += 1
			CountAllBagsThatBagContains(bag.Color, bm, containingBags)
		}
	}
}
