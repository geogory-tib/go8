package graphics

import (
	"go8/types"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var PREVIOUS_KEY int

// runs through 2d array of booleans and white pixel if the element is true
func Draw_Buffer(emu_screen_image *rl.Image, emu_screen_texture *rl.Texture2D, chip8 *types.Chip8) {
	if !chip8.Has_Drawn && !rl.IsWindowResized() {
		return
	}
	rl.ImageClearBackground(emu_screen_image, rl.Black) // the texture for the CHIP8 screen
	rl.ClearBackground(rl.Black)
	for y := range 32 {
		for x := range 64 {
			if chip8.Display[y][x] { // draws pixel if buffer == true
				rl.ImageDrawPixel(emu_screen_image, int32(x), int32(y), rl.White)
			}
		}
	}
	rl.UpdateTexture(*emu_screen_texture, rl.LoadImageColors(emu_screen_image))
	source_rect := rl.NewRectangle(0, 0, float32(emu_screen_texture.Width), float32(emu_screen_texture.Height))
	dest_rect := rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()))
	origin := rl.NewVector2(0.0, 0.0)
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.DrawTexturePro(*emu_screen_texture, source_rect, dest_rect, origin, 0.0, rl.White)
	rl.EndDrawing()
}
func Handle_key(chip8 *types.Chip8) {
	if rl.IsKeyDown(rl.KeyLeftControl) {
		if rl.IsKeyDown(rl.KeyF) {
			chip8.Frames += 60
			if chip8.Frames == 600 {
				chip8.Frames = 60
			}
			rl.SetTargetFPS(int32(chip8.Frames))
		}
		return
	}
	if rl.IsKeyDown(rl.KeyOne) {
		chip8.Key_board[1] = true
	} else {
		chip8.Key_board[1] = false

	}
	if rl.IsKeyDown(rl.KeyTwo) {
		chip8.Key_board[2] = true
	} else {
		chip8.Key_board[2] = false
	}
	if rl.IsKeyDown(rl.KeyThree) {
		chip8.Key_board[3] = true
	} else {
		chip8.Key_board[3] = false
	}
	if rl.IsKeyDown(rl.KeyFour) {
		chip8.Key_board[0xC] = true
	} else {
		chip8.Key_board[0xC] = false
	}
	if rl.IsKeyDown(rl.KeyQ) {
		chip8.Key_board[4] = true
	} else {
		chip8.Key_board[4] = false
	}
	if rl.IsKeyDown(rl.KeyW) {
		chip8.Key_board[5] = true
	} else {
		chip8.Key_board[5] = false
	}
	if rl.IsKeyDown(rl.KeyE) {
		chip8.Key_board[6] = true
	} else {
		chip8.Key_board[6] = false

	}
	if rl.IsKeyDown(rl.KeyR) {
		chip8.Key_board[0xD] = true
	} else {
		chip8.Key_board[0xD] = false

	}
	if rl.IsKeyDown(rl.KeyA) {
		chip8.Key_board[7] = true
	} else {
		chip8.Key_board[7] = false

	}
	if rl.IsKeyDown(rl.KeyS) {
		chip8.Key_board[8] = true
	} else {
		chip8.Key_board[8] = false

	}
	if rl.IsKeyDown(rl.KeyD) {
		chip8.Key_board[9] = true
	} else {
		chip8.Key_board[9] = false

	}
	if rl.IsKeyDown(rl.KeyF) {
		chip8.Key_board[0xE] = true
	} else {
		chip8.Key_board[0xE] = false

	}
	if rl.IsKeyDown(rl.KeyZ) {
		chip8.Key_board[0xA] = true
	} else {
		chip8.Key_board[0xA] = false

	}
	if rl.IsKeyDown(rl.KeyX) {
		chip8.Key_board[0] = true
	} else {
		chip8.Key_board[0] = false

	}
	if rl.IsKeyDown(rl.KeyC) {
		chip8.Key_board[0xB] = true
	} else {
		chip8.Key_board[0xB] = false

	}
	if rl.IsKeyDown(rl.KeyV) {
		chip8.Key_board[0xF] = true
	} else {
		chip8.Key_board[0xF] = false

	}
}
