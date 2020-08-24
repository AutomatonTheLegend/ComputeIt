package Lab1

import "math/rand"

type Automaton struct {
	Cells        [][]*Cell
	BackCells    [][]*Cell
	History      [][][]*Cell
	Size         *Vector2
	Depth        int
	StatesCount  int
	Rule         map[[3]int]int
	HistoryIndex int
}

func NewAutomaton(randomGenerator *rand.Rand, size *Vector2, depth, statesCount int) *Automaton {
	automaton := new(Automaton)
	automaton.Size = size
	automaton.HistoryIndex = 0
	automaton.StatesCount = statesCount
	automaton.Rule = make(map[[3]int]int)
	automaton.Depth = depth
	automaton.BuildCells(&automaton.Cells, randomGenerator)
	automaton.BuildCells(&automaton.BackCells, nil)
	automaton.BuildHistory(randomGenerator)
	return automaton
}

func (automaton *Automaton) TryToExplode(cell *Cell) {
	cell.Energy++
	cell.Energy %= automaton.StatesCount
	if cell.Energy == 0 && cell.CanExplode {
		cell.CanExplode = false
		automaton.TryToExplode(cell.Neighbors["right"])
		automaton.TryToExplode(cell.Neighbors["down"])
		automaton.TryToExplode(cell.Neighbors["left"])
		automaton.TryToExplode(cell.Neighbors["up"])
	}
}

func (automaton *Automaton) Iterate() {
	for x := 0; x < automaton.Size.X; x++ {
		for y := 0; y < automaton.Size.Y; y++ {
			for z := 1; z < automaton.Depth; z++ {
				cell := automaton.History[z][x][y]
				cell.CanExplode = true
			}
		}
	}
	for x := 0; x < automaton.Size.X; x++ {
		for y := 0; y < automaton.Size.Y; y++ {
			cell := automaton.Cells[x][y]
			cell.CanExplode = true
		}
	}
	for x := 0; x < automaton.Size.X; x++ {
		for y := 0; y < automaton.Size.Y; y++ {
			cell := automaton.History[0][x][y]
			automaton.TryToExplode(cell)
		}
	}
	for x := 0; x < automaton.Size.X; x++ {
		for y := 0; y < automaton.Size.Y; y++ {
			for z := 1; z < automaton.Depth; z++ {
				cell := automaton.History[z][x][y]
				pastCell := automaton.History[z-1][x][y]
				if !pastCell.CanExplode {
					automaton.TryToExplode(cell)
				}
				automaton.TryToExplode(cell)
			}
		}
	}
	for x := 0; x < automaton.Size.X; x++ {
		for y := 0; y < automaton.Size.Y; y++ {
			cell := automaton.Cells[x][y]
			pastCell := automaton.History[automaton.Depth-1][x][y]
			if !pastCell.CanExplode {
				automaton.TryToExplode(cell)
			}
			automaton.TryToExplode(cell)
		}
	}
	for z := 1; z < automaton.Depth; z++ {
		automaton.Copy(automaton.History[z], automaton.History[z-1])
	}
	automaton.Copy(automaton.Cells, automaton.History[automaton.Depth-1])
	//automaton.Copy(automaton.History[automaton.HistoryIndex], automaton.Cells)
	//automaton.HistoryIndex++
	//automaton.HistoryIndex %= automaton.Depth
}

func (automaton *Automaton) Copy(a, b [][]*Cell) {
	for x := 0; x < automaton.Size.X; x++ {
		for y := 0; y < automaton.Size.Y; y++ {
			b[x][y].Energy = a[x][y].Energy
		}
	}
}

func (automaton *Automaton) BuildHistory(randomGenerator *rand.Rand) {
	automaton.History = make([][][]*Cell, automaton.Depth)
	for z := 0; z < automaton.Depth; z++ {
		automaton.BuildCells(&automaton.History[z], randomGenerator)
	}
}

func (automaton *Automaton) BuildCells(cells *[][]*Cell, randomGenerator *rand.Rand) {
	(*cells) = make([][]*Cell, automaton.Size.X)
	for x := 0; x < automaton.Size.X; x++ {
		(*cells)[x] = make([]*Cell, automaton.Size.Y)
		for y := 0; y < automaton.Size.Y; y++ {
			if randomGenerator == nil {
				(*cells)[x][y] = NewCell(NewVector2(x, y), 0)
			} else {
				(*cells)[x][y] = NewCell(NewVector2(x, y), randomGenerator.Intn(automaton.StatesCount))
			}
		}
	}
	for x := 0; x < automaton.Size.X; x++ {
		for y := 0; y < automaton.Size.Y; y++ {
			cell := (*cells)[x][y]
			if x == automaton.Size.X-1 {
				cell.Neighbors["right"] = (*cells)[0][y]
			} else {
				cell.Neighbors["right"] = (*cells)[x+1][y]
			}
			if y == automaton.Size.Y-1 {
				cell.Neighbors["down"] = (*cells)[x][0]
			} else {
				cell.Neighbors["down"] = (*cells)[x][y+1]
			}
			if x == 0 {
				cell.Neighbors["left"] = (*cells)[automaton.Size.X-1][y]
			} else {
				cell.Neighbors["left"] = (*cells)[x-1][y]
			}
			if y == 0 {
				cell.Neighbors["up"] = (*cells)[x][automaton.Size.Y-1]
			} else {
				cell.Neighbors["up"] = (*cells)[x][y-1]
			}
		}
	}
}
