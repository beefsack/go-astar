package astar

import "testing"

func testPath(worldInput string, t *testing.T, expectedDist float64) {
	world := ParseWorld(worldInput)
	t.Logf("Input world\n%s", world.RenderPath([]Pather{}))
	p, dist := Path(world.From(), world.To())
	if p == nil {
		t.Log("Could not find a path")
	} else {
		t.Logf("Resulting path\n%s", world.RenderPath(p))
	}
	if p == nil && expectedDist >= 0 {
		t.Fatal("Could not find a path")
	}
	if p != nil && dist != expectedDist {
		t.Fatalf("Expected dist to be %v but got %v", expectedDist, dist)
	}
}

func TestStraightLine(t *testing.T) {
	testPath(`
.....~......
.....MM.....
.F........T.
....MMM.....
............
`, t, 9)
}

func TestPathAroundMountain(t *testing.T) {
	testPath(`
.....~......
.....MM.....
.F..MMMM..T.
....MMM.....
............
`, t, 13)
}

func TestBlocked(t *testing.T) {
	testPath(`
............
.........XXX
.F.......XTX
.........XXX
............
`, t, -1)
}

func TestMaze(t *testing.T) {
	testPath(`
FX.X........
.X...XXXX.X.
.X.X.X....X.
...X.X.XXXXX
.XX..X.....T
`, t, 27)
}

func TestMountainClimber(t *testing.T) {
	testPath(`
..F..M......
.....MM.....
....MMMM..T.
....MMM.....
............
`, t, 12)
}

func TestRiverSwimmer(t *testing.T) {
	testPath(`
.....~......
.....~......
.F...X...T..
.....M......
.....M......
`, t, 11)
}
