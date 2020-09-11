package main

import (
	"fmt"
	"os"
	"strconv"

	Automata "github.com/jtheiss19/Simple-Cell-Automata"
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

	//Set your CellTypes
	for col := 1; col <= testGrid.ColSize; col++ {
		for row := 1; row <= testGrid.RowSize; row++ {
			if col%50 == 0 {
				testGrid.GetCell(col, row).SimFunc = TreeSimulation
				testGrid.GetCell(col, row).DrawFunc = TreeDraw
			} else {
				testGrid.GetCell(col, row).SimFunc = AirSimulation
				testGrid.GetCell(col, row).DrawFunc = AirDraw
			}
		}
	}

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
