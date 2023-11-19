package playground

import (
	"fmt"
	"os"
)

// Determine a valid path through the farm, if one exists, and return it as a slice of room names
// in order of visiting
func FindValidPaths(antFarm *AntFarm) [][]Room {
	// fmt.Println(ScanPath(antFarm.TunnelGraph, antFarm.RoomNames, antFarm.StartRoom.RoomName, antFarm.EndRoom.RoomName))
	_, exists := antFarm.TunnelGraph[antFarm.StartRoom.RoomName]
	if !exists {
		return nil
	}

	for _, cncRoom := range antFarm.TunnelGraph[antFarm.StartRoom.RoomName] {
		antFarm.AllRoomsMap[cncRoom.RoomName] = Room{
			RoomName:  cncRoom.RoomName,
			IsChecked: false,
		}
		antFarm.PossiblePaths = append(antFarm.PossiblePaths, []Room{antFarm.AllRoomsMap[cncRoom.RoomName]})
	}

	var chosenPath []Room
	var validPaths [][]Room
	count := 0
	if len(antFarm.PossiblePaths) > 0 {
		ScanForPath(antFarm.TunnelGraph, antFarm.AllRoomsMap, antFarm.EndRoom.RoomName, antFarm.PossiblePaths, validPaths)
		ScanForPath(antFarm.TunnelGraph, antFarm.AllRoomsMap, antFarm.EndRoom.RoomName, antFarm.PossiblePaths, validPaths)
		ScanForPath(antFarm.TunnelGraph, antFarm.AllRoomsMap, antFarm.EndRoom.RoomName, antFarm.PossiblePaths, validPaths)

		count++

		// proceed until there are no rooms connected to the starting room still unchecked (not included in a determinedly valid path)
		for HasPossiblePathLeft(antFarm.TunnelGraph, antFarm.AllRoomsMap, antFarm.StartRoom.RoomName) {
			for !ValidPathFound(antFarm.PossiblePaths, antFarm.EndRoom.RoomName) {

				ScanForPath(antFarm.TunnelGraph, antFarm.AllRoomsMap, antFarm.EndRoom.RoomName, antFarm.PossiblePaths, validPaths)

				count++
				if count > 25 {
					fmt.Println(antFarm.AllRoomsMap)
					fmt.Println("Failurre")
					os.Exit(1)
				}
			}
			validPaths = append(validPaths, chosenPath)
		}
		fmt.Println("Nr of ScanForPath calls: ", count)
		fmt.Println(antFarm.PossiblePaths)
	} else {
		fmt.Println("No possible paths")
		return nil
	}
	return validPaths
}

// Go through possible paths, adding them to a slice of possible path slices, that is finally returned in FindValidPaths()
func ScanForPath(tunnelGraph Graph, allRoomsMap map[string]Room, endRoomName string, possiblePaths, validPaths [][]Room) {
	// chosenPath := []Room{}
	for i, path := range possiblePaths {
		if len(path) == 1 {
			path[0].IsChecked = true
			allRoomsMap[path[0].RoomName] = path[0]
		}
		currentMoveRoom := path[len(path)-1]
		_, exists := tunnelGraph[currentMoveRoom.RoomName]
		if exists {
			for j, connectedRoom := range tunnelGraph[currentMoveRoom.RoomName] {
				if !(allRoomsMap[connectedRoom.RoomName].IsChecked) && j > 0 {
					if connectedRoom.RoomName != endRoomName {
						allRoomsMap[connectedRoom.RoomName] = Room{
							RoomName:  connectedRoom.RoomName,
							IsChecked: true,
						}
					}
					// If the room has more than one unchecked connected room, an alternative possible path is appended
					altPath := append(path, allRoomsMap[connectedRoom.RoomName])
					possiblePaths = append(possiblePaths, altPath)
				} else if !(allRoomsMap[connectedRoom.RoomName].IsChecked) {
					fmt.Println("Moving...")
					// append room's name to corresponding subslice of possiblePaths
					if connectedRoom.RoomName != endRoomName {
						allRoomsMap[connectedRoom.RoomName] = Room{
							RoomName:  connectedRoom.RoomName,
							IsChecked: true,
						}
					} else {
						possiblePaths[i] = append(possiblePaths[i], allRoomsMap[connectedRoom.RoomName])
						validPaths = append(validPaths, possiblePaths[i])
						leftoverPaths := append(possiblePaths[:i], possiblePaths[i+1:]...)
						UnCheckLeftOverRooms(allRoomsMap, leftoverPaths)
						return
					}
					possiblePaths[i] = append(possiblePaths[i], allRoomsMap[connectedRoom.RoomName])
					// fmt.Println(possiblePaths[i])
					// fmt.Println(allRoomsMap[connectedRoom.RoomName])
					// chosenPath = possiblePaths[i]
				}
			}
		}
	}

}

func HasPossiblePathLeft(tunnelGraph Graph, allRoomsMap map[string]Room, startRoomName string) bool {
	for _, cncRoom := range tunnelGraph[startRoomName] {
		if allRoomsMap[cncRoom.RoomName].IsChecked == false {
			return true
		}
	}
	return false
}

// lower all flags of rooms not part of a valid path
func UnCheckLeftOverRooms(allRoomsMap map[string]Room, leftOverPaths [][]Room) {
	for _, path := range leftOverPaths {
		for j := range path {
			path[j].IsChecked = false
			allRoomsMap[path[j].RoomName] = path[j]
		}
	}
}

func ValidPathFound(possiblePaths [][]Room, endRoomName string) bool {
	for _, path := range possiblePaths {
		if path[len(path)-1].RoomName == endRoomName {
			return true
		}
	}
	return false
}
