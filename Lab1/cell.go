package Lab1

type Cell struct {
	Neighbors  map[string]*Cell
	Position   *Vector2
	Energy     int
	CanExplode bool
}

func NewCell(position *Vector2, energy int) *Cell {
	cell := new(Cell)
	cell.Energy = energy
	cell.Position = position
	cell.CanExplode = true
	cell.Neighbors = make(map[string]*Cell)
	return cell
}
