package playground

type AntFarm struct {
	PossiblePaths [][]Room
	ValidPaths    [][]Room
	AllRoomsMap   map[string]Room
	StartRoom     Room
	EndRoom       Room
	TunnelGraph   Graph
	AntNr         int
}
type Room struct {
	RoomName  string
	IsChecked bool
}

// can the graph type be changed to accommodate a boolean value as well? meaning, the value would be ([]string, bool)
type Graph map[string][]string
