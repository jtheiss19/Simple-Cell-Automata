package main

import (
	"os"

	Automata "github.com/jtheiss19/Simple-Cell-Automata"
)

func main() {
	//initialize the grid
	testGrid := Automata.NewGrid(gridSize, gridSize)

	//Set the initial Values Manually
	for col := 1; col <= testGrid.ColSize; col++ {
		testGrid.GetCell(col, 250).Value = 1
	}

	//OR set your cells using a function your CellTypes
	testGrid.SetupGrid(setup)

	testGrid.PrintPNG("Setup")

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
