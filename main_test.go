package main

import "testing"

func validOrder(n int, edges [][2]int, order []int) bool {
	if len(order) != n {
		return false
	}
	pos := make([]int, n)
	for i, v := range order {
		pos[v] = i
	}
	for _, e := range edges {
		if pos[e[0]] >= pos[e[1]] {
			return false
		}
	}
	return true
}

func TestTopoSort(t *testing.T) {
	edges := [][2]int{{5, 2}, {5, 0}, {4, 0}, {4, 1}, {2, 3}, {3, 1}}
	order, ok := TopoSort(6, edges)
	if !ok || !validOrder(6, edges, order) {
		t.Errorf("order %v invalid", order)
	}

	// Cyclic graph has no order.
	if _, ok := TopoSort(2, [][2]int{{0, 1}, {1, 0}}); ok {
		t.Error("cycle must report no order")
	}

	// Edgeless graph: identity order (deterministic smallest-first).
	order, ok = TopoSort(3, nil)
	if !ok || order[0] != 0 || order[1] != 1 || order[2] != 2 {
		t.Errorf("edgeless order = %v", order)
	}

	// Chain forces the unique order.
	order, _ = TopoSort(3, [][2]int{{2, 1}, {1, 0}})
	if order[0] != 2 || order[1] != 1 || order[2] != 0 {
		t.Errorf("chain order = %v", order)
	}
}
