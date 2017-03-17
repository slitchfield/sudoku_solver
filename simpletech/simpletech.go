
package simpletech

import (
    "github.com/slitchfield/sudoku_solver/board"
    "github.com/slitchfield/sudoku_solver/logging"
)

func NaiveElimSub(b *board.Board) {
    logging.Trace.Printf("Inside NaiveElimSub!")
    

    for row_index, row := range b.Board {
        
        for col_index, cell := range row {

            // If this is already a definite cell, move on
            if len(cell.Poss) == 1 {
                continue
            }

            // Get the boundaries for this subgrid
            row_min := row_index/3 * 3  // Note: this relies on integer truncation
            row_max := row_min + 2

            col_min := col_index/3 * 3
            col_max := col_min + 2

            // For the row coords in this sub grid...
            for sub_row := row_min; sub_row <= row_max; sub_row++ {

                // For the col coords in this sub grid...
                for sub_col := col_min; sub_col <= col_max; sub_col++ {
                    
                    // Examine every other cell
                    remote_cell := b.Board[sub_row][sub_col]

                    // If this other cell only has one possibility, remove it from us
                    if len(remote_cell.Poss) == 1 {
                        cell.RemovePoss(remote_cell.Poss[0])
                    }
                
                } // End sub_col

            } // End sub_row

        } // end col_index

    } // end row_index

}

func NaiveElimRow(b *board.Board) {

    logging.Trace.Printf("Inside NaiveElimRow!")

    for row_index, row := range b.Board {

        for col_index, cell := range row {

            // If this is already a definite cell, move on
            if len(cell.Poss) == 1 {
                continue
            }

            for remote_col := 0; remote_col < 9; remote_col++ {

                // If this is us, move on
                if remote_col == col_index { continue }

                remote_cell := b.Board[row_index][remote_col]

                // If this is a definite cell, eliminate it's possibility from us
                if len(remote_cell.Poss) == 1 {
                    cell.RemovePoss(remote_cell.Poss[0])
                }

            }

        }

    }

}

func NaiveElimCol(b *board.Board) {
        logging.Trace.Printf("Inside NaiveElimCol!")

    for row_index, row := range b.Board {

        for col_index, cell := range row {

            // If this is already a definite cell, move on
            if len(cell.Poss) == 1 {
                continue
            }

            for remote_row := 0; remote_row < 9; remote_row++ {

                // If this is us, move on
                if remote_row == row_index { continue }

                remote_cell := b.Board[remote_row][col_index]

                // If this is a definite cell, eliminate it's possibility from us
                if len(remote_cell.Poss) == 1 {
                    cell.RemovePoss(remote_cell.Poss[0])
                }

            }

        }

    }
}
