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
	antFarm.AllRoomsMap[antFarm.StartRoom.RoomName] = Room{
		RoomName:  antFarm.StartRoom.RoomName,
		IsChecked: true,
	}
	for _, cncRoomName := range antFarm.TunnelGraph[antFarm.StartRoom.RoomName] {
		antFarm.AllRoomsMap[cncRoomName] = Room{
			RoomName:  cncRoomName,
			IsChecked: false,
		}
		antFarm.PossiblePaths = append(antFarm.PossiblePaths, []Room{antFarm.AllRoomsMap[cncRoomName]})
	}

	count := 0
	if len(antFarm.PossiblePaths) > 0 {

		count++

		// proceed until there are no rooms connected to the starting room still unchecked (not included in a determinedly valid path)
		for HasPossiblePathLeft(antFarm.TunnelGraph, antFarm.AllRoomsMap, antFarm.StartRoom) {

			// TODO: change this to something that works. Issue: allRoomsMap's values falsely set to false after valid paths are found and saved
			for !ValidPathFound(&antFarm.ValidPaths, antFarm.PossiblePaths, antFarm.EndRoom.RoomName) {
				ScanForPath(antFarm.TunnelGraph, antFarm.AllRoomsMap, antFarm.EndRoom.RoomName, &antFarm.PossiblePaths, &antFarm.ValidPaths)
				count++
				if count > 5 {
					fmt.Println("Nr of ScanForPath calls: ", count)
					fmt.Println("Failurre")
					os.Exit(1)
				}
			}
			// fmt.Println(antFarm.AllRoomsMap)
			// os.Exit(1)
		}
		fmt.Println(antFarm.PossiblePaths)
	} else {
		fmt.Println("No possible paths")
		return nil
	}
	return antFarm.ValidPaths
}

// Go through possible paths, adding them to a slice of possible path slices, that is finally returned in FindValidPaths()
func ScanForPath(tunnelGraph Graph, allRoomsMap map[string]Room, endRoomName string, possiblePaths *[][]Room, validPaths *[][]Room) {
	fmt.Println("x")
	Paths := *possiblePaths
	for i, path := range Paths {
		if len(path) == 1 {
			path[0].IsChecked = true
			allRoomsMap[path[0].RoomName] = path[0]
		}
		currentMoveRoom := path[len(path)-1]
		if currentMoveRoom.RoomName == endRoomName {
			// !! Something else to do here? !!
			return
		}
		_, exists := tunnelGraph[currentMoveRoom.RoomName]
		nrOfConnections := 0
		connectedRoomNames := tunnelGraph[currentMoveRoom.RoomName]
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
				fmt.Println("Moving from", currentMoveRoom.RoomName, "to... ", connectedRoomName)
				if connectedRoomName != endRoomName {
					allRoomsMap[connectedRoomName] = Room{
						RoomName:  connectedRoomName,
						IsChecked: true,
					}
					// append room's name to corresponding subslice of possiblePaths
					fmt.Println("Appending", allRoomsMap[connectedRoomName], "to", path)
					Paths[i] = append(Paths[i], allRoomsMap[connectedRoomName])
					*possiblePaths = Paths
				} else {
					Paths[i] = append(Paths[i], allRoomsMap[connectedRoomName])
					*validPaths = append(*validPaths, Paths[i])
					fmt.Println("Voila:", *validPaths)
					UnCheckLeftOverRooms(allRoomsMap, validPaths)
					// for k := range possiblePaths[len(*validPaths):] {
					// 	if k != i {
					// 		possiblePaths[k] = []Room{possiblePaths[k][0]}
					// 	}
					// }
					return
				}
			}
		} else if nrOfConnections > 2 {
			// Alternative paths need to be added, all paths will be selected between later in the parent function FindValidPaths()
			unvisited := ReturnUnvisited(connectedRoomNames, allRoomsMap)
			for l := len(unvisited) - 1; l >= 0; l-- {
				fmt.Println("Moving from", currentMoveRoom.RoomName, "to... ", unvisited[l])
				allRoomsMap[unvisited[l]] = Room{
					RoomName:  unvisited[l],
					IsChecked: true,
				}
				if l > 0 {
					fmt.Println("Appending", allRoomsMap[unvisited[l]], "to", path)
					altPath := append(Paths[i], allRoomsMap[unvisited[l]])
					fmt.Println("altPath:", altPath)
					Paths = append(Paths, altPath)
					*possiblePaths = Paths
				} else {
					fmt.Println("Appending", allRoomsMap[unvisited[l]], "to", path)
					Paths[i] = append(Paths[i], allRoomsMap[unvisited[l]])
					*possiblePaths = Paths
				}
			}
		}

	}
	fmt.Println(*possiblePaths)
}

func HasPossiblePathLeft(tunnelGraph Graph, allRoomsMap map[string]Room, startRoom Room) bool {
	for _, cncRoomName := range tunnelGraph[startRoom.RoomName] {
		if allRoomsMap[cncRoomName].IsChecked == false {
			return true
		}
	}
	return false
}

// lower all flags of rooms not part of a valid path
func UnCheckLeftOverRooms(allRoomsMap map[string]Room, validPaths *[][]Room) {
	for _, path := range *validPaths {
		// Populate the validRooms map for quick lookups
		validRooms := make(map[string]struct{})
		for _, vRoom := range path {
			validRooms[vRoom.RoomName] = struct{}{}
		}

		// uncheck all rooms not part of a valid path
		for roomName, room := range allRoomsMap {
			if _, isValid := validRooms[roomName]; !isValid {
				room.IsChecked = false
				allRoomsMap[roomName] = room
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

func SomewhereToGo(connectedRoomNames []string, allRoomsMap map[string]Room) bool {
	for _, roomName := range connectedRoomNames {
		if !allRoomsMap[roomName].IsChecked {
			return true
		}
	}
	return false
}

func ReturnUnvisited(connectedRoomNames []string, allRoomsMap map[string]Room) []string {
	unvisitedRooms := []string{}
	for _, connectedRoomName := range connectedRoomNames {
		if !allRoomsMap[connectedRoomName].IsChecked {
			unvisitedRooms = append(unvisitedRooms, connectedRoomName)
		}
	}
	return unvisitedRooms
}
