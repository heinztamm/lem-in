package main

import (
	"os"

	"01.kood.tech/git/kartamm/lem-in/playground"
)

func main() {

	examplefilepath := os.Args[1]

	playground.CreateGraph(examplefilepath)
}
