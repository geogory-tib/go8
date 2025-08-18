package main

import (
	"fmt"
	"gotext/list-ext"
	"gotext/logic"
	"gotext/termfunc"
	"gotext/textio"
	"log"
	"os"
	"reflect"

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
			if event.Ch == 'i' && screen.InsertMode == false {
				screen.InsertMode = true
				screen.EditNode = listext.FindNodeAt(screen.CursorY, screen) // once you enter insert mode it grabs the node at the line
				nodeVaule := reflect.ValueOf(screen.EditNode.Value)          // reflection to cast the value to a string
				screen.EditLen = len(nodeVaule.String())
				termfuc.DrawBottomBar(Mwidth, Mheight, screen, os.Args[1])
			} else if screen.InsertMode == true {
				termbox.SetChar(screen.CursorX, screen.CursorY, event.Ch)
				screen.CursorX++
				if screen.CursorX > screen.EditLen {
					screen.EditLen += (screen.CursorX - screen.EditLen)
				}
				termbox.SetCursor(screen.CursorX, screen.CursorY)
				termbox.Flush()
			}

		}

	}
}
