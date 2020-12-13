package main

import "fmt"

// Graph shows how to use a map of maps to represent a directed graph.
// 图表显示了如何使用地图来表示有向图。
var graph = make(map[string]map[string]bool)

// add map to map
func add_edge(from, to string) {
	edge := graph[from]

	if edge == nil {
		edge = make(map[string]bool)
		// fmt.Printf("before %v\n", graph[from])
		graph[from] = edge
		fmt.Printf("after %v\n", graph[from])

	}

	edge[to] = true
		fmt.Printf("after %v\n", graph[from])

}

func has_edge(from, to string) bool {
	return graph[from][to]
}

func main() {
	add_edge("a", "b")
	add_edge("a", "d")
	add_edge("c", "b")



	fmt.Println(has_edge("a", "b"))
}