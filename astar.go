package astar

// astar is an A* pathfinding implementation.

// Pather is an interface which allows A* searching on arbitrary objects which
// can represent a weighted graph.
type Pather interface {
	// PathNeighbors returns the direct neighboring nodes of this node which
	// can be pathed to.
	PathNeighbors() []Pather
	// PathCost calculated the exact movement cost to neighbor nodes, and
	// approximates the movement cost to any non-adjacent nodes.
	PathCost(to Pather) float64
}

// node is a wrapper to store A* data for a Pather node.
type node struct {
	pather Pather
	cost   float64
	rank   float64
	parent *node
	open   bool
	closed bool
}

// nodeMap is a collection of nodes keyed by Pather nodes for quick reference.
type nodeMap map[Pather]*node

// get gets the Pather object wrapped in a node, instantiating if required.
func (nm nodeMap) get(p Pather) *node {
	n, ok := nm[p]
	if !ok {
		n = &node{
			pather: p,
		}
		nm[p] = n
	}
	return n
}

// lowestOpen gets the lowest ranked open node in the node map.
//
// The storage and / or searching for the lowest ranked open node needs to be
// optimised.
func (nm nodeMap) lowestOpen() (n *node) {
	for _, i := range nm {
		if i.open && (n == nil || i.rank < n.rank) {
			n = i
		}
	}
	return
}

// Path calculates a short path and the distance between the two Pather nodes.
//
// If no path is found, it will return nil.
//
// See:
func Path(from, to Pather) ([]Pather, float64) {
	nm := nodeMap{}
	fromNode := nm.get(from)
	fromNode.open = true
	for {
		current := nm.lowestOpen()
		if current == nil {
			// There's no path, return nil.
			return nil, 0
		}
		current.open = false
		current.closed = true
		for _, neighbor := range current.pather.PathNeighbors() {
			cost := current.cost + current.pather.PathCost(neighbor)
			neighborNode := nm.get(neighbor)
			if neighbor == to {
				// Found a path to the goal.
				p := []Pather{}
				curr := neighborNode
				curr.parent = current
				for curr != nil {
					p = append(p, curr.pather)
					curr = curr.parent
				}
				return p, cost
			}
			if cost < neighborNode.cost {
				neighborNode.open = false
				neighborNode.closed = false
			}
			if !neighborNode.open && !neighborNode.closed {
				neighborNode.cost = cost
				neighborNode.open = true
				neighborNode.rank = cost + neighbor.PathCost(to)
				neighborNode.parent = current
			}
		}
	}
}
