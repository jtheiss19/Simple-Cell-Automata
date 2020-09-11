package main

import (
	"github.com/fogleman/gg"
	Automata "github.com/jtheiss19/Simple-Cell-Automata"
)

func TreeSimulation(c *Automata.Cell) {
	if c.Value < 0.5 {
		c.Value =
			c.Value +
				//Conduction Heat transfer
				conductionTransferRate*(c.Grid.GetCell(c.XPos, c.YPos-1).Value-c.Grid.GetCell(c.XPos, c.YPos).Value) + //Below Cell
				conductionTransferRate*(c.Grid.GetCell(c.XPos, c.YPos+1).Value-c.Grid.GetCell(c.XPos, c.YPos).Value) + //Above Cell
				conductionTransferRate*(c.Grid.GetCell(c.XPos-1, c.YPos).Value-c.Grid.GetCell(c.XPos, c.YPos).Value) + //Left Cell
				conductionTransferRate*(c.Grid.GetCell(c.XPos+1, c.YPos).Value-c.Grid.GetCell(c.XPos, c.YPos).Value) //Right Cell
	} else {
		c.Value = 1

		c.SimFunc = BurningTreeSim
		c.DrawFunc = BurningTreeDraw
	}
}

func TreeDraw(dc *gg.Context, cell *Automata.Cell) {
	dc.SetRGB255(150, 75, 0)
}

func BurningTreeSim(c *Automata.Cell) {
	c.Value = 1
}

func BurningTreeDraw(dc *gg.Context, cell *Automata.Cell) {
	dc.SetRGB(cell.Value, 0, 0)
}
