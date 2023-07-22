package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "BlokusAI")
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 400, 400
}

var tanSquareImage *ebiten.Image

func initiateImages() {
	var err error
	tanSquareImage, _, err = ebitenutil.NewImageFromFile("tan_square.png")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("BlokusAI")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
