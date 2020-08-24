package Lab1

import "github.com/veandco/go-sdl2/sdl"

type MainMenu struct {
	Manager *Manager
}

func NewMainMenu(manager *Manager) *MainMenu {
	mainMenu := new(MainMenu)
	mainMenu.Manager = manager
	return mainMenu
}

func (mainMenu *MainMenu) React(event sdl.Event) {

}
