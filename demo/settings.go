package main

import "os"

//Settings
var (
	timeSteps       = 1000
	gridSize        = 400
	interval        = 500
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
