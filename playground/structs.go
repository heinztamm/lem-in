package playground

type AntFarm struct {
	PossiblePaths [][]string
	ValidPaths3D  [][][]string
	AllRoomsMap   map[string]bool
	StartRoom     Room
	EndRoom       Room
	TunnelGraph   Graph
	AntNr         int
}
type Room struct {
	RoomName  string
	IsChecked bool
}

type Graph map[string][]string
