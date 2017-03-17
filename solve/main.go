package main

import "github.com/slitchfield/sudoku_solver/board"
import "github.com/slitchfield/sudoku_solver/logging"
import "github.com/slitchfield/sudoku_solver/simpletech"
import "os"

func main() {

    logging.Log_Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

    filename := "/home/samuel/go/src/github.com/slitchfield/sudoku_solver/test_boards/1.board"

    logging.Trace.Printf("Reading in board from %s", filename)
    myBoard := board.BoardFromCSV(filename)
    myBoard.Solved = false
    myBoard.Dimension = 9

    logging.Trace.Printf("Beginning solving!")
    myBoard.Print()

    logging.Trace.Printf("Eliminating possibilities in subgrid!")
    simpletech.NaiveElimSub(myBoard)
    myBoard.Print()

    logging.Trace.Printf("Eliminating poss in rows!")
    simpletech.NaiveElimRow(myBoard)
    myBoard.Print()

    logging.Trace.Printf("Eliminating poss in cols!")
    simpletech.NaiveElimCol(myBoard)
    myBoard.Print()
}
