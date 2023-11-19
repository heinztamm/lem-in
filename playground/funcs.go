package playground

import "fmt"

// addEdge adds a directed edge from node1 to node2
func (g Graph) addDirectedEdge(node1Name string, node2 Room) {
	g[node1Name] = append(g[node1Name], node2)
}

// printGraph prints the adjacency list representation of the graph
func (g Graph) PrintGraph() {
	for node, neighbors := range g {
		fmt.Printf("%s: %v\n", node, neighbors)
	}
}
