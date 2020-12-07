package bags

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseDirectionsWithNoContains(t *testing.T) {
	testString := "faded blue bags contain no other bags."
	expectedColor := "faded blue"
	actualColor, actualContains := ParseDirections(testString)
	if actualColor != expectedColor {
		t.Errorf(fmt.Sprintf("Expected %s received %s", expectedColor, actualColor))
	}
	if len(actualContains) > 0 {
		t.Errorf(fmt.Sprintf("Expected empty contains, got %+v", actualContains))
	}
}

func TestParseDirectionWithOneContains(t *testing.T) {
	testString := "bright white bags contain 1 shiny gold bag."
	expectedColor := "bright white"
	expectedContains := Contains{Number: "1", Color: "shiny gold"}
	actualColor, actualContains := ParseDirections(testString)
	if actualColor != expectedColor {
		t.Errorf(fmt.Sprintf("Expected %s received %s", expectedColor, actualColor))
	}
	if len(actualContains) == 0 {
		t.Errorf(fmt.Sprintf("Expected non-empty contains, got %+v", actualContains))
	}
	if actualContains[0] != expectedContains {
		t.Errorf(fmt.Sprintf("Expected %+v contains, got %+v", expectedContains, actualContains))
	}
}

func TestParseDirectionWithMultipleContains(t *testing.T) {
	testString := "dark olive bags contain 3 faded blue bags, 4 dotted black bags."
	expectedColor := "dark olive"
	expectedContains := []Contains{{Number: "3", Color: "faded blue"}, {Number: "4", Color: "dotted black"}}
	actualColor, actualContains := ParseDirections(testString)
	if actualColor != expectedColor {
		t.Errorf(fmt.Sprintf("Expected %s received %s", expectedColor, actualColor))
	}
	if len(actualContains) == 0 {
		t.Errorf(fmt.Sprintf("Expected non-empty contains, got %+v", actualContains))
	}
	if actualContains[0] != expectedContains[0] {
		t.Errorf(fmt.Sprintf("Expected %+v contains, got %+v", expectedContains[0], actualContains[0]))
	}
	if actualContains[1] != expectedContains[1] {
		t.Errorf(fmt.Sprintf("Expected %+v contains, got %+v", expectedContains[1], actualContains[1]))
	}
}

func TestAddsNewBagToBagMap(t *testing.T) {
	bagMap := BagMap{}
	testString := "dark olive bags contain 3 faded blue bags, 4 dotted black bags."
	bag := CreateNewBag(testString)
	bagMap[bag.Color] = bag
	expectedBag := Bag{Color: "dark olive", Contains: []Contains{{Number: "3", Color: "faded blue"}, {Number: "4", Color: "dotted black"}}}
	if !reflect.DeepEqual(bagMap["dark olive"], expectedBag) {
		t.Errorf(fmt.Sprintf("Expected populated map, got %+v", bagMap))
	}
}

func TestCanAddContainingColors(t *testing.T) {
	bagMap := BagMap{}
	testString1 := "dark olive bags contain 3 faded blue bags, 4 dotted black bags."
	testString2 := "showy brown bags contain 1 dark olive bag."
	bag := CreateNewBag(testString1)
	bag2 := CreateNewBag(testString2)
	bagMap[bag.Color] = bag
	bagMap[bag2.Color] = bag2
	resp := AddAllContainingBags("dark olive", bagMap)
	if resp[0] != "showy brown" {
		t.Errorf("Did not get containing bag")
	}
}

func TestCanListAllContainingBags(t *testing.T) {
	bagMap := BagMap{
		"blue":  Bag{Color: "blue", ContainedBy: []string{"red", "gold"}, Contains: []Contains{}},
		"red":   Bag{Color: "red", ContainedBy: []string{}, Contains: []Contains{}},
		"gold":  Bag{Color: "gold", ContainedBy: []string{"azure"}, Contains: []Contains{}},
		"azure": Bag{Color: "azure", ContainedBy: []string{}, Contains: []Contains{}},
	}
	containingBags := map[string]int{}
	ListAllContainingBags("blue", bagMap, containingBags)
	if len(containingBags) != 3 {
		t.Errorf("Expected 3 containing bags, got %+v", containingBags)
	}
}
