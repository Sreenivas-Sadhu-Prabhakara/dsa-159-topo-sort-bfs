# Topological Sort (Kahn's BFS)

> Striver SDE Sheet 159 · Day 23 Graph — repeatedly emit a vertex with no unmet prerequisites; indegree counters make "no prerequisites" an O(1) test. Go, stdlib only, tested.

## Problem

Order a directed acyclic graph's vertices so every edge points forward — the dependency-resolution problem (build systems, course schedules, package managers). Report failure on cyclic input.

## Approach

Kahn's algorithm is the executable definition: a vertex with indegree zero has no unmet dependency, so emit it and release its outgoing edges, decrementing targets' indegrees; newly-zero vertices join the ready pool. Falling short of n emissions means a cycle (same residue argument as #158 — detection comes free). One deliberate upgrade here: the ready pool is kept sorted so ties break toward smaller vertices, making the output canonical and exactly testable; a plain queue is fine when any valid order will do. The validity *property* (every edge points forward) is still what the main test checks.

**Complexity:** O(V+E) with a plain queue; the deterministic pool adds a small insertion cost.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
