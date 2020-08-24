package GameA1

import "github.com/veandco/go-sdl2/sdl"

type Rectangle struct {
	Position *Vector2
	Size     *Vector2
}

func NewRectangle(position, size *Vector2) *Rectangle {
	rectangle := new(Rectangle)
	rectangle.Position = position
	rectangle.Size = size
	return rectangle
}

func (rectangle *Rectangle) ToSDL() *sdl.Rect {
	sdlRectangle := new(sdl.Rect)
	sdlRectangle.X = int32(rectangle.Position.X)
	sdlRectangle.Y = int32(rectangle.Position.Y)
	sdlRectangle.W = int32(rectangle.Size.X)
	sdlRectangle.H = int32(rectangle.Size.Y)
	return sdlRectangle
}
