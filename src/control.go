package engine

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func PlayerControlInput() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		fmt.Println("up is pressed")
	}
}
