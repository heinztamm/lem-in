package playground

import (
	"sort"
)

func ReOrderPaths(ChosenPaths [][]string) [][]string {
	sort.Slice(ChosenPaths, func(i, j int) bool {
		return len(ChosenPaths[i]) < len(ChosenPaths[j])
	})
	return ChosenPaths
}

func ReturnLongest(setsOfPaths [][][]string) [][]string {
	indexOfLongest := 0
	maxLength := len(setsOfPaths[0])
	for i, set := range setsOfPaths {
		if len(set) > maxLength {
			maxLength = len(set)
			indexOfLongest = i
		}
	}
	return setsOfPaths[indexOfLongest]
}
