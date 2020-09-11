//Package automata is a simple package for adding a felxable cellular automata simulations to go!
package automata

import (
	"fmt"
	"strconv"

	"github.com/cheggaaa/pb"
	"github.com/fogleman/gg"
)

//NewGrid is the constructor function for the Grid object
func NewGrid(xSize, ySize int) *Grid {
	myNewGrid := &Grid{ColSize: xSize, RowSize: ySize}
	for row := 1; row <= ySize; row++ {
		for col := 1; col <= xSize; col++ {
			myNewCell := Cell{XPos: col, YPos: row, Value: 0, Grid: myNewGrid, SimFunc: defaultSimFunc, DrawFunc: defaultDrawFunc}
			myNewGrid.CellList = append(myNewGrid.CellList, &myNewCell)
		}
	}
	return (myNewGrid)
}

//Grid carrys a reference to and overidable function called simFunction. This is so the user can define a function and pass it to the specific Grid. That way they can run parrell simulations
type Grid struct {
	ColSize  int
	RowSize  int
	CellList []*Cell
}

func (g *Grid) SetupGrid(SetupFunc func(*Grid, int, int)) {
	for col := 1; col <= g.ColSize; col++ {
		for row := 1; row <= g.RowSize; row++ {
			if row != 1 && row != g.RowSize {
				if col != 1 && col != g.ColSize {
					SetupFunc(g, col, row)
				}
			}
		}
	}
}

//CopyGrid preforms a deep copy of the Grid to create new references in memory
func (g *Grid) CopyGrid() *Grid {
	tempGrid := *g
	var tempCellList []*Cell
	for _, v := range g.CellList {
		newCell := *v
		newCell.Grid = g
		tempCellList = append(tempCellList, &newCell)
	}
	tempGrid.CellList = tempCellList
	return (&tempGrid)
}

//GetCell is used to grab a specific cell
func (g *Grid) GetCell(x, y int) *Cell {
	return (g.CellList[(y-1)*g.ColSize+x-1])
}

//PrettyPrint prints the Grid to terminal, if it isn't pretty your Grid is wider then your terminal
func (g *Grid) PrettyPrint() {
	var stringArray []string
	for row := 1; row <= g.RowSize; row++ {
		tempString := ""
		for col := 1; col <= g.ColSize; col++ {
			tempString = tempString + " " + strconv.FormatFloat(g.GetCell(col, row).Value, 'f', 3, 64)
		}
		stringArray = append(stringArray, tempString)
	}

	stringArrayLen := len(stringArray)
	for k := range stringArray {
		fmt.Println(stringArray[stringArrayLen-k-1])
	}
}

//PrintPNG prints a png of the current Grid to the folder the program is ran in
func (g *Grid) PrintPNG(fileType string) {
	dc := gg.NewContext(g.ColSize, g.RowSize)

	for _, v := range g.CellList {
		dc.DrawRectangle(float64(g.ColSize-v.XPos), float64(g.RowSize-v.YPos), 1, 1)
		v.draw(dc)
		dc.Fill()
	}

	dc.SavePNG("./pictures/" + fileType + ".png")
}

//Simulate calls the Grids simulate function, defined here explicitly for potential pre-flight checks
func (g *Grid) Simulate() {

	tempGrid := g.CopyGrid()

	for row := 1; row <= tempGrid.RowSize; row++ {
		for col := 1; col <= tempGrid.ColSize; col++ {
			//ignore the edges
			if row != 1 && row != tempGrid.RowSize {
				if col != 1 && col != tempGrid.ColSize {
					tempGrid.GetCell(col, row).simulate()
				}
			}
		}
	}
	g.CellList = tempGrid.CellList
}

//RunSimulation calls the Grids simulate function multiple times
func (g *Grid) RunSimulation(steps int, interupt func(*Grid, int)) {
	//Used for the progressbar (pb)
	bar := pb.StartNew(steps)

	//Simulation Loop
	for i := 0; i < steps; i++ {
		g.Simulate()

		interupt(g, i)

		bar.Increment()
	}

	bar.Finish()

}

//defaultSimulation is a catch all simulation function so that the user can't accidentally crash the program.
func defaultSimulation(g *Grid) {
	tempGrid := g.CopyGrid()
	g.CellList = tempGrid.CellList
}
