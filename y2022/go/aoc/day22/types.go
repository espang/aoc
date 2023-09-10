package day22

type Grid [][]Cell

func (g Grid) At(row, col int) Cell {
	if row < 0 || row >= len(g) {
		return Void
	}
	theRow := g[row]
	if col < 0 || col >= len(theRow) {
		return Void
	}
	return theRow[col]
}

func (g Grid) IsVoid(row, col int) bool {
	return g.At(row, col) == Void
}

func (g Grid) MoveInCol(forward bool, row, col int) (int, int) {
	newRow := row
	if forward {
		newRow++
	} else {
		newRow--
	}
	if newRow >= len(g) {
		for r := 0; r < len(g); r++ {
			if !g.IsVoid(r, col) {
				return r, col
			}
		}
	}
	if newRow < 0 {
		for r := len(g) - 1; r >= 0; r-- {
			if !g.IsVoid(r, col) {
				return r, col
			}
		}
	}
	return newRow, col
}

func (g Grid) MoveInRow(forward bool, row, col int) (int, int) {
	newCol := col
	if forward {
		newCol++
	} else {
		newCol--
	}
	if newCol >= len(g[row]) {
		for c := 0; c < len(g[row]); c++ {
			if !g.IsVoid(row, c) {
				return row, c
			}
		}
	}
	if newCol < 0 {
		for c := len(g[row]) - 1; c >= 0; c-- {
			if !g.IsVoid(row, c) {
				return row, c
			}
		}
	}
	return row, newCol
}

func (g Grid) Handle(m Move, row, col int, orientation Orientation) (int, int) {
	var move func(int, int) (int, int)
	switch orientation {
	case Left:
		move = func(r, c int) (int, int) { return g.MoveInRow(false, r, c) }
	case Right:
		move = func(r, c int) (int, int) { return g.MoveInRow(true, r, c) }
	case Up:
		move = func(r, c int) (int, int) { return g.MoveInCol(false, r, c) }
	case Down:
		move = func(r, c int) (int, int) { return g.MoveInCol(true, r, c) }
	}

	for i := 0; i < int(m); i++ {
		nextRow, nextCol := move(row, col)
		for g.At(nextRow, nextCol) == Void {
			nextRow, nextCol = move(nextRow, nextCol)
		}
		switch g.At(nextRow, nextCol) {
		case Void:
			panic("unreachable void")
		case Open:
			row, col = nextRow, nextCol
		case Wall:
			return row, col
		}
	}
	return row, col
}

func (g Grid) MoveOnCube(row, col int, orientation Orientation, jumps map[State]State) (int, int, Orientation) {
	nextRow, nextCol := row, col
	switch orientation {
	case Left:
		nextCol--
	case Right:
		nextCol++
	case Up:
		nextRow--
	case Down:
		nextRow++
	}
	nextOrientation := orientation
	if g.IsVoid(nextRow, nextCol) {
		jumpTo := jumps[s(nextRow, nextCol, orientation)]
		nextRow = jumpTo.row
		nextCol = jumpTo.col
		nextOrientation = jumpTo.orientation
	}
	switch g.At(nextRow, nextCol) {
	case Void:
		panic("unreachable void")
	case Open:
		return nextRow, nextCol, nextOrientation
	case Wall:
		return row, col, orientation
	}
	panic("unreachable code")
}

type Cube struct {
	Top    [][]Cell
	Bottom [][]Cell
	Front  [][]Cell
	Back   [][]Cell
	Left   [][]Cell
	Right  [][]Cell
}

func SideFromBoard(board Grid, rowSection, colSection, size int) [][]Cell {
	rowOffset := rowSection * size
	colOffset := colSection * size
	matrix := [][]Cell{}
	for rowPosition := 0; rowPosition < size; rowPosition++ {
		aRow := []Cell{}
		row := rowOffset + rowPosition
		for colPosition := 0; colPosition < size; colPosition++ {
			col := colOffset + colPosition
			if board.IsVoid(row, col) {
				panic("you did it wrong!")
			}
			aRow = append(aRow, board.At(row, col))
		}
		matrix = append(matrix, aRow)
	}
	return matrix
}
