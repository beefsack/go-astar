package astar

type Pather interface {
	Neighbors() []Pather
	Cost(to Pather) float64
}

type openPather struct {
	pather Pather
	rank   float64
	parent *openPather
}

type pathQueue struct {
	Values []*openPather
}

func (q *pathQueue) Add(p *openPather) {

}

func (q *pathQueue) Remove(p *openPather) {
}

func (q *pathQueue) Pop() *openPather {
	return q.Values[0]
}

func (q *pathQueue) Len() int {
	return len(q.Values)
}

func Path(from, to Pather) []Pather {
	openQueue := openQueue{}
	openIndex := map[string]*openPather{
		from.LocationIdentifier(): {
			pather: from,
			rank:   0,
			parent: nil,
		},
	}
	for len(open) > 0 && open[0].pather != to {
		open = open[1:]
	}
	if len(open) == 0 {
		return nil
	}
	curr := &open[0]
	path := []Pather{}
	for curr != nil {
		path = append(path, curr.pather)
		curr = curr.parent
	}
	return nil
}
