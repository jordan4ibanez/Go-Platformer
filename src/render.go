package engine

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//cache reuse variables

var currentPosition [2]float64 = [2]float64{0, 0}

var currentSize [2]float64 = [2]float64{0, 0}

var player *ebiten.Image

var flarple *ebiten.Image

var tile *ebiten.Image

var imageManipulation = &ebiten.DrawImageOptions{}

//

func GraphicsInitialization() {
	var err error //gc swept

	player, _, err = ebitenutil.NewImageFromFile("textures/go-gopher.png")

	if err != nil {
		log.Fatal(err)
	}

	flarple, _, err = ebitenutil.NewImageFromFile("textures/flarple.png")

	if err != nil {
		log.Fatal(err)
	}

	tile, _, err = ebitenutil.NewImageFromFile("textures/tile.png")

	if err != nil {
		log.Fatal(err)
	}
}

func DrawEntities(screen *ebiten.Image) {

	imageManipulation.GeoM.Reset()

	imageManipulation.GeoM.Translate(getCameraPositionX(), getCameraPositionY())

	//map has be loaded inversely because data structure
	for y := 0; y < GetMapHeight(); y++ {
		for x := 0; x < GetMapLength(); x++ {
			if GameMap[y][x] > 0 {
				screen.DrawImage(tile, imageManipulation)
			}
			imageManipulation.GeoM.Translate(32, 0)
		}

		imageManipulation.GeoM.Translate(-32*float64(GetMapLength()), 32)
	}

	for i := 0; i < GetNumberOfEntities(); i++ {

		imageManipulation.GeoM.Reset()

		if i == 0 {
			imageManipulation.GeoM.Translate((640/2)+(GetWidth(0)/2), (480/2)+(GetHeight(0)/2))
			imageManipulation.GeoM.Scale(1, 1)
			screen.DrawImage(player, imageManipulation)
		} else {

			currentPosition = GetPosition(i)

			currentSize = GetSize(i)

			imageManipulation.GeoM.Translate(GetPositionX(i)+getCameraPositionX(), GetPositionY(i)+getCameraPositionY())
			imageManipulation.GeoM.Scale(1, 1)
			screen.DrawImage(flarple, imageManipulation)

			//ebitenutil.DrawRect(screen, currentPosition[0], currentPosition[1], currentSize[0], currentSize[1], color.White)
		}
	}

	//screen.DrawImage(player, nil)

	//ebiten.SetWindowTitle(fmt.Sprintf("Hello, World! %v", CalculateFPS())) this basically chops the FPS down by 25-ish percent

	ebitenutil.DebugPrint(screen, fmt.Sprint(onGround))
}
