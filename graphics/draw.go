package graphics

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// runs through 2d array of booleans and white pixel if the element is true
func Draw_Buffer(emu_screen *rl.RenderTexture2D, chip8_screen [32][64]bool) {
	rl.BeginTextureMode(*emu_screen) // the texture for the CHIP8 screen
	for y := range len(chip8_screen) {
		for x := range len(chip8_screen[y]) {
			if chip8_screen[y][x] { // draws pixel if buffer == true
				rl.DrawPixel(int32(x), int32(y), rl.White)
			}
		}
	}

	rl.EndTextureMode()
	source_rect := rl.NewRectangle(0, 0, float32(emu_screen.Texture.Width), float32(-emu_screen.Texture.Height))
	dest_rect := rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()))
	origin := rl.NewVector2(0.0, 0.0)
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.DrawTexturePro(emu_screen.Texture, source_rect, dest_rect, origin, 0.0, rl.White)
	rl.EndDrawing()
}
