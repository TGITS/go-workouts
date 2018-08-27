package main

import (
	"fmt"
	"os"
	"sort"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// N: the total number of nodes in the level, including the gateways
	// L: the number of links
	// E: the number of exit gateways
	var N, L, E int
	fmt.Scan(&N, &L, &E)

	graph := CreateGraph()

	for i := 0; i < L; i++ {
		// N1: N1 and N2 defines a link between these nodes
		var N1, N2 int
		fmt.Scan(&N1, &N2)

		var node1, node2 *Node
		if !graph.NodeExists(N1) {
			node1 = CreateNode(N1)
			graph.AddNode(node1)
		} else {
			node1 = graph.RetrieveNode(N1)
			fmt.Fprintln(os.Stderr, node1)
		}

		if !graph.NodeExists(N2) {
			node2 = CreateNode(N2)
			fmt.Fprintln(os.Stderr, node2)
			graph.AddNode(node2)
		} else {
			node2 = graph.RetrieveNode(N2)
			fmt.Fprintln(os.Stderr, node2)
		}
		node1.AddLink(N2)
		node2.AddLink(N1)
	}

	var gateways []*Node
	for i := 0; i < E; i++ {
		// EI: the index of a gateway node
		var EI int
		fmt.Scan(&EI)
		graph.MarkAsGateway(EI)
		gateways = append(gateways, graph.RetrieveNode(EI))
	}
	for {
		// SI: The index of the node on which the Skynet agent is positioned this turn
		var SI int
		fmt.Scan(&SI)

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		shortestPathByGateway := make(map[int][]int)
		for _, gateway := range gateways {
			shortestPathByGateway[gateway.id] = graph.ShortestPath(SI, gateway.id)
		}

		var shortestOfShortestPath []int
		for _, value := range shortestPathByGateway {
			if shortestOfShortestPath == nil {
				shortestOfShortestPath = value
			} else {
				if len(shortestOfShortestPath) > len(value) {
					shortestOfShortestPath = value
				}
			}
		}
		// Example: 0 1 are the indices of the nodes you wish to sever the link between
		//fmt.Println("0 1")
		fmt.Println(fmt.Sprintf("%d %d", shortestOfShortestPath[1], shortestOfShortestPath[0]))
	}
}

// Node is a struct to represent the node of a graph
type Node struct {
	id          int
	adjacencies []int
	gateway     bool
}

//Graph is a struct to represent the graph in its entierety
//It's a Map of pointers on node, the keys of the map are the id of the nodes
//The idea is to be able to retrieve fast a node from its id
//and to have a data structure that store all the nodes used in the program
type Graph struct {
	nodes map[int]*Node
}

//Row is a struct that is used in the computing of the shortest path from a node to another
//as a row of the distance table, materialized by a map[int]Row
type Row struct {
	Distance        int
	PrecedingNodeID int
}

//CreateNode is a builder method to create a node
func CreateNode(id int) *Node {
	return &Node{
		id:          id,
		adjacencies: nil,
		gateway:     false,
	}
}

//CreateGraph creates and initializes a Graph
func CreateGraph() *Graph {
	return &Graph{
		nodes: make(map[int]*Node),
	}
}

//Degree gives the degree of the node, this is the number of adjacency nodes
func (n *Node) Degree() int {
	return len(n.adjacencies)
}

// AddLink is a method to add a node linked to the node on which the method is applied
func (n *Node) AddLink(destination int) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf("Node on which we add a link : %v", n))
	if n.adjacencies == nil || len(n.adjacencies) == 0 {
		fmt.Fprintln(os.Stderr, "The adjacency list is currently nil")
		n.adjacencies = append(n.adjacencies, destination)
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Adjacency list %v after adding %d", n.adjacencies, destination))
	} else {
		index := sort.SearchInts(n.adjacencies, destination)
		if index >= len(n.adjacencies) || n.adjacencies[index] != destination {
			n.adjacencies = append(n.adjacencies, destination)
			sort.Ints(n.adjacencies)
		}
	}
}

// RemoveLink is a method to remove a node linked to the node on which the method is applied
func (n *Node) RemoveLink(id int) {
	index := sort.SearchInts(n.adjacencies, id)
	if n.adjacencies[index] == id {
		n.adjacencies = append(n.adjacencies[:index], n.adjacencies[index+1:]...)
	}
}

//IsGateway return true if the node is a gateway, false otherwise
func (n *Node) IsGateway() bool {
	return n.gateway
}

//MarkAsGateway modifies the flag that indicates that a node is a gateway or not.
func (n *Node) MarkAsGateway() {
	n.gateway = true
}

//RetrieveNode is a method to get a node from its index
func (g *Graph) RetrieveNode(id int) *Node {
	return g.nodes[id]
}

//AddNode is a method to add a node to the graph, if it is not already present
func (g *Graph) AddNode(n *Node) {
	id := n.id
	_, ok := g.nodes[id]
	if !ok {
		g.nodes[id] = n
	}
}

//NodeExists is a method to test if a node has already been created
func (g *Graph) NodeExists(id int) bool {
	_, ok := g.nodes[id]
	return ok
}

//MarkAsGateway modifies find the node identified by the provided id and modifies the flag
//that indicates this node is a gateway or not.
func (g *Graph) MarkAsGateway(id int) {
	n, ok := g.nodes[id]
	if ok {
		n.gateway = true
	}
}

//ShortestPath finds the shortest path in the graph between the two nodes which the ids are provided in parameter
//The Path is materialized by a slide of id of node
func (g *Graph) ShortestPath(source, destination int) []int {
	var path []int
	distanceTable := make(map[int]*Row)
	for key := range g.nodes {
		if key == source {
			distanceTable[key] = &Row{Distance: 0, PrecedingNodeID: key}
		} else {
			distanceTable[key] = &Row{Distance: -1, PrecedingNodeID: -1}
		}
	}
	fmt.Fprintln(os.Stderr, distanceTable)
	var processingQueue []int
	processingQueue = append(processingQueue, source)
	enqueuedNodes := make(map[int]bool)
	for len(processingQueue) > 0 {
		currentNodeID := processingQueue[0]
		processingQueue = processingQueue[1:]
		currentNode := g.RetrieveNode(currentNodeID)
		for _, id := range currentNode.adjacencies {
			currentRow := distanceTable[id]
			if currentRow.PrecedingNodeID == -1 {
				currentRow.PrecedingNodeID = currentNodeID
				currentRow.Distance = distanceTable[currentNodeID].Distance + 1
			}
			if _, present := enqueuedNodes[id]; !present {
				processingQueue = append(processingQueue, id)
				enqueuedNodes[id] = true
			}
		}
	}

	currentNodeID := destination
	currentRow := distanceTable[currentNodeID]
	precedingNodeID := currentRow.PrecedingNodeID
	path = append(path, currentNodeID)
	for currentNodeID != source {
		currentNodeID = precedingNodeID
		currentRow = distanceTable[currentNodeID]
		precedingNodeID = currentRow.PrecedingNodeID
		path = append(path, currentNodeID)
	}

	return path
}

//NumberOfNodes return the number of nodes in the graph
func (g *Graph) NumberOfNodes() int {
	return len(g.nodes)
}
