package grids

type BoolGrid struct {
	data          []bool
	height, width int
}

func CreateBoolGrid(height, width int) BoolGrid {
	data := make([]bool, height*width)
	return BoolGrid{height: height, width: width, data: data}
}

func (g *BoolGrid) IsInsideBounds(x, y int) bool {
	return y >= 0 && y < g.width && x >= 0 && x < g.height
}

func (g *BoolGrid) Get(x, y int) bool {
	return g.data[g.getIndex(x, y)]
}

func (g *BoolGrid) Set(x, y int) {
	g.data[g.getIndex(x, y)] = true
}

func (g *BoolGrid) Reset(x, y int) {
	g.data[g.getIndex(x, y)] = false
}

func (g *BoolGrid) CountTrue() (count int) {
	for _, v := range g.data {
		if v {
			count += 1
		}
	}
	return
}

func (g *BoolGrid) getIndex(x, y int) int {
	return y*g.width + x
}
