package main

import (
	"log"

	physics "github.com/jordan4ibanez/Go-Platformer/src"

	deltaTime "github.com/jordan4ibanez/Go-Platformer/src"

	render "github.com/jordan4ibanez/Go-Platformer/src"

	control "github.com/jordan4ibanez/Go-Platformer/src"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var delta float64 = 0

func init() {
	render.GraphicsInitialization()
}

func (g *Game) Update() error {

	delta = deltaTime.CalculateDelta()

	control.PlayerControlInput(delta)

	physics.RunPhysics(delta)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	render.DrawEntities(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Go Flarple Terror 2")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
