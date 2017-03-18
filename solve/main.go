package main

import "github.com/slitchfield/sudoku_solver/board"
import "github.com/slitchfield/sudoku_solver/logging"
import "github.com/slitchfield/sudoku_solver/simpletech"
import "os"
import "fmt"
import "github.com/fatih/color"
import "time"

func solve(b *board.Board) {

  start_time := time.Now()

  for ;; {
    simpletech.ResolvePoss(b)
    simpletech.ElimLastRemainingSub(b)
    simpletech.ResolvePoss(b)
    simpletech.ElimLastRemainingRow(b)
    simpletech.ResolvePoss(b)
    simpletech.ElimLastRemainingCol(b)

		if b.CheckIfSolved() {
			success := color.New(color.BgGreen).Add(color.FgRed).PrintfFunc()
			success("Solved it!")
			fmt.Printf("\n")
			b.Print()
			break
		}

    if time.Since(start_time).Seconds() > 1 {
      failure := color.New(color.BgRed).Add(color.FgGreen).PrintfFunc()
      failure("Could not solve it :(")
      fmt.Printf("\n")
      b.Print()
      break
    }
  }

}

func main() {

	logging.Log_Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	filename := "/home/samuel/go/src/github.com/slitchfield/sudoku_solver/test_boards/1.board"

	logging.Trace.Printf("Reading in board from %s", filename)
	myBoard := board.BoardFromCSV(filename)

	logging.Trace.Printf("Beginning solving!")
	myBoard.Print()

  solve(myBoard)

}
