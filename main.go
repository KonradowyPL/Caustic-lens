package main

import (
	"fmt"
	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	count int
}

var screenX int = 1000
var screenY int = 1000

var pixel *ebiten.Image

func (g *Game) Update() error {
	g.count++

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	update()
	draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenX, screenY
}

func main() {
	spawn()


	pixel = ebiten.NewImage(1, 1)
	pixel.Fill(color.RGBA{255, 255, 255, 255})

	ebiten.SetWindowSize(screenX, screenY)
	ebiten.SetWindowTitle("Animation (Ebitengine Demo)")
	err := ebiten.RunGame(&Game{})
	if err != nil {
		fmt.Println(err)
	}
}

func draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	for i := 0; i < len(particles); i++ {
		var particle = particles[i]
		var x = int(particle.x)
		var y = int(particle.y)
		opts.GeoM.Reset()
		opts.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(pixel, opts)
	}
}
