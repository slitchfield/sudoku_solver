
package board

import (
    "os"
    "encoding/csv"
    "fmt"
    "strconv"
    "github.com/slitchfield/sudoku_solver/logging"
    "github.com/fatih/color"
)

type Cell struct {
    Poss []bool
}

type Board struct {
    Solved bool
    Dimension int
    Board [][]Cell
}

func (c *Cell) Contains(i int) bool {
    return c.Poss[i-1]
}

func (c *Cell) PossCount() int {
    
    return_val := 0

    for _, test := range c.Poss {
        if test { return_val += 1 }
    }

    return return_val
}

func (c *Cell) OnlyContains(i int) bool {

    if (c.PossCount() == 1) && c.Contains(i) {
        return true
    }
    
    return false
}

func (c *Cell) GetDefiniteVal() int {

    if c.PossCount() != 1 {
        logging.Error.Printf("Tried to call GetDefiniteVal on cell with no definite val!")
        return 0
    }

    for i, test := range c.Poss {
        if test {
            return i + 1
        }        
    }

    logging.Error.Printf("Found no definite value!")
    return 0
}

func (c *Cell) SetDefiniteVal(i int) {

    for idx, test := range c.Poss {
        if idx == i-1 {
            if !test {
                logging.Error.Printf("Tried to set definite value to something we can't be!")
            }
        } else {
            c.Poss[idx] = false
        }
    }

}

func (c *Cell) RemovePoss(i int) {

    c.Poss[i-1] = false

}

func (b *Board) Print() {

    header  := "╭─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────╮\n"
    between := "├─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┤\n"
    footer  := "╰─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────╯\n"

    fmt.Print(header)

    // Iterate over all the rows
    for row_index, row := range b.Board {

        // We'll have 3 rows of text per cell
        for i := 0; i < 3; i++ {
            fmt.Print("│")  // Start the line off right

            // Iterate over all the cells in that row
            for _, cell := range row {

                if cell.PossCount() != 1 {
                    if cell.Contains(i*3+1) {
                        fmt.Printf("%d ", i*3+1)
                    } else {
                        fmt.Printf("  ")
                    }

                    if cell.Contains(i*3+2) {
                        fmt.Printf("%d ", i*3+2)
                    } else {
                        fmt.Printf("  ")
                    }
      
                    if cell.Contains(i*3+3) {
                        fmt.Printf("%d│", i*3+3)
                    } else {
                        fmt.Printf(" │")
                    }
                } else {
                    bg := color.New(color.BgGreen).Add(color.FgRed).PrintfFunc()

                    if cell.Contains(i*3+1) {
                        bg("%d ", i*3+1)
                    } else {
                        bg("  ")
                    }

                    if cell.Contains(i*3+2) {
                        bg("%d ", i*3+2)
                    } else {
                        bg("  ")
                    }
      
                    if cell.Contains(i*3+3) {
                        bg("%d", i*3+3)
                    } else {
                        bg(" ")
                    }
                    fmt.Printf("│")
                }
            }
            fmt.Print("\n")

        }

        if row_index < 8 {
            fmt.Print(between)
        }
    }

    fmt.Print(footer)
}

func (b *Board) CheckIfSolved() bool {

    for _, row := range b.Board {
        for _, cell := range row {
            if cell.PossCount() != 1 {
                return false
            }
        }
    }
    
    return true
}

func BoardFromCSV(filename string) *Board {

    // Allocate the board to be created
    return_board := new(Board)
    
    // Try to open the file and read it in if we can
    f, err := os.Open(filename)
    if err != nil {
        logging.Error.Printf("Could not open board with filename: %s", filename)
        panic(err)
    }
    defer f.Close()

    csv_reader := csv.NewReader(f)

    records, err := csv_reader.ReadAll()
    if err != nil {
        logging.Error.Printf("Could not read from file. Is it really csv?")
        panic(err)
    }

    // Check to see we've got the right dimensions!
    if (len(records) == 9) && (len(records[0]) == 9) {
        // Allocate the Cells
        return_board.Board = make([][]Cell, 9)
        for i := 0; i < 9; i++ {
            return_board.Board[i] = make([]Cell, 9)
        }

        // Give the cells values as appropriate
        for row_index, row := range records {
            for col_index, col := range row {

                // Try to convert the value into an integer
                val, err := strconv.Atoi(col)
                if err != nil {
                    logging.Error.Printf("Could not convert found number to int!")
                    panic(err)
                }

                // If there's actually a number in the board, it's the only possibility
                // Otherwise, anything is possible
                if val != 0 {
                    // Make the Possibility slice, and set only val true
                    return_board.Board[row_index][col_index].Poss = make([]bool, 9) 
                    return_board.Board[row_index][col_index].Poss[val-1] = true
                } else {
                    // Make the Possibility slice, and set all values possible
                    return_board.Board[row_index][col_index].Poss = make([]bool, 9)
                    for i, _ := range return_board.Board[row_index][col_index].Poss {
                        return_board.Board[row_index][col_index].Poss[i] = true
                    }
                }
            }
        }

    } else {
        logging.Error.Printf("Appears this csv has the wrong dimensions!")
    }
    
    return return_board
}
