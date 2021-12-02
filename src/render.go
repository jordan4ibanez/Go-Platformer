package engine

import (
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

var imageManipulation = &ebiten.DrawImageOptions{}

//

func GraphicsInitialization() {
	var err error

	player, _, err = ebitenutil.NewImageFromFile("textures/go-gopher.png")

	if err != nil {
		log.Fatal(err)
	}

	flarple, _, err = ebitenutil.NewImageFromFile("textures/flarple.png")

	if err != nil {
		log.Fatal(err)
	}
}

func DrawEntities(screen *ebiten.Image) {

	for i := 0; i < GetNumberOfEntities(); i++ {

		imageManipulation.GeoM.Reset()

		if i == 0 {
			imageManipulation.GeoM.Translate(GetPositionX(0), GetPositionY(0))
			imageManipulation.GeoM.Scale(1, 1)
			screen.DrawImage(player, imageManipulation)
		} else {

			currentPosition = GetPosition(i)

			currentSize = GetSize(i)

			imageManipulation.GeoM.Translate(GetPositionX(i), GetPositionY(i))
			imageManipulation.GeoM.Scale(1, 1)
			screen.DrawImage(flarple, imageManipulation)

			//ebitenutil.DrawRect(screen, currentPosition[0], currentPosition[1], currentSize[0], currentSize[1], color.White)
		}
	}

	//screen.DrawImage(player, nil)

	//ebiten.SetWindowTitle(fmt.Sprintf("Hello, World! %v", CalculateFPS())) this basically chops the FPS down by 25-ish percent
}
