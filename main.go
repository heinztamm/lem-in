package main

import (
	"fmt"
	"os"

	"01.kood.tech/git/kartamm/lem-in/playground"
)

func main() {

	examplefilepath := os.Args[1]
	antFarm := &playground.AntFarm{}
	playground.FetchParams(examplefilepath, antFarm)

	for _, path := range playground.FindValidPaths(antFarm) {
		fmt.Println(path)
	}
}
