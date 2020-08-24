package GameA1

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
)

type Drawer struct {
	Size     *Vector2
	Manager  *Manager
	Renderer *sdl.Renderer
	Display  *sdl.Texture
	MustDraw bool
}

func NewDrawer(manager *Manager, size *Vector2, renderer *sdl.Renderer, texture *sdl.Texture) *Drawer {
	drawer := new(Drawer)
	drawer.Manager = manager
	drawer.Renderer = renderer
	drawer.Display = texture
	drawer.Size = size
	drawer.MustDraw = true
	return drawer
}

func (drawer *Drawer) Color(theColor *color.RGBA) {
	drawer.Renderer.SetDrawColor(theColor.R, theColor.G, theColor.B, theColor.A)
}

func (drawer *Drawer) Clear() {
	drawer.Renderer.SetRenderTarget(drawer.Display)
	drawer.Renderer.Clear()
}

func (drawer *Drawer) Rectangle(rectangle *Rectangle) {
	drawer.Renderer.SetRenderTarget(drawer.Display)
	drawer.Renderer.FillRect(rectangle.ToSDL())
}

func (drawer *Drawer) Texture(rectangle *Rectangle, name string) {
	texture := drawer.Manager.TextureManager.Textures[name]
	drawer.Renderer.SetRenderTarget(drawer.Display)
	drawer.Renderer.Copy(texture, nil, rectangle.ToSDL())
}

func (drawer *Drawer) Draw() {
	drawer.Renderer.SetRenderTarget(nil)
	drawer.Renderer.Copy(drawer.Display, nil, nil)
	drawer.Renderer.Present()
}
