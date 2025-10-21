package main

import (
	"fmt"
	"go8/emu"
	"go8/graphics"
	"go8/types"
	"image/color"
	"os"

	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No rom provided. Exiting")
		return
	}
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(640, 320, "Go-8")

	defer rl.CloseWindow()
	emu_screen_image := rl.GenImageColor(64, 32, rl.Black)
	emu_screen_text := rl.LoadTextureFromImage(emu_screen_image)
	//emulator_screen := rl.LoadRenderTexture(64, 32)
	//rl.SetTextureFilter(emulator_screen.Texture, rl.TextureFilterNearest)
	var chip8 types.Chip8
	rl.SetTargetFPS(int32(chip8.Frames))
	rl.BeginDrawing()
	rl.ClearBackground(color.RGBA{0, 0, 0, 0})
	rl.EndDrawing()
	emu.Load_rom(os.Args[1], &chip8)
	for !rl.WindowShouldClose() {
		emu.Chip8_cycle(&chip8)
		graphics.Draw_Buffer(emu_screen_image, &emu_screen_text, &chip8)
		graphics.Handle_key(&chip8)
	}
}
