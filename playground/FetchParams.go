package playground

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FetchParams(examplefilepath string, antFarm *AntFarm) {
	file, err := os.Open(examplefilepath)
	if err != nil {
		fmt.Printf("No file %v found", examplefilepath)
		os.Exit(0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	tunnelGraph := make(Graph)
	var skip bool
	antFarm.AllRoomsMap = make(map[string]bool)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i, line := range lines {
		if !skip && strings.Contains(line, "##start") {
			skip = true
			antFarm.StartRoom.RoomName = strings.Split(lines[i+1], " ")[0]
			antFarm.AllRoomsMap[antFarm.StartRoom.RoomName] = true
		} else if !skip && strings.Contains(line, "##end") {
			skip = true
			antFarm.EndRoom.RoomName = strings.Split(lines[i+1], " ")[0]
			antFarm.AllRoomsMap[strings.Split(lines[i+1], " ")[0]] = false
		} else if strings.Contains(line, "-") {
			skip = false

			fromRoom := Room{
				RoomName:  strings.Split(line, "-")[0],
				IsChecked: false,
			}
			toRoom := Room{
				RoomName:  strings.Split(line, "-")[1],
				IsChecked: false,
			}
			tunnelGraph.addEdge(fromRoom, toRoom)
		} else if !skip && i != 0 && !strings.Contains(line, "-") && strings.Contains(line, " ") {
			antFarm.AllRoomsMap[strings.Split(lines[i], " ")[0]] = false
			skip = false
		} else {
			skip = false
		}
	}
	antFarm.TunnelGraph = tunnelGraph
	if _, exists := tunnelGraph[antFarm.EndRoom.RoomName]; !exists {
		fmt.Println("ERROR: invalid data format. No path between start and end room.")
		os.Exit(1)
	}
	antFarm.AntNr, err = strconv.Atoi(lines[0])
	if antFarm.AntNr < 1 {
		fmt.Println("ERROR: invalid data format, invalid number of Ants")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println("Number of ants could not be read from file: issue with format")
		os.Exit(1)
	}
}
