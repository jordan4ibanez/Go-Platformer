package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawEntities(screen *ebiten.Image) {

	ebitenutil.DrawRect(screen, 0, 0, 20, 20, color.White)
}
