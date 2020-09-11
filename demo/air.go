package main

import (
	"github.com/fogleman/gg"
	Automata "github.com/jtheiss19/Simple-Cell-Automata"
)

func AirSimulation(c *Automata.Cell) {
	//PHYSICS GOES HERE
	c.Value =
		c.Value +

			//Conduction Heat transfer
			conductionTransferRate*(c.Grid.GetCell(c.XPos, c.YPos-1).Value-c.Grid.GetCell(c.XPos, c.YPos).Value) + //Below Cell
			conductionTransferRate*(c.Grid.GetCell(c.XPos, c.YPos+1).Value-c.Grid.GetCell(c.XPos, c.YPos).Value) + //Above Cell
			conductionTransferRate*(c.Grid.GetCell(c.XPos-1, c.YPos).Value-c.Grid.GetCell(c.XPos, c.YPos).Value) + //Left Cell
			conductionTransferRate*(c.Grid.GetCell(c.XPos+1, c.YPos).Value-c.Grid.GetCell(c.XPos, c.YPos).Value) + //Right Cell

			//Convection Heat Transfer
			convectionTransferRate*c.Grid.GetCell(c.XPos, c.YPos-1).Value -
			convectionTransferRate*c.Grid.GetCell(c.XPos, c.YPos).Value

}

func AirDraw(dc *gg.Context, cell *Automata.Cell) {
	dc.SetRGB(cell.Value, 0, 0)
}
