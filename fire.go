package main

import "github.com/cheggaaa/pb"

var GridSize int = 400

func main() {
	//initialize the grid
	testGrid := NewGrid(GridSize, GridSize)

	//Set the initial values
	for col := 1; col <= testGrid.colSize; col++ {
		testGrid.getCell(col, 1).value = 1
	}

	//Set your simulation function of kind func(*grid)
	testGrid.simFunction = mySimulationFunction

	//run the sim function
	timesteps := 700

	//Used for the progressbar (pb)
	bar := pb.StartNew(timesteps)
	for i := 0; i < timesteps; i++ {
		bar.Increment()
		testGrid.simulate()
	}

	bar.Finish()
	testGrid.printPNG("FinalPicture")
}

func mySimulationFunction(g *grid) {
	//copy the grid so sim timesteps don't overlap
	tempGrid := g.copyGrid()

	//We are referencing one cell, inside is where the sim code goes
	for row := 1; row <= g.rowSize; row++ {
		for col := 1; col <= g.colSize; col++ {

			//PHYSICS GOES HERE: This sim code ignores the bottom level (perhaps a burning branch or something) and top level (crashes) and then shifts heat around the air
			if row != 1 && row != g.rowSize {
				tempGrid.getCell(col, row).value = g.getCell(col, row).value + 0.3*(g.getCell(col, row-1).value-g.getCell(col, row).value) + 0.3*(g.getCell(col, row+1).value-g.getCell(col, row).value)
			}

		}
	}

	//Setting our newly calculated values as the true ones
	g.cellList = tempGrid.cellList
}
