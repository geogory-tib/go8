package main

import (
	"container/list"
	"fmt"
	"gotext/list-ext"
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
	var editNode *list.Element //current node being editied
	var editLen int            // the len of the edited text to know how long the edited string is
	defer termbox.Close()
	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey { // block dealing with event keys
			switch event.Key {
			case termbox.KeyEsc:
				if screen.InsertMode == false { // exit program if not in edit mode
					return
				} else if screen.InsertMode == true {
					screen.InsertMode = false
					editNode.Value = termfuc.GetStringAtLine(screen.CursorY, editLen)
					editNode = nil
					editLen = 0
					termfuc.Draw(screen, Mwidth, Mheight)
					termfuc.DrawBottomBar(Mwidth, Mheight, screen, os.Args[1])
				}
			case termbox.KeyPgdn:
				if (screen.Offset + uint32(Mheight-1)) < uint32(screen.StrList.Len()) {
					screen.Offset += uint32(Mheight - 1)
					termfuc.Draw(screen, Mwidth, Mheight)
					termfuc.DrawBottomBar(Mwidth, Mheight, screen, os.Args[1])
				}
			case termbox.KeyPgup:
				if screen.Offset != 0 {
					screen.Offset -= uint32(Mheight - 1)
					termfuc.Draw(screen, Mwidth, Mheight)
					termfuc.DrawBottomBar(Mwidth, Mheight, screen, os.Args[1])
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
					file, err = os.Create(os.Args[1])
					if err != nil {
						log.Fatal(err)
					}
					textio.WriteFileData(file, screen.StrList)
					file.Close()
					termfuc.Draw(screen, Mwidth, Mheight)
					termfuc.DrawBottomBar(Mwidth, Mwidth, screen, os.Args[1])
				}
			}

		}
		if event.Type == termbox.EventKey && screen.InsertMode == true {
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
			if event.Key == termbox.KeyEnter {
				editNode.Value = termfuc.GetStringAtLine(screen.CursorY, editLen)
				editNode = screen.StrList.InsertAfter("\n", editNode) // if user presses enter w hile in edit mode on a line a new node is pushed after the line currently being ediited
				screen.CursorY++
				if screen.CursorY == Mheight-1 {
					screen.Offset += uint32(Mheight - 1)
					screen.CursorY = 0
				}
				screen.CursorX = 0 //cursor is reset
				termbox.SetCursor(screen.CursorX, screen.CursorY)
				termfuc.Draw(screen, Mwidth, Mheight)
				termfuc.DrawBottomBar(Mwidth, Mheight, screen, os.Args[1])
			}
		}
		if event.Ch != 0 {
			if event.Ch == 'i' && screen.InsertMode == false {
				screen.InsertMode = true
				editNode = listext.FindNodeAt(screen.CursorY, screen) // once you enter insert mode it grabs the node at the line
				nodeVaule := reflect.ValueOf(editNode.Value)          // reflection to cast the value to a string
				editLen = len(nodeVaule.String())
				termfuc.DrawBottomBar(Mwidth, Mheight, screen, os.Args[1])
			} else if screen.InsertMode == true {
				termbox.SetChar(screen.CursorX, screen.CursorY, event.Ch)
				screen.CursorX++
				if screen.CursorX > editLen {
					editLen += (screen.CursorX - editLen)
				}
				termbox.SetCursor(screen.CursorX, screen.CursorY)
				termbox.Flush()
			}

		}

	}
}
