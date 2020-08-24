package GameA1

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type TextureManager struct {
	Manager  *Manager
	Textures map[string]*sdl.Texture
}

func NewTextureManager(manager *Manager) *TextureManager {
	textureManager := new(TextureManager)
	textureManager.Manager = manager
	textureManager.Textures = make(map[string]*sdl.Texture)
	names := []string{}
	for _, name := range names {
		textureManager.Load(name)
	}
	return textureManager
}

func (textureManager *TextureManager) Load(name string) {
	surface, err := img.Load("./textures/" + name + ".png")
	if err != nil {
		panic(err)
	}
	texture, err := textureManager.Manager.Drawer.Renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	textureManager.Textures[name] = texture
}
