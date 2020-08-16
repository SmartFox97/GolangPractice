package main

import (
	"fmt"
)

func main() {
	var n , m int
	fmt.Scan(&n,&m)
	var graph [][]int
	for i := 0; i < n; i++ {
		inline := make([]int, m)
		graph = append(graph, inline)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&graph[i][j])
		}
	}

	var count int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if graph[i][j] == 1 {
				search(graph, i, j,n ,m)
				count++
			}
		}
	}
	fmt.Println(count)
}

func search(graph1 [][]int, i, j, n, m int) {
	if i < 0 || i >= n || j >= m || j < 0 || graph1[i][j] == 0 {
		return
	}
	graph1[i][j] = 0
	search(graph1, i-1, j, n, m)
	search(graph1, i+1, j, n, m)
	search(graph1, i, j+1, n, m)
	search(graph1, i, j-1, n, m)

}
