package main

import (
	"fmt"
)

type Node struct {
	distance int
	neighbor int
}

type Graph struct {
	adjList map[int][]Node
}

func (graph *Graph) addVertex(vertex int) {

	if _, isPresent := graph.adjList[vertex]; isPresent {
		return
	} else {
		graph.adjList[vertex] = []Node{}
	}
}

func (graph *Graph) addEdge(from, to, dist int) {

	graph.addVertex(from)
	graph.addVertex(to)

	// bi-directional graph

	// checking to avoid multiple edges between same nodes
	if !graph.containsEdge(from, to) {
		graph.adjList[from] = append(graph.adjList[from], Node{distance: dist, neighbor: to})
	}
}

func (graph *Graph) containsEdge(from, to int) bool {
	for _, neighbors := range graph.adjList[from] {
		if neighbors.neighbor == to {
			return true
		}
	}

	return false
}

func (graph *Graph) DeleteEdge(from, to int) {
	for index, neighbors := range graph.adjList[from] {
		if neighbors.neighbor == to {
			graph.adjList[from] = append(graph.adjList[from][:index], graph.adjList[from][index+1:]...)
		}
	}
}

func main() {
	fmt.Println("Let start with graphs")

	graph := Graph{adjList: map[int][]Node{}}

	graph.addEdge(1, 2, 4)
	graph.addEdge(1, 4, 9)
	graph.addEdge(5, 1, 6)
	graph.addEdge(2, 3, 7)
	graph.addEdge(4, 2, 3)

	fmt.Println(graph.containsEdge(2, 5))
	fmt.Println(graph.containsEdge(2, 4))
	fmt.Println(graph.containsEdge(2, 3))

	graph.DeleteEdge(5, 1)

	for vertex, neighbors := range graph.adjList {
		fmt.Printf("Vertex %v: ", vertex)
		for _, neighbor := range neighbors {
			fmt.Printf("(%v, %v) ", neighbor.neighbor, neighbor.distance)
		}

		fmt.Println()
	}

}
