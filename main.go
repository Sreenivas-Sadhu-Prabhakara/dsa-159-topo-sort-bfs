package main

import "fmt"

// TopoSort returns a topological order of a directed graph via Kahn's
// algorithm; ok is false when the graph has a cycle (no order exists).
// Smaller vertices are preferred among the simultaneously available, making
// the output deterministic.
func TopoSort(n int, edges [][2]int) ([]int, bool) {
	adj := make([][]int, n)
	indegree := make([]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		indegree[e[1]]++
	}
	// A sorted "ready" list instead of a plain queue keeps output canonical.
	ready := []int{}
	for v := n - 1; v >= 0; v-- {
		if indegree[v] == 0 {
			ready = append(ready, v) // descending, so pop from the end = smallest
		}
	}
	out := make([]int, 0, n)
	for len(ready) > 0 {
		v := ready[len(ready)-1]
		ready = ready[:len(ready)-1]
		out = append(out, v)
		for _, w := range adj[v] {
			indegree[w]--
			if indegree[w] == 0 {
				// Insert keeping ready sorted descending.
				i := len(ready)
				ready = append(ready, 0)
				for i > 0 && ready[i-1] < w {
					ready[i] = ready[i-1]
					i--
				}
				ready[i] = w
			}
		}
	}
	if len(out) < n {
		return nil, false
	}
	return out, true
}

func main() {
	order, ok := TopoSort(6, [][2]int{{5, 2}, {5, 0}, {4, 0}, {4, 1}, {2, 3}, {3, 1}})
	fmt.Println(order, ok)
}
