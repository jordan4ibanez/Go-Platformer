package engine

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func playerControlInput() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		fmt.Println("up is pressed")
	}
}
