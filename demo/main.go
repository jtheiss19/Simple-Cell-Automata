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
	collectData     = true
	collectPictures = true
)

//Physics Variables
var (
	conductionTransferRate = 0.01
	convectionTransferRate = 0.01
)

//Global Variables

var (
	dataFile    *os.File
	currentStep = 0
)

func main() {
	//initialize the grid
	testGrid := Automata.NewGrid(gridSize, gridSize)

	/*
		//Set the initial Values
		for col := 150; col <= 250; col++ {
			for row := 150; row <= 250; row++ {
				testGrid.GetCell(col, row).Value = 1
			}
		}
	*/

	testGrid.GetCell(gridSize/2, gridSize/2).Value = 1

	testGrid.PrintPNG("Setup")

	//Set your simulation function of kind func(*grid)
	testGrid.SimFunction = mySimulationFunction

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

func mySimulationFunction(g *Automata.Grid) {
	//copy the grid so sim timesteps don't overlap
	tempGrid := g.CopyGrid()

	//We are referencing one cell, inside is where the sim code goes
	for row := 1; row <= g.RowSize; row++ {
		for col := 1; col <= g.ColSize; col++ {

			//PHYSICS GOES HERE: This sim code ignores the bottom level (perhaps a burning branch or something) and top level (crashes) and then shifts heat around the air
			if row != 1 && row != g.RowSize {
				if col != 1 && col != g.ColSize {
					tempGrid.GetCell(col, row).Value =
						g.GetCell(col, row).Value +
							//Conduction Heat transfer
							conductionTransferRate*(g.GetCell(col, row-1).Value-g.GetCell(col, row).Value) + //Below Cell
							conductionTransferRate*(g.GetCell(col, row+1).Value-g.GetCell(col, row).Value) + //Above Cell
							conductionTransferRate*(g.GetCell(col-1, row).Value-g.GetCell(col, row).Value) + //Left Cell
							conductionTransferRate*(g.GetCell(col+1, row).Value-g.GetCell(col, row).Value) + //Right Cell
							//Convection Heat Transfer
							g.GetCell(col, row-1).Value*convectionTransferRate -
							g.GetCell(col, row).Value*convectionTransferRate
				}
			}
		}
	}
	if currentStep <= 4000 {
		tempGrid.GetCell(200, 200).Value = 1
		currentStep++
	}
	//Setting our newly calculated Values as the true ones
	g.CellList = tempGrid.CellList
}

func myRunFunction(grid *Automata.Grid, i int) {
	if i%interval == 0 {

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

		if collectPictures {
			//Create Image
			grid.PrintPNG("Frame_" + strconv.Itoa(i))
		}
	}

}
