package GameA1

type Visuals struct {
	Array [][]Element
	Size  *Vector2
}

func NewVisuals(size *Vector2) *Visuals {
	visuals := new(Visuals)
	visuals.Size = size
	visuals.Array = BuildBidimensionalElementArray(visuals.Size)
}
