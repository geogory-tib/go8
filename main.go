package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"go8/emu"
	"go8/types"
)

func main() {
	rl.InitWindow(500, 500, "Go-8")
	defer rl.CloseWindow()
	emulator_screen := rl.LoadRenderTexture(64, 32)
	rl.SetTextureFilter(emulator_screen.Texture, rl.TextureFilterNearest)

	var chip8 types.Chip8
	emu.Load_rom("ibmlogo.ch8", &chip8)
	for !rl.WindowShouldClose() {
		emu.Chip8_cycle(&chip8)
		chip8.Print_State()
	}
}
