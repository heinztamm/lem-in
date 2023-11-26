package main

import (
	"os"

	"01.kood.tech/git/kartamm/lem-in/playground"
)

func main() {

	examplefilepath := os.Args[1]
	antFarm := &playground.AntFarm{}
	playground.FetchParams(examplefilepath, antFarm)

	setsOfPaths := playground.FindSetsOfValidPaths(antFarm)
	chosenPaths := playground.ReturnLongest(setsOfPaths)
	playground.Movement(antFarm.AntNr, antFarm.StartRoom.RoomName, antFarm.EndRoom.RoomName, chosenPaths)
}
