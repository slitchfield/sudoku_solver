package main

import "github.com/slitchfield/sudoku_solver/board"
import "github.com/slitchfield/sudoku_solver/logging"
import "github.com/slitchfield/sudoku_solver/simpletech"
import "os"
import "fmt"
import "github.com/fatih/color"

func main() {

	logging.Log_Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	filename := "/home/samuel/go/src/github.com/slitchfield/sudoku_solver/test_boards/1.board"

	logging.Trace.Printf("Reading in board from %s", filename)
	myBoard := board.BoardFromCSV(filename)
	myBoard.Solved = false
	myBoard.Dimension = 9

	logging.Trace.Printf("Beginning solving!")
	myBoard.Print()

	for i := 0; i < 5; i++ {
		simpletech.ResolvePoss(myBoard)

		simpletech.ElimLastRemainingSub(myBoard)

		simpletech.ResolvePoss(myBoard)

		simpletech.ElimLastRemainingRow(myBoard)

		simpletech.ResolvePoss(myBoard)

		simpletech.ElimLastRemainingCol(myBoard)

		if myBoard.CheckIfSolved() {
			success := color.New(color.BgGreen).Add(color.FgRed).PrintfFunc()
			success("Solved it!")
			fmt.Printf("\n")
			myBoard.Print()
			break
		}
	}
}
