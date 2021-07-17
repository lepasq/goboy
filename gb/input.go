package gb

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Window struct {
	window  *pixelgl.Window
	picture *pixel.PictureData
}

var direction map[pixelgl.Button]int
var action map[pixelgl.Button]int

func configKeys() {
	direction = make(map[pixelgl.Button]int)
	action = make(map[pixelgl.Button]int)

	direction[pixelgl.KeyL] = 0 // right
	direction[pixelgl.KeyH] = 1 // left
	direction[pixelgl.KeyK] = 2 // up
	direction[pixelgl.KeyJ] = 3 // down

	action[pixelgl.KeyD] = 0 // A
	action[pixelgl.KeyF] = 1 // B
	action[pixelgl.KeyS] = 2 // SELECT
	action[pixelgl.KeyA] = 3 // START

}

func run() {
	configKeys()
	cfg := pixelgl.WindowConfig{
		Title:  "Goboy",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	var mem Memory

	for !win.Closed() {
		mem.input(win)
		fmt.Printf("%b\n", mem.IO[0])
		win.Update()
	}
}

func Run() {
	pixelgl.Run(run)
}

func (mem *Memory) input(win *pixelgl.Window) {
	for k, v := range direction {
		if win.Pressed(k) {
			mem.writeDirection(true, v)
		}
	}

	for k, v := range action {
		if win.Pressed(k) {
			mem.writeAction(true, v)
		}
	}

	for k, v := range direction {
		if win.JustReleased(k) {
			mem.writeDirection(false, v)
		}
	}

	for k, v := range action {
		if win.JustReleased(k) {
			mem.writeAction(false, v)
		}
	}

	if win.JustReleased(pixelgl.KeyEscape) {
		win.SetClosed(true)
	}
}

func (mem *Memory) writeJoypad(pressed bool, bit int) {
	mem.clearJoypad()
	mem.updateJoypad(pressed, bit)
}

func (mem *Memory) writeDirection(pressed bool, bit int) {
	mem.writeJoypad(pressed, bit)
	if pressed {
		mem.updateJoypad(pressed, 4)
		mem.updateJoypad(!pressed, 5)
	}
}

func (mem *Memory) writeAction(pressed bool, bit int) {
	mem.writeJoypad(pressed, bit)
	if pressed {
		mem.updateJoypad(pressed, 5)
		mem.updateJoypad(!pressed, 4)
	}
}

func (mem *Memory) updateJoypad(pressed bool, bit int) {
	var val uint8 = 0
	if pressed {
		val = 1
	}
	bits := mem.Read(0xFF00)
	var mask uint8 = 1 << bit
	bits = ((bits & ^mask) | (val << bit))
	mem.Write(0xFF00, bits)
}

func (mem *Memory) clearJoypad() {
	for i := range []int{0, 1, 2, 3} {
		mem.updateJoypad(false, i)
	}
}
