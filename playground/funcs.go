package playground

// addEdge adds a directed edge from node1 to node2
func (g Graph) addEdge(node1, node2 Room) {
	g[node1.RoomName] = append(g[node1.RoomName], node2.RoomName)
	g[node2.RoomName] = append(g[node2.RoomName], node1.RoomName)
}
