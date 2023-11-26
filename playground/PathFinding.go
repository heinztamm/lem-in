package playground

import (
	"fmt"
	"os"
	"reflect"
)

func SelectSliceOfValidPaths(validPaths3D [][][]string) [][]string {

	return [][]string{}
}

func FindSetsOfValidPaths(antFarm *AntFarm) [][][]string {
	_, exists := antFarm.TunnelGraph[antFarm.StartRoom.RoomName]
	if !exists {
		return nil
	}
	antFarm.AllRoomsMap[antFarm.StartRoom.RoomName] = true
	masterIndex := 0
	pathRepeat := false
	antFarm.ValidPaths3D = make([][][]string, 0)
	for _, cncRoomName := range antFarm.TunnelGraph[antFarm.StartRoom.RoomName] {
		antFarm.AllRoomsMap[cncRoomName] = true
		antFarm.PossiblePaths = append(antFarm.PossiblePaths, []string{cncRoomName})
	}
	if len(antFarm.PossiblePaths) > 0 {
		for masterIndex < 2 {
			for !AllEndsAreDead(&antFarm.PossiblePaths, antFarm.AllRoomsMap, antFarm.TunnelGraph, antFarm.EndRoom.RoomName) {
				ScanForPath(antFarm.TunnelGraph, antFarm.AllRoomsMap, antFarm.EndRoom.RoomName, &antFarm.PossiblePaths, &antFarm.ValidPaths3D, masterIndex)
				if len(antFarm.ValidPaths3D) > 0 {
					if len(antFarm.ValidPaths3D[masterIndex]) > 1 {
						lastElIndex := len(antFarm.ValidPaths3D[masterIndex]) - 1
						if sliceExists(antFarm.ValidPaths3D[masterIndex][:lastElIndex], antFarm.ValidPaths3D[masterIndex][lastElIndex]) {
							antFarm.ValidPaths3D[masterIndex] = antFarm.ValidPaths3D[masterIndex][:lastElIndex]
							pathRepeat = true
							break
						}
					}
				}
			}
			if pathRepeat {
				break
			}
			if len(antFarm.PossiblePaths) == len(antFarm.ValidPaths3D[masterIndex]) {
				break
			}
			RetraceValidPaths(antFarm.AllRoomsMap, &antFarm.ValidPaths3D, masterIndex)
			// fmt.Println(AllEndsAreDead(&antFarm.PossiblePaths, antFarm.AllRoomsMap, antFarm.TunnelGraph, antFarm.EndRoom.RoomName))
			masterIndex++
		}
		/*
			repeat loop until the last and second to last validPaths slices (of the 2nd dimension in ValidPaths3D)
			are equal with imported package reflect's DeepEqual() function.
			This would mean that no new combination of valid paths was found
			and the search can be concluded
		*/
		// mistake: comparing an element to itself
		for !reflect.DeepEqual(antFarm.ValidPaths3D[len(antFarm.ValidPaths3D)-1], antFarm.ValidPaths3D[len(antFarm.ValidPaths3D)-1]) {
			for !AllEndsAreDead(&antFarm.PossiblePaths, antFarm.AllRoomsMap, antFarm.TunnelGraph, antFarm.EndRoom.RoomName) {
				ScanForPath(antFarm.TunnelGraph, antFarm.AllRoomsMap, antFarm.EndRoom.RoomName, &antFarm.PossiblePaths, &antFarm.ValidPaths3D, masterIndex)
			}
			RetraceValidPaths(antFarm.AllRoomsMap, &antFarm.ValidPaths3D, masterIndex)
			masterIndex++
		}

	} else {
		fmt.Println("No possible paths")
		return nil
	}
	return antFarm.ValidPaths3D
}

// Go through possible paths, adding them to a slice of possible path slices, that is finally returned in FindValidPaths()
func ScanForPath(tunnelGraph Graph, allRoomsMap map[string]bool, endRoomName string, possiblePaths *[][]string, validPaths3D *[][][]string, masterIndex int) {
	Paths := *possiblePaths
	valPaths3DCopy := *validPaths3D
	for i, path := range Paths {
		if len(path) == 1 {
			allRoomsMap[path[0]] = true
		}
		currentMoveRoomName := path[len(path)-1]
		if currentMoveRoomName == endRoomName {
			// !! Something else to do here? !!
			// !!  Potentially problematic spot. !!
			valPaths3DCopy = append(valPaths3DCopy, [][]string{})
			valPaths3DCopy[masterIndex] = append(valPaths3DCopy[masterIndex], path)
			*validPaths3D = valPaths3DCopy
			continue
		}
		_, exists := tunnelGraph[currentMoveRoomName]
		nrOfConnections := 0
		connectedRoomNames := tunnelGraph[currentMoveRoomName]
		if exists {
			nrOfConnections = len(connectedRoomNames)
		} else if !SomewhereToGo(connectedRoomNames, allRoomsMap) {
			fmt.Println("DEAD END?!")
			return
		} else {
			fmt.Println("inexistent room, huhhh?!!")
			os.Exit(1)
		}
		if nrOfConnections == 2 {
			unvisited := ReturnUnvisited(connectedRoomNames, allRoomsMap)
			for _, connectedRoomName := range unvisited {
				// append room's name to corresponding subslice of possiblePaths
				Paths[i] = append(Paths[i], connectedRoomName)
				*possiblePaths = Paths
				if connectedRoomName != endRoomName {
					allRoomsMap[connectedRoomName] = true
				} else {
					if len(valPaths3DCopy) == masterIndex {
						valPaths3DCopy = append(valPaths3DCopy, [][]string{})
					}
					valPaths3DCopy[masterIndex] = append(valPaths3DCopy[masterIndex], Paths[i])
					*validPaths3D = valPaths3DCopy
				}
				break
			}
		} else if nrOfConnections > 2 {
			// Alternative paths need to be added, all paths will be selected between later in the parent function FindValidPaths()
			unvisited := ReturnUnvisited(connectedRoomNames, allRoomsMap)
			for _, connectedRoomName := range unvisited {
				Paths[i] = append(Paths[i], connectedRoomName)
				*possiblePaths = Paths
				if connectedRoomName != endRoomName {
					allRoomsMap[connectedRoomName] = true
				} else {
					if len(valPaths3DCopy) == masterIndex {
						valPaths3DCopy = append(valPaths3DCopy, [][]string{})
					}
					valPaths3DCopy[masterIndex] = append(valPaths3DCopy[masterIndex], Paths[i])
					*validPaths3D = valPaths3DCopy
				}
				break
			}
		}
	}
	// fmt.Println(*possiblePaths)
}

func HasPossiblePathLeft(tunnelGraph Graph, allRoomsMap map[string]Room, startRoom Room) bool {
	for _, cncRoomName := range tunnelGraph[startRoom.RoomName] {
		if allRoomsMap[cncRoomName].IsChecked == false {
			return true
		}
	}
	return false
}

// lower all flags of rooms part of a valid path
func RetraceValidPaths(allRoomsMap map[string]bool, validPaths3D *[][][]string, masterIndex int) {
	valPaths3DCopy := *validPaths3D
	for _, path := range valPaths3DCopy[masterIndex] {
		// Populate a validRooms map for quick lookups
		validRooms := make(map[string]struct{})
		for _, vRoomName := range path {
			validRooms[vRoomName] = struct{}{}
		}
		// uncheck all rooms part of a valid path
		for roomName := range allRoomsMap {
			if _, isValid := validRooms[roomName]; isValid {
				allRoomsMap[roomName] = false
			} else {
				allRoomsMap[roomName] = true
			}
		}
	}
}

func ValidPathFound(validPaths *[][]Room, possiblePaths [][]Room, endRoomName string) bool {
	for _, path := range possiblePaths[len(*validPaths):] {
		if path[len(path)-1].RoomName == endRoomName {
			return true
		}
	}
	return false
}

// IF a room has more than two connections (meaning there's a choice of where to move next)
// Prioritize selecting rooms with the fewest connections leading to them
func RoomSelection(sourceRoomName string, MouthRoomNames []string, tunnelGraph Graph) string {
	// select the room with shortest mouthroom array in tunnelGraph
	selectedRoom := MouthRoomNames[0]
	shortestLen := len(MouthRoomNames[0])
	for _, room := range MouthRoomNames[1:] {
		if len(tunnelGraph[room]) < shortestLen {
			shortestLen = len(room)
			selectedRoom = room
		}
	}
	return selectedRoom
}

func SomewhereToGo(connectedRoomNames []string, allRoomsMap map[string]bool) bool {
	for _, roomName := range connectedRoomNames {
		// allRoomsMap[roomName] == false if room is unvisited
		if !allRoomsMap[roomName] {
			return true
		}
	}
	return false
}

// Return true if all rooms the current last room in each path is connected to are set to 'true' in allRoomsMap (they have been visited)
func AllEndsAreDead(possiblePaths *[][]string, allRoomsMap map[string]bool, tunnelGraph Graph, endRoomName string) bool {
	// a check here, whether all possible paths end with endRoom? In which case, return true
	allReachedEnd := true
	for _, path := range *possiblePaths {
		if path[len(path)-1] != endRoomName {
			allReachedEnd = false
		}
	}
	if allReachedEnd {
		return true
	}
	for _, path := range *possiblePaths {
		for _, roomName := range path {
			for _, connectedRoomName := range tunnelGraph[roomName] {
				// fmt.Println(connectedRoomName, ":", allRoomsMap[connectedRoomName])
				if !allRoomsMap[connectedRoomName] {
					return false
				}
			}
		}
	}
	return true
}

func ReturnUnvisited(connectedRoomNames []string, allRoomsMap map[string]bool) []string {
	unvisitedRooms := []string{}
	for _, connectedRoomName := range connectedRoomNames {
		if !allRoomsMap[connectedRoomName] {
			unvisitedRooms = append(unvisitedRooms, connectedRoomName)
		}
	}
	return unvisitedRooms
}

func sliceExists(sliceOfSlices [][]string, targetSlice []string) bool {
	for _, slice := range sliceOfSlices {
		if reflect.DeepEqual(slice, targetSlice) {
			return true
		}
	}
	return false
}
