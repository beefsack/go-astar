package astar

import "testing"

type TestPather struct{}

func (tp *TestPather) Neighbors() []Pather {
	return []Pather{}
}

func (tp *TestPather) Cost(to Pather) float64 {
	return 1
}

func TestSomething(t *testing.T) {
	Path(&TestPather{}, &TestPather{})
}
