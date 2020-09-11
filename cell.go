package automata

import (
	"github.com/fogleman/gg"
)

//Cell remain mostly unchanged, they carry the Value Type which can carry type such as "air" or "wood" for tree simulation
type Cell struct {
	Type     string
	XPos     int
	YPos     int
	Value    float64
	Grid     *Grid
	SimFunc  func(*Cell)
	DrawFunc func(*gg.Context, *Cell)
}

func (c *Cell) simulate() {
	c.SimFunc(c)
}

func (c *Cell) draw(dc *gg.Context) {
	c.DrawFunc(dc, c)
}

func defaultSimFunc(cell *Cell) {
}

func defaultDrawFunc(dc *gg.Context, cell *Cell) {
	dc.SetRGB(cell.Value, 0, 0)
}
