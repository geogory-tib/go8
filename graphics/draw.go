package graphics

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw_Buffer(emu_screen *rl.RenderTexture2D, chip8_screen [][]bool) {
	rl.ClearBackground(color.RGBA{0, 0, 0, 0})
	rl.BeginTextureMode(*emu_screen)
	for y := range len(chip8_screen) {
		for x := range len(chip8_screen[y]) {
			if chip8_screen[y][x] {
				rl.DrawPixel(int32(x), int32(y), color.RGBA{255, 255, 255, 0})
			}
		}
	}

}
