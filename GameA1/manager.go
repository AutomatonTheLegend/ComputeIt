package GameA1

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Manager struct {
	Visuals        *Visuals
	TextureManager *TextureManager
	Drawer         *Drawer
	TileSize       *Vector2
	LogicalSize    *Vector2
	Window         *sdl.Window
	DrawEvent      *sdl.UserEvent
	Running        bool
}

func NewManager() *Manager {
	manager := new(Manager)
	manager.LogicalSize = NewVector2(16, 8)
	manager.InitializeSDL()
	manager.CreateWindow()
	manager.CreateDrawer()
	manager.TextureManager = NewTextureManager(manager)
	drawEventType := sdl.RegisterEvents(1)
	manager.DrawEvent = new(sdl.UserEvent)
	manager.DrawEvent.Type = drawEventType
	manager.TileSize = NewVector2(manager.Drawer.Size.X/manager.LogicalSize.X, manager.Drawer.Size.Y/manager.LogicalSize.Y)
	manager.Visuals = NewVisuals()
	manager.Visuals.Text(NewVector2(0, 0), "Compute It!!!", manager.Fonts["main"])
	return manager
}

func (manager *Manager) CreateWindow() {
	window, err := sdl.CreateWindow("ComputeIt", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 1024, 512, sdl.WINDOW_SHOWN)
	if err != nil {
		img.Quit()
		sdl.Quit()
		panic(err)
	}
	manager.Window = window
}

func (manager *Manager) CreateDrawer() {
	renderer, err := sdl.CreateRenderer(manager.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		manager.Window.Destroy()
		img.Quit()
		sdl.Quit()
		panic(err)
	}
	displaySize := NewVector2(2048, 1024)
	display, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_TARGET, int32(displaySize.X), int32(displaySize.Y))
	if err != nil {
		renderer.Destroy()
		manager.Window.Destroy()
		img.Quit()
		sdl.Quit()
		panic(err)
	}
	manager.Drawer = NewDrawer(manager, displaySize, renderer, display)
}

func (manager *Manager) InitializeSDL() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	if err := img.Init(img.INIT_PNG); err != nil {
		sdl.Quit()
		panic(err)
	}
}

func (manager *Manager) Main() {
	defer sdl.Quit()
	defer img.Quit()
	defer manager.Window.Destroy()
	defer manager.Drawer.Renderer.Destroy()
	defer manager.Drawer.Display.Destroy()
	manager.Running = true
	for manager.Running {
		event := sdl.WaitEvent()
		switch event.(type) {
		case *sdl.QuitEvent:
			manager.Running = false
		case *sdl.UserEvent:
			specialEvent := event.(*sdl.UserEvent)
			switch specialEvent.Type {
			case manager.DrawEvent.Type:
				manager.Draw()
			}
		}
	}
}

func (manager *Manager) RequestDraw() {
	filtered, err := sdl.PushEvent(manager.DrawEvent)
	if filtered {
		fmt.Println("Draw request was filtered")
	}
	if err != nil {
		panic(err)
	}
}

func (manager *Manager) Draw() {

}
