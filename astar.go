package astar

type Pather interface {
	Neighbors() []Pather
	Cost(to Pather) float64
}

type node struct {
	pather Pather
	cost   float64
	rank   float64
	parent *node
	open   bool
	closed bool
}

type nodeMap map[Pather]*node

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

func (nm nodeMap) lowestOpen() (n *node) {
	for _, i := range nm {
		if i.open && (n == nil || i.rank < n.rank) {
			n = i
		}
	}
	return
}

func Path(from, to Pather) ([]Pather, float64) {
	nm := nodeMap{}
	fromNode := nm.get(from)
	fromNode.open = true
	for {
		current := nm.lowestOpen()
		if current == nil {
			// There's no path
			return nil, 0
		}
		current.open = false
		current.closed = true
		for _, neighbor := range current.pather.Neighbors() {
			cost := current.cost + current.pather.Cost(neighbor)
			neighborNode := nm.get(neighbor)
			if neighbor == to {
				// Goal!
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
				neighborNode.rank = cost + neighbor.Cost(to)
				neighborNode.parent = current
			}
		}
	}
}
