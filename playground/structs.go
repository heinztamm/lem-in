package playground

type Room struct {
	RoomName               any
	Next                   *Room
	ConnectedRoomAddresses []*Room
	IsChecked              bool
}

type Path struct {
	Head   *Room
	Tail   *Room
	Length int
}

type StartRoom struct {
	RoomName any
	Next     *Room
}

type EndRoom struct {
	RoomName any
}

type Graph map[any][]any
