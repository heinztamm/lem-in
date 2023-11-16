package playground

import "fmt"

// addEdge adds a directed edge from node1 to node2
func (g Graph) addDirectedEdge(node1, node2 string) {
	g[node1] = append(g[node1], node2)
}

// printGraph prints the adjacency list representation of the graph
func (g Graph) PrintGraph() {
	for node, neighbors := range g {
		fmt.Printf("%s: %v\n", node, neighbors)
	}
}
