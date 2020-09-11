package main

import (
	"github.com/fogleman/gg"
	Automata "github.com/jtheiss19/Simple-Cell-Automata"
)

func GroundSimulation(c *Automata.Cell) {

	c.Value =
		c.Value +
			//Conduction Heat transfer
			3*conductionTransferRate*(c.Grid.GetCell(c.XPos, c.YPos-1).Value-c.Grid.GetCell(c.XPos, c.YPos).Value) + //Below Cell
			3*conductionTransferRate*(c.Grid.GetCell(c.XPos, c.YPos+1).Value-c.Grid.GetCell(c.XPos, c.YPos).Value) + //Above Cell
			3*conductionTransferRate*(c.Grid.GetCell(c.XPos-1, c.YPos).Value-c.Grid.GetCell(c.XPos, c.YPos).Value) + //Left Cell
			3*conductionTransferRate*(c.Grid.GetCell(c.XPos+1, c.YPos).Value-c.Grid.GetCell(c.XPos, c.YPos).Value) //Right Cell

}

func GroundDraw(dc *gg.Context, cell *Automata.Cell) {
	otherValue := (1 - cell.Value)
	if otherValue < 0.4 {
		otherValue = 0.4
	}
	dc.SetRGB(1, otherValue, otherValue)
}
