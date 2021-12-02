package main

import (
	"image"
	"log"
	"os"

	_ "image/png"

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
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowSizeLimits(680, 480, -1, -1)
	ebiten.SetWindowResizable(true)

	//this should probably handle errors but no one will play this game anyways
	f, _ := os.Open("textures/flarple.png")
	defer f.Close()
	img, _, _ := image.Decode(f)
	var test []image.Image = []image.Image{img}

	ebiten.SetWindowIcon(test)
	ebiten.SetWindowTitle("Go Flarple Terror 2")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
