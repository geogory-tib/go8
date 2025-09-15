package main

import (
<<<<<<< HEAD
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
=======
	"fmt"
	"gotext/logic"
	"gotext/termfunc"
	"gotext/textio"
	"log"
	"os"

	"github.com/nsf/termbox-go"
)

/*
* Main function for "gotext" a simple text editor
* Most control flow is stored in here IO is done in the textio module
* and extended terminal functions are done in termfunc module
* list-ext module deals with functions extending on top of the linked list module
* */
func main() {
	if len(os.Args) < 2 {
		fmt.Print("gotext Error: No file specified?\n")
		os.Exit(1)
	}
	var screen termfuc.Screen
	file, err := os.Open(os.Args[1])
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(os.Args[1])
		} else {
			log.Fatal(err)
		}
	}
	screen.FileSize, screen.StrList = textio.LoadFileData(file)
	file.Close()
	err = termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	Mwidth, Mheight := termbox.Size() // size of the screen
	termfuc.Draw(screen, Mwidth, Mheight)
	termfuc.DrawBottomBar(Mwidth, Mheight, screen, os.Args[1])
	defer termbox.Close()
	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			code := logic.HandleEventKeys(&screen, event, Mwidth, Mheight) // should always return true unless esc is pressed
			if code == false {
				return
			}
		}
		if event.Type == termbox.EventKey && screen.InsertMode == true {
			logic.HandleEditorEventKeys(&screen, event, Mwidth, Mheight)
		}
		if event.Ch != 0 {
			logic.HandleUserCharKeys(&screen, event, Mwidth, Mheight)
			logic.HandleSearchKeys(&screen, event.Ch, Mwidth, Mheight)
		}
>>>>>>> cb0c755ecf22288f1ff680fa63acac29a82c8028
	}
}
