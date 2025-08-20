package logic

import (
	"gotext/list-ext"
	termfuc "gotext/termfunc"
	"gotext/textio"
	"log"
	"os"
	"reflect"

	"github.com/nsf/termbox-go"
)

/*
* flow.go
*	Main file in the logic module
* Each function contains the logic for features in the editor
*
* */
const EXIT_CODE = false

// handles all the event keys that dont relate to the editing of the bufffer returns false when user wants to exit
func HandleEventKeys(screen *termfuc.Screen, event termbox.Event, Mwidth, Mheight int) bool {
	switch event.Key {
	case termbox.KeyEsc:
		if screen.InsertMode == false { // exit program if not in edit mode
			return EXIT_CODE
		} else if screen.InsertMode == true {
			screen.InsertMode = false
			screen.EditNode.Value = termfuc.GetStringAtLine(screen.CursorY, screen.EditLen)
			screen.EditNode = nil
			screen.EditLen = 0
			termfuc.Draw(*screen, Mwidth, Mheight)
			termfuc.DrawBottomBar(Mwidth, Mheight, *screen, os.Args[1])
		}
	case termbox.KeyPgdn:
		if (screen.Offset + uint32(Mheight-1)) < uint32(screen.StrList.Len()) {
			screen.Offset += uint32(Mheight - 1)
			termfuc.Draw(*screen, Mwidth, Mheight)
			termfuc.DrawBottomBar(Mwidth, Mheight, *screen, os.Args[1])
		}
	case termbox.KeyPgup:
		if screen.Offset != 0 {
			screen.Offset -= uint32(Mheight - 1)
			termfuc.Draw(*screen, Mwidth, Mheight)
			termfuc.DrawBottomBar(Mwidth, Mheight, *screen, os.Args[1])
		}
	case termbox.KeyArrowDown:
		if screen.CursorY < (Mheight-1) && screen.InsertMode != true {
			screen.CursorY++
			termbox.SetCursor(screen.CursorX, screen.CursorY)
			termbox.Flush()
		}
	case termbox.KeyArrowUp:
		if screen.CursorY > 0 && screen.InsertMode != true {
			screen.CursorY--
			termbox.SetCursor(screen.CursorX, screen.CursorY)
			termbox.Flush()
		}
	case termbox.KeyArrowRight:
		if screen.CursorX < Mwidth {
			screen.CursorX++
			termbox.SetCursor(screen.CursorX, screen.CursorY)
			termbox.Flush()
		}
	case termbox.KeyArrowLeft:
		if screen.CursorX > 0 {
			screen.CursorX--
			termbox.SetCursor(screen.CursorX, screen.CursorY)
			termbox.Flush()
		}
	case termbox.KeyCtrlW:
		if screen.InsertMode != true {
			file, err := os.Create(os.Args[1])
			if err != nil {
				log.Fatal(err)
			}
			textio.WriteFileData(file, screen.StrList)
			file.Close()
			termfuc.Draw(*screen, Mwidth, Mheight)
			termfuc.DrawBottomBar(Mwidth, Mheight, *screen, os.Args[1])
		}
	}
	return true
}

// handles event keys that retain to modifying the text buffer ex: backspace, space
func HandleEditorEventKeys(screen *termfuc.Screen, event termbox.Event, Mwidth, Mheight int) {
	if event.Key == termbox.KeyBackspace2 || event.Key == termbox.KeyDelete {
		termbox.SetChar(screen.CursorX, screen.CursorY, ' ')
		if screen.CursorX > 0 {
			screen.CursorX--
		}
		termbox.SetCursor(screen.CursorX, screen.CursorY)
		termbox.Flush()
	}
	if event.Key == termbox.KeySpace {
		termbox.SetChar(screen.CursorX, screen.CursorY, ' ')
		if screen.CursorX < Mwidth {
			screen.CursorX++
		}
		termbox.SetCursor(screen.CursorX, screen.CursorY)
		termbox.Flush()
	}
	if event.Key == termbox.KeyTab {
		for x := screen.CursorX; x < screen.CursorX+4; x++ {
			termbox.SetChar(x, screen.CursorY, ' ')
		}
		screen.CursorX += 4
		termbox.SetCursor(screen.CursorX, screen.CursorY)
		termbox.Flush()
	}
	if event.Key == termbox.KeyEnter {
		screen.EditNode.Value = termfuc.GetStringAtLine(screen.CursorY, int(screen.EditLen))
		screen.EditNode = screen.StrList.InsertAfter("\n", screen.EditNode) // if user presses enter w hile in edit mode on a line a new node is pushed after the line currently being ediited
		screen.CursorY++
		if screen.CursorY == Mheight-1 {
			screen.Offset += uint32(Mheight - 1)
			screen.CursorY = 0
		}
		screen.CursorX = 0 //cursor is reset
		termbox.SetCursor(screen.CursorX, screen.CursorY)
		termfuc.Draw(*screen, Mwidth, Mheight)
		termfuc.DrawBottomBar(Mwidth, Mheight, *screen, os.Args[1])
	}

}

// deals with adding user chars to the buffer
func HandleUserCharKeys(screen *termfuc.Screen, event termbox.Event, Mwidth, Mheight int) {
	if event.Ch == 'i' && screen.InsertMode == false {
		screen.InsertMode = true
		screen.EditNode = listext.FindNodeAt(screen.CursorY, *screen) // once you enter insert mode it grabs the node at the line
		nodeVaule := reflect.ValueOf(screen.EditNode.Value)           // reflection to cast the value to a string
		screen.EditLen = len(nodeVaule.String())
		termfuc.DrawBottomBar(Mwidth, Mheight, *screen, os.Args[1])
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
