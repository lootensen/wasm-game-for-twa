package app

import (
	"fmt"
	"image/color"
	"log"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/lootensen/wasm-demo-game/src/component"
)

type Game struct {
	btn      []*component.Button
	TouchIDs []ebiten.TouchID
}

func (g *Game) Update() error {

	inputs := g.gatherInputs() // Collect all inputs
	for _, input := range inputs {
		for _, btn := range g.btn {
			btn.Update(input.X, input.Y)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0x1f, 0xa0, 0xff})
	for _, btn := range g.btn {
		btn.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) gatherInputs() []struct{ X, Y int } {
	var inputs []struct{ X, Y int }

	// Collect touch inputs
	for _, id := range ebiten.AppendTouchIDs(nil) {
		if inpututil.IsTouchJustReleased(id) {
			fmt.Println("IsTouchJustReleased")

			x, y := ebiten.TouchPosition(id)
			inputs = append(inputs, struct{ X, Y int }{X: x, Y: y})
		}
	}

	// Collect mouse input if the left button is just pressed
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		fmt.Println("IsMouseButtonJustPressed")
		x, y := ebiten.CursorPosition()
		inputs = append(inputs, struct{ X, Y int }{X: x, Y: y})
	}

	return inputs
}

func RunApp(jsInterface *component.JSInterface) {
	btn := &component.Button{
		X:           110,
		Y:           100,
		Width:       100,
		Height:      60,
		JSInterface: jsInterface,
		OnClickHandler: func() {
			jsInterface.CallFunction("Telegram.WebApp.expand", nil)
			jsInterface.CallFunction("window.console.log", "Test Output")
		},
	}

	closeBtn := &component.Button{
		X:           250,
		Y:           60,
		Width:       60,
		Height:      60,
		JSInterface: jsInterface,
		OnClickHandler: func() {
			jsInterface.CallFunction("Telegram.WebApp.close", nil)
		},
	}
	game := &Game{
		btn:      []*component.Button{btn, closeBtn},
		TouchIDs: make([]ebiten.TouchID, 0),
	}
	ebiten.SetWindowSize(720, 1280)
	ebiten.SetWindowTitle("Fill")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
