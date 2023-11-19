package playground

type AntFarm struct {
	PossiblePaths [][]Room // keys for the map of all room names, filled in order of following a path
	AllRoomsMap   map[string]Room
	StartRoom     StartRoom
	EndRoom       EndRoom
	TunnelGraph   Graph
	AntNr         int
}
type Room struct {
	RoomName  string
	IsChecked bool
}

type StartRoom struct {
	RoomName string
	AntCount int
}

type EndRoom struct {
	RoomName string
	AntCount int
}

// can the graph type be changed to accommodate a boolean value as well? meaning, the value would be ([]string, bool)
type Graph map[string][]Room
