package playground

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CreateGraph(examplefilepath string) {
	file, err := os.Open(examplefilepath)
	if err != nil {
		fmt.Printf("No file %v found", examplefilepath)
		os.Exit(0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var startRoomName any
	var endRoomName any
	var lines []string
	tunnelGraph := make(Graph)
	var skip bool

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i, line := range lines {
		if !skip && strings.Contains(line, "##start") {
			startRoomName = strings.Split(lines[i+1], " ")[0]
			skip = true
		} else if !skip && strings.Contains(line, "##end") {
			endRoomName = strings.Split(lines[i+1], " ")[0]
			skip = true
		} else if strings.Contains(line, "-") {
			tunnelGraph.addDirectedEdge(strings.Split(line, "-")[0], strings.Split(line, "-")[1])
			skip = false
		} else {
			skip = false
		}
	}

	antNr, err := strconv.Atoi(lines[0])
	if err != nil {
		fmt.Println("Number of ants could not be read from file: issue with format")
		os.Exit(1)
	}
	fmt.Println(antNr)
	fmt.Println()
	fmt.Println(startRoomName)
	fmt.Println(endRoomName)
	tunnelGraph.PrintGraph()
}
