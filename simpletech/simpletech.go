package simpletech

import (
	"github.com/slitchfield/sudoku_solver/board"
	//"github.com/slitchfield/sudoku_solver/logging"
)

func NaiveElimSub(b *board.Board) {

	for row_index, row := range b.Board {

		for col_index, cell := range row {

			// If this is already a definite cell, move on
			if cell.PossCount() == 1 {
				continue
			}

			// Get the boundaries for this subgrid
			row_min := row_index / 3 * 3 // Note: this relies on integer truncation
			row_max := row_min + 2

			col_min := col_index / 3 * 3
			col_max := col_min + 2

			// For the row coords in this sub grid...
			for sub_row := row_min; sub_row <= row_max; sub_row++ {

				// For the col coords in this sub grid...
				for sub_col := col_min; sub_col <= col_max; sub_col++ {

					// Skip us!
					if sub_row == row_index && sub_col == col_index {
						continue
					}

					// Examine every other cell
					remote_cell := b.Board[sub_row][sub_col]

					// If this other cell only has one possibility, remove it from us
					if remote_cell.PossCount() == 1 {
						cell.RemovePoss(remote_cell.GetDefiniteVal())
					}

				} // End sub_col

			} // End sub_row

		} // end col_index

	} // end row_index

}

func NaiveElimRow(b *board.Board) {

	for row_index, row := range b.Board {

		for col_index, cell := range row {

			// If this is already a definite cell, move on
			if cell.PossCount() == 1 {
				continue
			}

			for remote_col := 0; remote_col < 9; remote_col++ {

				// If this is us, move on
				if remote_col == col_index {
					continue
				}

				remote_cell := b.Board[row_index][remote_col]

				// If this is a definite cell, eliminate it's possibility from us
				if remote_cell.PossCount() == 1 {
					cell.RemovePoss(remote_cell.GetDefiniteVal())
				}

			}

		}

	}

}

func NaiveElimCol(b *board.Board) {

	for row_index, row := range b.Board {

		for col_index, cell := range row {

			// If this is already a definite cell, move on
			if cell.PossCount() == 1 {
				continue
			}

			for remote_row := 0; remote_row < 9; remote_row++ {

				// If this is us, move on
				if remote_row == row_index {
					continue
				}

				remote_cell := b.Board[remote_row][col_index]

				// If this is a definite cell, eliminate it's possibility from us
				if remote_cell.PossCount() == 1 {
					cell.RemovePoss(remote_cell.GetDefiniteVal())
				}

			}

		}

	}
}

func ResolvePoss(b *board.Board) {
	NaiveElimSub(b)
	NaiveElimRow(b)
	NaiveElimCol(b)
}

func ElimLastRemainingSub(b *board.Board) {

	for row_index, row := range b.Board {

		for col_index, cell := range row {

			// If this is already a definite cell, move on
			if cell.PossCount() == 1 {
				continue
			}

			// Get the boundaries for this subgrid
			row_min := row_index / 3 * 3 // Note: this relies on integer truncation
			row_max := row_min + 2

			col_min := col_index / 3 * 3
			col_max := col_min + 2

			// For everything I can be, am I the only one that can be that?
		PossLoop:
			for pos_i, test := range cell.Poss {

				if !test {
					continue
				} // No way we can be this number!

				// For the row coords in this sub grid...
				for sub_row := row_min; sub_row <= row_max; sub_row++ {

					// For the col coords in this sub grid...
					for sub_col := col_min; sub_col <= col_max; sub_col++ {

						// Skip ourselves
						if (sub_row == row_index) && (sub_col == col_index) {
							continue
						}

						// Examine every other cell
						remote_cell := b.Board[sub_row][sub_col]

						// If this other cell can also be this number, we
						// don't make any conclusions, and move on
						if remote_cell.Contains(pos_i + 1) {
							continue PossLoop
						}

					} // End sub_col

				} // End sub_row

				// If we're here, we're the only cell in this block that can be this possibility
				cell.SetDefiniteVal(pos_i + 1)
				break

			} // End poss iteration

		} // end col_index

	} // end row_index

}

func ElimLastRemainingRow(b *board.Board) {

	for row_index, row := range b.Board {

		for col_index, cell := range row {

			// If this is already a definite cell, move on
			if cell.PossCount() == 1 {
				continue
			}

			// For everything I can be, am I the only one in this row that can be that?
		PossLoop:
			for pos_i, test := range cell.Poss {

				if !test {
					continue
				} // No way we can be this number!

				// For the columns in this row...
				for sub_col := 0; sub_col < 9; sub_col++ {

					// Skip ourselves
					if sub_col == col_index {
						continue
					}

					// Examine every other cell
					remote_cell := b.Board[row_index][sub_col]

					// If this other cell can also be this number, we
					// don't make any conclusions, and move on
					if remote_cell.Contains(pos_i + 1) {
						continue PossLoop
					}

				} // End sub_col

				// If we're here, we're the only cell in this block that can be this possibility
				cell.SetDefiniteVal(pos_i + 1)
				break

			} // End poss iteration

		} // end col_index

	} // end row_index

}

func ElimLastRemainingCol(b *board.Board) {

	for row_index, row := range b.Board {

		for col_index, cell := range row {

			// If this is already a definite cell, move on
			if cell.PossCount() == 1 {
				continue
			}

			// For everything I can be, am I the only one in this row that can be that?
		PossLoop:
			for pos_i, test := range cell.Poss {

				if !test {
					continue
				} // No way we can be this number!

				// For the rows in this column...
				for sub_row := 0; sub_row < 9; sub_row++ {

					// Skip ourselves
					if sub_row == row_index {
						continue
					}

					// Examine every other cell
					remote_cell := b.Board[sub_row][col_index]

					// If this other cell can also be this number, we
					// don't make any conclusions, and move on
					if remote_cell.Contains(pos_i + 1) {
						continue PossLoop
					}

				} // End sub_col

				// If we're here, we're the only cell in this column that can be this possibility
				cell.SetDefiniteVal(pos_i + 1)
				break

			} // End poss iteration

		} // end col_index

	} // end row_index

}
