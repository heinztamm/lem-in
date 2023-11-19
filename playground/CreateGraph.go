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
	// var startRoomName any
	// var endRoomName any
	var lines []string
	tunnelGraph := make(Graph)
	var skip bool
	antFarm.AllRoomsMap = make(map[string]Room)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i, line := range lines {
		if !skip && strings.Contains(line, "##start") {
			skip = true
			antFarm.StartRoom.RoomName = strings.Split(lines[i+1], " ")[0]
		} else if !skip && strings.Contains(line, "##end") {
			skip = true
			antFarm.EndRoom.RoomName = strings.Split(lines[i+1], " ")[0]
			antFarm.AllRoomsMap[strings.Split(lines[i+1], " ")[0]] = Room{
				RoomName:  strings.Split(lines[i+1], " ")[0],
				IsChecked: false,
			}
		} else if strings.Contains(line, "-") {
			skip = false
			fromRoomWithName := strings.Split(line, "-")[0]
			toRoom := Room{
				RoomName:  strings.Split(line, "-")[1],
				IsChecked: false,
			}
			tunnelGraph.addDirectedEdge(fromRoomWithName, toRoom)
		} else if !skip && i != 0 && !strings.Contains(line, "-") {
			antFarm.AllRoomsMap[strings.Split(lines[i], " ")[0]] = Room{
				RoomName:  strings.Split(lines[i], " ")[0],
				IsChecked: false,
			}
			skip = false
		} else {
			skip = false
		}
	}

	antFarm.TunnelGraph = tunnelGraph
	antFarm.AntNr, err = strconv.Atoi(lines[0])
	if err != nil {
		fmt.Println("Number of ants could not be read from file: issue with format")
		os.Exit(1)
	}

	// fmt.Println(antNr)
	// fmt.Println()
	// fmt.Println(startRoomName)
	// fmt.Println(antFarm.EndRoom.RoomName)
	// tunnelGraph.PrintGraph()
}
