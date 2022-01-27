package astar

import "container/heap"

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
	index  int
}

// nodeMap is a collection of nodes keyed by Pather nodes for quick reference.
type nodeMap map[Pather]*node

// Path calculates a short path and the distance between the two Pather nodes.
//
// If no path is found, found will be false.
func Path(from, to Pather) (path []Pather, distance float64, found bool) {
	fromNode := &node{pather: from}
	closeset := nodeMap{from: fromNode}
	openset := &priorityQueue{fromNode}
	heap.Init(openset)
	for openset.Len() > 0 {
		current := heap.Pop(openset).(*node)

		if current.pather == to {
			// Found a path to the goal.
			p := []Pather{}
			curr := current
			for curr != nil {
				p = append(p, curr.pather)
				curr = curr.parent
			}
			return p, current.cost, true
		}

		for _, neighbor := range current.pather.PathNeighbors() {
			neighborNode, exists := closeset[neighbor]
			if !exists {
				neighborNode = &node{pather: neighbor}
				closeset[neighbor] = neighborNode
			}
			cost := current.cost + current.pather.PathNeighborCost(neighbor)
			if !exists || cost < neighborNode.cost {
				neighborNode.cost = cost
				neighborNode.rank = cost + neighbor.PathEstimatedCost(to)
				neighborNode.parent = current
				if exists {
					heap.Fix(openset, neighborNode.index)
				} else {
					heap.Push(openset, neighborNode)
				}
			}
		}
	}
	return
}
