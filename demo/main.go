package main

import (
	"fmt"
	"os"
	"strconv"

	Automata "github.com/jtheiss19/Simple-Cell-Automata"
)

//Settings
var (
	timeSteps       = 10000
	gridSize        = 400
	interval        = 200
	collectData     = false
	collectPictures = false
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

func main() {
	//initialize the grid
	testGrid := Automata.NewGrid(gridSize, gridSize)

	//Set the initial Values
	for col := 150; col <= 250; col++ {
		for row := 150; row <= 250; row++ {
			testGrid.GetCell(col, row).Value = 1
		}
	}

	testGrid.PrintPNG("Setup")

	//Set your simulation function of kind func(*Cell)
	for _, c := range testGrid.CellList {
		c.SimFunc = mySimulationFunction
	}

	if collectData {
		var err error
		dataFile, err = os.OpenFile("./data/mydata.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 77777)
		if err != nil {
			panic(err.Error())
		}
	}

	testGrid.RunSimulation(timeSteps, myRunFunction)

	testGrid.PrintPNG("FinalPicture")
}

func mySimulationFunction(c *Automata.Cell) {
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
