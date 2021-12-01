package main

import (
	"image/color"
	"log"

	physics "github.com/jordan4ibanez/Go-Platformer/src"

	deltaTime "github.com/jordan4ibanez/Go-Platformer/src"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {

	physics.TestCode()

	deltaTime.CalculateDelta()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, physics.X, physics.Y, 20, 20, color.White)
	//ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
