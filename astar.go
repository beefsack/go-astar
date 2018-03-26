package astar

import (
	"container/heap"
	//"fmt"
)

// astar is an A* pathfinding implementation.

// Pather is an interface which allows A* searching on arbitrary objects which
// can represent a weighted graph.
type Pather interface {
	// PathNeighbors returns the direct neighboring nodes of this node which
	// can be pathed to.
	PathNeighbors() []Pather
	// PathNeighborCost calculates the exact movement cost to neighbor nodes.
	PathNeighborCost(to Pather) float64
	// PathEstimatedCost is a heuristic method for estimating movement costs
	// between non-adjacent nodes.
	PathEstimatedCost(to Pather) float64
}

// node is a wrapper to store A* data for a Pather node.
type node struct {
	pather Pather
	cost   float64
	rank   float64
	parent *node
	open   bool
	closed bool
	index  int
}

// nodeMap is a collection of nodes keyed by Pather nodes for quick reference.
type nodeMap map[Pather]*node

// get gets the Pather object wrapped in a node, instantiating if required.
func (nm nodeMap) get_node_from_pather(p Pather) *node {
	n, ok := nm[p]
	if !ok {
		n = &node{
			pather: p,
		}
		nm[p] = n
	}
	return n
}

func expand(nm nodeMap, nq *priorityQueue, curnode *node, dest Pather) {
	for _, neighbor := range curnode.pather.PathNeighbors() {

		cost := curnode.cost + curnode.pather.PathNeighborCost(neighbor)
		neighborNode := nm.get_node_from_pather(neighbor)

		if cost < neighborNode.cost {
			if neighborNode.open {
				heap.Remove(nq, neighborNode.index)
			}
			neighborNode.open = false
			neighborNode.closed = false
		}
		if !neighborNode.open && !neighborNode.closed {
			neighborNode.cost = cost
			neighborNode.open = true
			neighborNode.rank = cost + neighbor.PathEstimatedCost(dest)
			neighborNode.parent = curnode
			heap.Push(nq, neighborNode)
		}
	}
}

// Path calculates a short path and the distance between the two Pather nodes.
//
// If no path is found, found will be false.
func Path(from Pather, to Pather) (path []Pather, distance float64, found bool) {
	fwd_nodemap := nodeMap{}
	fwd_nq := &priorityQueue{}

	heap.Init(fwd_nq)

	fromNode := fwd_nodemap.get_node_from_pather(from)
	fromNode.open = true

	heap.Push(fwd_nq, fromNode)

	for {
		if fwd_nq.Len() == 0 {
			// There's no path, return found false.
			return
		}

		fwd_curnode := heap.Pop(fwd_nq).(*node)
		fwd_curnode.open = false
		fwd_curnode.closed = true
		fwd_pather := fwd_curnode.pather

		if fwd_pather == to {
			// Found a path to the goal.
			p := []Pather{}

			curr := fwd_curnode
			for curr != nil {
				p = append(p, curr.pather)
				curr = curr.parent
			}

			return p, fwd_curnode.cost, true
		}

		expand(fwd_nodemap, fwd_nq, fwd_curnode, to)
	}
}

// Path calculates a short path and the distance between the two Pather nodes.
//
// If no path is found, found will be false.
func PathBidir(from Pather, to Pather) (path []Pather, distance float64, found bool) {
	fwd_nodemap := nodeMap{}
	fwd_nq := &priorityQueue{} // fwd priq

	rev_nodemap := nodeMap{}
	rev_nq := &priorityQueue{} // rev priq

	heap.Init(fwd_nq)
	heap.Init(rev_nq)

	fromNode := fwd_nodemap.get_node_from_pather(from)
	fromNode.open = true

	toNode := rev_nodemap.get_node_from_pather(to)
	toNode.open = true

	heap.Push(fwd_nq, fromNode)
	heap.Push(rev_nq, toNode)

	for {
		if fwd_nq.Len() == 0 || rev_nq.Len() == 0 {
			// There's no path, return found false.
			return
		}

		fwd_curnode := heap.Pop(fwd_nq).(*node)
		fwd_curnode.open = false
		fwd_curnode.closed = true
		fwd_pather := fwd_curnode.pather

		rev_curnode := heap.Pop(rev_nq).(*node)
		rev_curnode.open = false
		rev_curnode.closed = true

		fwd_node_in_rev_map := rev_nodemap.get_node_from_pather(fwd_pather)
		if fwd_node_in_rev_map.closed || fwd_node_in_rev_map.open || fwd_pather == to {
			// Found a path to the goal.
			rp := []Pather{}
			curr := fwd_node_in_rev_map
			for curr != nil {
				rp = append(rp, curr.pather)
				curr = curr.parent
			}

			// reverse the path from touchpoint to "to"
			p := []Pather{}
			for i := len(rp) - 1; i >= 0; i = i - 1 {
				p = append(p, rp[i])
			}
			curr = fwd_curnode.parent
			for curr != nil {
				p = append(p, curr.pather)
				curr = curr.parent
			}

			return p, fwd_curnode.cost + fwd_node_in_rev_map.cost, true
		}

		expand(fwd_nodemap, fwd_nq, fwd_curnode, to)
		expand(rev_nodemap, rev_nq, rev_curnode, from)

	}
}
