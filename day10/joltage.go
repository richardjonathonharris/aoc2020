package joltage

import (
	"errors"
	"fmt"
)

type JoltDiff int

const (
	One   JoltDiff = 1
	Two            = 2
	Three          = 3
)

type JoltRecord map[string]int

func GetNextJoltDiff(currJolt int, nextJolt int) (JoltDiff, error) {
	if nextJolt-currJolt < 1 || nextJolt-currJolt > 3 {
		return One, errors.New("Jolt Difference is too big")
	}
	return JoltDiff(nextJolt - currJolt), nil
}

func GetPossibleNextJoltDiffs(currJolt int, data []int) []int {
	goodJolts := []int{}
	for _, nextJolt := range data {
		if nextJolt-currJolt < 1 || nextJolt-currJolt > 3 {
			continue
		} else {
			goodJolts = append(goodJolts, nextJolt)
		}
	}
	return goodJolts
}

func MapOfAllJoltDiffs(data []int) map[int][]int {
	returnMap := map[int][]int{}
	for _, val := range data {
		possibleDiffs := GetPossibleNextJoltDiffs(val, data)
		if len(possibleDiffs) == 0 {
			continue
		} else {
			returnMap[val] = possibleDiffs
		}
	}
	return returnMap
}

func GetMapOfJoltDiffs(data []int) (JoltRecord, error) {
	jolts := JoltRecord{"1": 0, "2": 0, "3": 0}
	for idx, jolt := range data {
		if idx == len(data)-1 {
			break
		}
		joltDiff, err := GetNextJoltDiff(jolt, data[idx+1])
		if err != nil {
			return jolts, errors.New("Couldn't create map because invalid jolt diff")
		}
		switch joltDiff {
		case One:
			jolts["1"] += 1
		case Two:
			jolts["2"] += 1
		case Three:
			jolts["3"] += 1
		}
	}
	return jolts, nil
}

type Graph struct {
	Adj map[int][]int
}

func (g *Graph) AddEdge(node int, connectingNode int) {
	if len(g.Adj[node]) == 0 {
		g.Adj[node] = []int{connectingNode}
	} else {
		g.Adj[node] = append(g.Adj[node], connectingNode)
	}
}

func (g *Graph) GetAllPaths(startNode int, endNode int, paths *int) {
	visited := make(map[int]bool)
	for key, _ := range g.Adj {
		visited[key] = false
	}
	path := []int{}
	g.GetAllPathsUtil(startNode, endNode, visited, path, paths)
}

func (g *Graph) GetAllPathsUtil(currentNode int, endNode int, visited map[int]bool, path []int, paths *int) {
	visited[currentNode] = true
	path = append(path, currentNode)
	if currentNode == endNode {
		fmt.Println(*paths + 1)
		*paths += 1
	} else {
		for _, node := range g.Adj[currentNode] {
			if (visited)[node] == false {
				g.GetAllPathsUtil(node, endNode, visited, path, paths)
			}
		}
	}
	path = path[0:len(path) - 1]
	visited[currentNode] = false
}
