package GameA1

type Vector2 struct {
	X int
	Y int
}

func NewVector2(x, y int) *Vector2 {
	vector := new(Vector2)
	vector.X = x
	vector.Y = y
	return vector
}
