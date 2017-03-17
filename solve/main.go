package main

import "github.com/slitchfield/sudoku_solver/board"
import "io"
import "log"
import "os"

var (
    Trace   *log.Logger
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
)

func Init(
    traceHandle io.Writer,
    infoHandle io.Writer,
    warningHandle io.Writer,
    errorHandle io.Writer) {

    Trace = log.New(traceHandle,
            "TRACE: ",
            log.Ldate|log.Ltime|log.Lshortfile)

    Info = log.New(infoHandle,
            "INFO: ",
            log.Ldate|log.Ltime|log.Lshortfile)

    Warning = log.New(warningHandle,
            "Warning: ",
            log.Ldate|log.Ltime|log.Lshortfile)

    Error = log.New(errorHandle,
        "Error: ",
        log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

    Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

    myBoard := new(board.Board)
    myBoard.Testing = true
    Trace.Printf("Hello World, with a board!! Testing: %v\n", myBoard.Testing)
}
