package gb

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Window struct {
	window  *pixelgl.Window
	picture *pixel.PictureData
}

func Run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Goboy",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Update()
	}
}
