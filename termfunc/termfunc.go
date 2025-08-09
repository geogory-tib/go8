package termfuc

import (
	"container/list"
	"reflect"
	"strconv"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Screen struct {
	StrList          *list.List
	FileSize         int64
	Offset           uint32
	CursorX, CursorY int
	InsertMode       bool
}

// prints a string starting at the given input
func tbPrint(width, height int, str string, fg, bg termbox.Attribute) {
	x := width
	y := height
	for _, ch := range str {
		termbox.SetCell(x, y, ch, fg, bg)
		x += runewidth.RuneWidth(ch)
	}
}

// draws the bar at the bottom of the screen
func DrawBottomBar(Mwidth, Mheight int, screenData Screen, fileName string) {
	lineCountX := (len(fileName) + len("File: ")) + 3                                    // the start postion for the line counter at the bottom
	tbPrint(0, Mheight-1, ("File: " + fileName), termbox.ColorBlack, termbox.ColorWhite) // prints the fileName of the file
	NoOfLines := strconv.Itoa(screenData.StrList.Len())                                  // convert the length of the linked list into a string
	tbPrint(lineCountX, Mheight-1, NoOfLines, termbox.ColorBlack, termbox.ColorWhite)    // prints the number of lines
	bytesCountX := (len(NoOfLines) + lineCountX + 3)
	NoOfBytes := strconv.Itoa(int(screenData.FileSize))
	tbPrint(bytesCountX, Mheight-1, NoOfBytes, termbox.ColorBlack, termbox.ColorWhite)
	if screenData.InsertMode == true {
		tbPrint(bytesCountX+3, Mheight-1, "EDIT!", termbox.ColorBlack|termbox.AttrBold, termbox.ColorWhite) // if in edit mode it will display it on the bar
	}
	for x := range Mwidth {
		termbox.SetBg(x, Mheight-1, termbox.ColorWhite) // paint the bottom bar white with text black
		termbox.SetFg(x, Mheight-1, termbox.ColorBlack)
	}
	termbox.Flush()
}

// draws the text on the screen -1 for the search bar
func Draw(screenData Screen, Mwidth, Mheight int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	startElement := screenData.StrList.Front()
	var height, width int
	for i := 0; uint32(i) < screenData.Offset; i++ { //for loop that iterates through the list screenData.offset times to account for the lines scrolled
		startElement = startElement.Next()
	}

	for height < Mheight-1 && startElement != nil {
		value := reflect.ValueOf(startElement.Value) // reflection because im using the built in linked list libary
		line := value.String()                       // converting the "value" type to a string then it will be iterated through in the for loop
		width = 0
		for _, ch := range line {
			termbox.SetChar(width, height, ch)
			width += runewidth.RuneWidth(ch)
			if width > Mwidth {
				break
			}
		}
		height++
		startElement = startElement.Next()
	}
	termbox.SetCursor(screenData.CursorX, screenData.CursorY)
	termbox.Flush()
}

func GetStringAtLine(lineY, length int) (str string) {
	x := 0
	var retString string
	for x < length {

		cell := termbox.GetCell(x, lineY)
		retString += string(cell.Ch)
		x += runewidth.RuneWidth(cell.Ch)
	}
	retString += "\n"
	return retString
}
