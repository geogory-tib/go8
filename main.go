package main

import (
	"go8/emu"
	"go8/types"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	window, err := sdl.CreateWindow("go8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 500, 500, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()
	_, _ = window.GetSurface()
	if err != nil {
		log.Fatal(err)
	}

	var chip8 types.Chip8
	emu.Load_rom("ibmlogo.ch8", &chip8)
	for chip8.Emu_state {
		emu.Chip8_cycle(&chip8)
	}
}
