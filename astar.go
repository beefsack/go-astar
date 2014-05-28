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

func Path(from, to Pather) []Pather {
	open := []openPather{
		{
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
