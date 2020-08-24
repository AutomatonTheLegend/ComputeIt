package Lab1

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Manager struct {
	Window          *sdl.Window
	Drawer          *Drawer
	TextureManager  *TextureManager
	DrawEvent       *sdl.UserEvent
	IterationEvent  *sdl.UserEvent
	Running         bool
	WaitGroup       sync.WaitGroup
	Automaton       *Automaton
	CellSize        *Vector2
	Colors          []*color.RGBA
	MaxColorCount   float64
	ColorStep       float64
	RandomGenerator *rand.Rand
}

func NewManager() *Manager {
	manager := new(Manager)
	manager.Running = true
	manager.MaxColorCount = math.Pow(2, 24)
	//manager.MaxColorCount = math.Pow(2, 2)
	source := rand.NewSource(0)
	manager.RandomGenerator = rand.New(source)
	manager.Automaton = NewAutomaton(manager.RandomGenerator, NewVector2(64, 32), 16, 31)
	manager.BuildColors()
	return manager
}

func (manager *Manager) BuildColors() {
	manager.ColorStep = (manager.MaxColorCount - 1) / float64(manager.Automaton.StatesCount-1)
	//fmt.Println(manager.MaxColorCount)
	values := make([]float64, manager.Automaton.StatesCount)
	values[0] = 0
	for i := 1; i < manager.Automaton.StatesCount; i++ {
		values[i] = values[i-1] + manager.ColorStep
	}
	//fmt.Println(values)
	valuesUInt32 := make([]uint32, manager.Automaton.StatesCount)
	for i := 0; i < manager.Automaton.StatesCount; i++ {
		valuesUInt32[i] = uint32(values[i])
	}
	//fmt.Println(valuesUInt32)
	manager.Colors = make([]*color.RGBA, manager.Automaton.StatesCount)
	for i := 0; i < manager.Automaton.StatesCount; i++ {
		manager.Colors[i] = new(color.RGBA)
		theColor := manager.Colors[i]
		theColor.A = 255
		theColor.R = byte(valuesUInt32[i] >> (8 * 0))
		theColor.G = byte(valuesUInt32[i] >> (8 * 1))
		theColor.B = byte(valuesUInt32[i] >> (8 * 2))
	}
	//fmt.Println(manager.Colors)
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

func (manager *Manager) RequestIteration() {
	filtered, err := sdl.PushEvent(manager.IterationEvent)
	if filtered {
		fmt.Println("Iteration request was filtered")
	}
	if err != nil {
		panic(err)
	}
}

func (manager *Manager) Iterator() {
	defer manager.WaitGroup.Done()
	for manager.Running {
		//fmt.Println("Iteration")
		time.Sleep(16 * time.Millisecond)
		manager.RequestIteration()
	}
}

func (manager *Manager) Main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	if err := img.Init(img.INIT_PNG); err != nil {
		panic(err)
	}
	defer img.Quit()
	window, err := sdl.CreateWindow("ComputeIt", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 1024, 512, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	size := NewVector2(2048, 1024)
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_TARGET, int32(size.X), int32(size.Y))
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()
	manager.Drawer = NewDrawer(manager, size, renderer, texture)
	manager.TextureManager = NewTextureManager(manager)

	drawEventType := sdl.RegisterEvents(1)
	manager.DrawEvent = new(sdl.UserEvent)
	manager.DrawEvent.Type = drawEventType
	iterationEventType := sdl.RegisterEvents(1)
	manager.IterationEvent = new(sdl.UserEvent)
	manager.IterationEvent.Type = iterationEventType
	manager.CellSize = NewVector2(manager.Drawer.Size.X/manager.Automaton.Size.X, manager.Drawer.Size.Y/manager.Automaton.Size.Y)
	manager.WaitGroup.Add(1)
	go manager.Iterator()
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
			case manager.IterationEvent.Type:
				manager.Automaton.Iterate()
				manager.RequestDraw()
			}
		}
	}
	manager.WaitGroup.Wait()
}

func (manager *Manager) Draw() {
	//fmt.Println("Drawing")
	for x := 0; x < manager.Automaton.Size.X; x++ {
		for y := 0; y < manager.Automaton.Size.Y; y++ {
			position := NewVector2(x*manager.CellSize.X, y*manager.CellSize.Y)
			rectangle := NewRectangle(position, manager.CellSize)
			energy := manager.Automaton.Cells[x][y].Energy
			manager.Drawer.Color(manager.Colors[energy])
			manager.Drawer.Rectangle(rectangle)
		}
	}
	manager.Drawer.Draw()
}
