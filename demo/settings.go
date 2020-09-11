package main

import (
	"fmt"
	"os"
	"strconv"

	Automata "github.com/jtheiss19/Simple-Cell-Automata"
)

//Settings
var (
	timeSteps       = 20000
	gridSize        = 400
	interval        = 200
	slope           = 0
	collectData     = false
	collectPictures = true
)

//Physics Variables
var (
	conductionTransferRate = 0.03
	convectionTransferRate = 0.01
)

//Global Variables
var (
	dataFile *os.File
)

func setup(g *Automata.Grid, col int, row int) {
	if col%50 == 0 {
		g.GetCell(col, row).SimFunc = TreeSimulation
		g.GetCell(col, row).DrawFunc = TreeDraw
	} else {
		g.GetCell(col, row).SimFunc = AirSimulation
		g.GetCell(col, row).DrawFunc = AirDraw
	}

	if row <= g.RowSize/2 {
		g.GetCell(col, row).SimFunc = GroundSimulation
		g.GetCell(col, row).DrawFunc = GroundDraw
	}
}

func myRunFunction(grid *Automata.Grid, i int) {
	if i%interval == 0 {

		//Should we collect data
		if collectData {
			//Calculate Data
			myValue := 0.0
			for _, testCell := range grid.CellList {
				myValue += testCell.Value
			}

			//Print data
			_, err := dataFile.WriteString(fmt.Sprintf("Total Heat: %.2f On Trail:"+strconv.Itoa(i), myValue))
			if err != nil {
				panic(err.Error())
			}
			dataFile.WriteString("\n")
		}

		//Should we make a picture
		if collectPictures {
			//Create Image
			grid.PrintPNG("Frame_" + strconv.Itoa(i))
		}
	}
}
