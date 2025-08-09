package main

import (
	"container/list"
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
		log.Fatal("gotext Error: No file specified?")
	}
	var screen termfuc.Screen
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	screen.FileSize, screen.StrList = textio.LoadFileData(file)
	file.Close()
	err = termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	Mwidth, Mheight := termbox.Size()
	termfuc.Draw(screen, Mwidth, Mheight)
	termfuc.DrawBottomBar(Mwidth, Mheight, screen, os.Args[1])
	var editNode *list.Element
	var editLen int
	defer termbox.Close()
	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			switch event.Key {
			case termbox.KeyEsc:
				if screen.InsertMode == false {
					return
				}
				screen.InsertMode = false
				editNode.Value = termfuc.GetStringAtLine(screen.CursorY, editLen)
				editNode = nil
				editLen = 0
				termfuc.Draw(screen, Mwidth, Mheight)
				termfuc.DrawBottomBar(Mwidth, Mheight, screen, os.Args[1])
			case termbox.KeyPgdn:
				if (screen.Offset + uint32(Mheight)) < uint32(screen.StrList.Len()) {
					screen.Offset += uint32(Mheight)
					termfuc.Draw(screen, Mwidth, Mheight)
					termfuc.DrawBottomBar(Mwidth, Mheight, screen, os.Args[1])
				}
			case termbox.KeyPgup:
				if screen.Offset != 0 {
					screen.Offset -= uint32(Mheight)
					termfuc.Draw(screen, Mwidth, Mheight)
					termfuc.DrawBottomBar(Mwidth, Mheight, screen, os.Args[1])
				}
			case termbox.KeyArrowDown:
				if screen.CursorY < Mheight && screen.InsertMode != true {
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
				listext.InsertValueAfter(screen.CursorY, "\n", screen.StrList)
				screen.CursorY++
				screen.CursorX = 0
				termbox.SetCursor(screen.CursorX, screen.CursorY)
				termfuc.Draw(screen, Mwidth, Mheight)

			}
		}

		if event.Ch != 0 {
			if event.Ch == 'i' && screen.InsertMode == false {
				screen.InsertMode = true
				editNode = listext.FindNodeAt(screen.CursorY, screen.StrList)
				nodeVaule := reflect.ValueOf(editNode.Value)
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
