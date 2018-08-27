package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateGraph(t *testing.T) {
	graph := CreateGraph()
	if graph.NumberOfNodes() != 0 {
		t.Error("The expected result for the number of nodes is ", 0)
	}
	fmt.Fprintln(os.Stderr, graph)

	node1 := CreateNode(0)
	graph.AddNode(node1)
	fmt.Fprintln(os.Stderr, node1)
	fmt.Fprintln(os.Stderr, graph)
	if graph.NumberOfNodes() != 1 {
		t.Error("The expected result for the number of nodes is ", 1)
	}

	node2 := CreateNode(1)
	graph.AddNode(node2)
	fmt.Fprintln(os.Stderr, node2)
	fmt.Fprintln(os.Stderr, graph)
	if graph.NumberOfNodes() != 2 {
		t.Error("The expected result for the number of nodes is ", 2)
	}

	node3 := CreateNode(2)
	graph.AddNode(node3)
	fmt.Fprintln(os.Stderr, node3)
	fmt.Fprintln(os.Stderr, graph)
	if graph.NumberOfNodes() != 3 {
		t.Error("The expected result for the number of nodes is ", 3)
	}

	if !graph.NodeExists(0) {
		t.Error("The test to verify that node 0 exists should have return  ", true)
	}

	if !graph.NodeExists(1) {
		t.Error("The test to verify that node 1 exists should have return  ", true)
	}

	if !graph.NodeExists(2) {
		t.Error("The test to verify that node 2 exists should have return  ", true)
	}

	if graph.NodeExists(3) {
		t.Error("The test to verify that node 3 exists should have return  ", false)
	}

	node1.AddLink(1)
	node2.AddLink(0)
	node1.AddLink(2)
	node3.AddLink(0)
	fmt.Fprintln(os.Stderr, node1)
	fmt.Fprintln(os.Stderr, node2)
	fmt.Fprintln(os.Stderr, node3)
	if node1.Degree() != 2 {
		t.Error("The degree of node1 should be ", 2)
	}

	if node2.Degree() != 1 {
		t.Error("The degree of node2 should be ", 1)
	}

	if node3.Degree() != 1 {
		t.Error("The degree of node3 should be ", 1)
	}

	graph.MarkAsGateway(2)
	gateway := graph.RetrieveNode(2)
	fmt.Fprintln(os.Stderr, graph)
	fmt.Fprintln(os.Stderr, node2)
	fmt.Fprintln(os.Stderr, gateway)
	if !gateway.IsGateway() {
		t.Error("Node 3 should be a gateway and the method IsGateway should return", true)
	}
	path := graph.ShortestPath(0, gateway.id)
	fmt.Fprintln(os.Stderr, path)
	fmt.Println(fmt.Sprintf("%d %d", path[1], path[0]))
	if path[1] != 0 && path[0] != 2 {
		t.Error("The last part of the path should be ", fmt.Sprintf("%d %d", path[1], path[0]))
	}
}
