package main

import (
	"fmt"
	"strconv"

	"github.com/fogleman/gg"
)

// This package defines all the back ground static code. Think of this section as the simulation engine

//Cell remain mostly unchanged, they carry the value name which can carry type such as "air" or "wood" for tree simulation
type cell struct {
	name  string
	xPos  int
	yPos  int
	value float64
}

//NewGrid is the constructor function for the grid object
func NewGrid(xSize, ySize int) grid {
	myNewGrid := grid{colSize: xSize, rowSize: ySize, simFunction: defaultSimulation}
	for row := 1; row <= ySize; row++ {
		for col := 1; col <= xSize; col++ {
			myNewCell := cell{name: "air", xPos: col, yPos: row, value: 0}
			myNewGrid.cellList = append(myNewGrid.cellList, &myNewCell)
		}
	}
	return (myNewGrid)
}

//grid carrys a reference to and overidable function called simFunction. This is so the user can define a function and pass it to the specific grid. That way they can run parrell simulations
type grid struct {
	colSize     int
	rowSize     int
	cellList    []*cell
	simFunction func(*grid)
}

//copyGrid preforms a deep copy of the grid to create new references in memory
func (g *grid) copyGrid() grid {
	tempGrid := *g
	var tempCellList []*cell
	for _, v := range g.cellList {
		newCell := *v
		tempCellList = append(tempCellList, &newCell)
	}
	tempGrid.cellList = tempCellList
	return (tempGrid)
}

//getCell is used to grab a specific cell
func (g *grid) getCell(x, y int) *cell {
	return (g.cellList[(y-1)*g.colSize+x-1])
}

//prettyPrint prints the grid to terminal, if it isn't pretty your grid is wider then your terminal
func (g *grid) prettyPrint() {
	var stringArray []string
	for row := 1; row <= g.rowSize; row++ {
		tempString := ""
		for col := 1; col <= g.colSize; col++ {
			tempString = tempString + " " + strconv.FormatFloat(g.getCell(col, row).value, 'f', 3, 64)
		}
		stringArray = append(stringArray, tempString)
	}

	stringArrayLen := len(stringArray)
	for k := range stringArray {
		fmt.Println(stringArray[stringArrayLen-k-1])
	}
}

//printPNG prints a png of the current grid to the folder the program is ran in
func (g *grid) printPNG(fileName string) {
	dc := gg.NewContext(g.colSize, g.rowSize)

	for _, v := range g.cellList {
		dc.DrawRectangle(float64(g.colSize-v.xPos), float64(g.rowSize-v.yPos), 1, 1)
		dc.SetRGB(v.value, 0, 0)
		dc.Fill()
	}

	dc.SavePNG(fileName + ".png")
}

//simulate calls the grids simulate function, defined here explicitly for potential pre-flight checks
func (g *grid) simulate() {
	g.simFunction(g)
}

//defaultSimulation is a catch all simulation function so that the user can't accidentally crash the program.
func defaultSimulation(g *grid) {
	tempGrid := g.copyGrid()
	g.cellList = tempGrid.cellList
}
