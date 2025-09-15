package main

import (
	"go8/emu"
	"go8/graphics"
	"go8/types"
	"image/color"

	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(500, 500, "Go-8")
	defer rl.CloseWindow()
	emulator_screen := rl.LoadRenderTexture(64, 32)
	rl.SetTextureFilter(emulator_screen.Texture, rl.TextureFilterNearest)
	rl.BeginDrawing()
	rl.ClearBackground(color.RGBA{0, 0, 0, 0})
	rl.EndDrawing()
	var chip8 types.Chip8
	emu.Load_rom("ibmlogo.ch8", &chip8)
	for !rl.WindowShouldClose() {
		emu.Chip8_cycle(&chip8)
		graphics.Draw_Buffer(&emulator_screen, chip8.Display)
		chip8.Print_State()
	}
}
