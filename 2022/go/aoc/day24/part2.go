package day24

import (
	"aoc/aoc"
	"fmt"
)

func minimumEscape2(start, end Position, width, height int, tornados aoc.Set[Tornado]) int {
	if start == end {
		return 0
	}
	currentPositions := aoc.Set[Position]{}
	currentPositions.Add(start)
	currentTornados := tornados

	target := end
	reachedBefore := false

outer:
	for turn := 0; ; turn++ {
		if len(currentPositions) == 0 {
			return -1
		}
		currentTornados = moveTornados(width, height, currentTornados)
		occupied := occupiedPositions(currentTornados)
		nextPositions := aoc.Set[Position]{}
		for pos := range currentPositions {
			moves := potentialPositions(pos, start, end, width, height)
			for _, move := range moves {
				if !occupied.IsMember(move) {
					if move == target {
						if move == end && reachedBefore {
							return turn + 1
						}
						if move == end {
							// the first time we reach the end. This will
							// be a potential fastest path because we could
							// wait any number of turns and therefore be at
							// least as fast as any path that reaches the end
							// later.
							target = start
							reachedBefore = true
							currentPositions = aoc.Set[Position]{}
							currentPositions.Add(end)
							continue outer
						}
						if move == start {
							target = end
							currentPositions = aoc.Set[Position]{}
							currentPositions.Add(start)
							continue outer
						}
						panic("unreachable")
					}
					nextPositions.Add(move)
				}
			}
		}
		currentPositions = nextPositions
	}
}

func Part2(input string) {
	board := Parse(input)
	start := Position{row: 0}
	end := Position{row: len(board) - 1}
	for col, c := range board[0] {
		if c == '.' {
			start.col = col
			break
		}
	}
	for col, c := range board[len(board)-1] {
		if c == '.' {
			end.col = col
			break
		}
	}
	tornados := aoc.Set[Tornado]{}
	for rowIdx, row := range board {
		for colIdx, cell := range row {
			switch cell {
			case '>', '<', '^', 'v':
				tornados.Add(Tornado{
					pos:       Position{row: rowIdx, col: colIdx},
					direction: cell,
				})
			}
		}
	}
	width := len(board[0])
	height := len(board)
	fmt.Println(minimumEscape2(start, end, width, height, tornados))
}
