package playground

import (
	"fmt"
	"sort"
	"strconv"
)

func Movement(numberOfAnts int, startRoom string, endRoom string, paths [][]string) { // Simulate the movement and print results to terminal
	for i := range paths {
		paths[i] = append([]string{startRoom}, paths[i]...)
	}
	paths = ReOrderPaths(paths)
	antLocations := make(map[int]string) // create a map AntName equals to location. We can keep track of ants
	for i := 1; i <= numberOfAnts; i++ { // initialize all the ants in the starting room
		antLocations[i] = startRoom
	}
	for len(antLocations) > 0 { // Loops until all the ants are at the end. One loop is one turn
		var ants []int                  // create array based on the map
		for ant := range antLocations { // 1,2,3,4,5...
			ants = append(ants, ant)
		}
		sort.Ints(ants)            // sort them in ascending order
		var directLinkUsed bool    // bool used to detect if paths contain direct linkage between starting room and ending room. Changes movement a bit
		for _, ant := range ants { // used ants array for looping due to golang loops over maps in random order. I need same order every loop.
			if antLocations[ant] == endRoom { // ant reached end
				delete(antLocations, ant) // delete the ant, but it is still in the ants array. Otherwise it will break the ants looping
			} else {
				roomsAvailable := getNextRooms(antLocations[ant], &directLinkUsed, startRoom, endRoom, antLocations, paths) // Find all possible movement options
				if len(roomsAvailable) > 0 {                                                                                // We have a option to move somewhere
					antLocations[ant] = roomsAvailable[0]                              // ant moves to the first available room
					fmt.Print("L" + strconv.Itoa(ant) + "-" + roomsAvailable[0] + " ") // print the movement
				} // If no rooms then the ant will do nothing. It means it will wait for the next turn
			}
		}
		if len(antLocations) != 0 { //unless last line, move to newline for next turn
			fmt.Println()
		}
	}
}
func getNextRooms(antLoc string, directLinkUsed *bool, startRoom string, endRoom string, antLocations map[int]string, paths [][]string) []string { // Find all possible movement options
	// Find connected links from this room
	var nextRooms []string
	for _, path := range paths { // find current room progress from all the paths
		for index, content := range path {
			if content == antLoc {
				nextRooms = append(nextRooms, path[index+1]) // add the next room from that
			}
		}
	}
	// Check if any of those links are free to move into
	var roomsAvailable []string
	if len(antLocations) < 5 && antLoc == startRoom && *directLinkUsed == true { // special check-up for mazes that have direct links between start and end
		return roomsAvailable[:0] // last ants will wait out for the direct link to free up instead starting the long route. More time-efficient
	}
	for _, room := range nextRooms { // loop through the links from the room ant is located
		var occupied bool
		if room == endRoom && *directLinkUsed == false { // if no directLink detected and next room is the end
			roomsAvailable = append(roomsAvailable, room) // add it as possible movement
			if antLoc == startRoom {                      // that ant now used directlink
				*directLinkUsed = true
			}
		} else {
			for _, locs := range antLocations { // otherwise loop as normal
				if room == locs { // check if any of the possible links are already occupied
					occupied = true
					break
				}
			}
		}
		if !occupied { // if the room was empty. No ants in it
			roomsAvailable = append(roomsAvailable, room) // add it as possible movement
		}
	}
	return roomsAvailable // return the possible movements
}
