package main

import (
	"fmt"
	"os"

	"01.kood.tech/git/kartamm/lem-in/playground"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: wrong command format. Expected format: 'go run . <path-to-textfile-name>'")
		os.Exit(1)
	}
	examplefilepath := os.Args[1]
	antFarm := &playground.AntFarm{}
	playground.FetchParams(examplefilepath, antFarm)

	// setsOfPaths := playground.FindSetsOfValidPaths(antFarm)
	playground.FindSetsOfValidPaths(antFarm)
	fmt.Println(antFarm.ValidPaths3D)
	// chosenPaths := playground.ReturnLongest(setsOfPaths)
	// playground.Movement(antFarm.AntNr, antFarm.StartRoom.RoomName, antFarm.EndRoom.RoomName, chosenPaths)
}
