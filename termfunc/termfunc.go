package termfuc

import (
	"container/list"
	"reflect"
	"strconv"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

// struct repersenting the screen data being displayed
type Screen struct {
	StrList          *list.List
	FileSize         int64
	Offset           uint32
	CursorX, CursorY int
	InsertMode       bool
	EditNode         *list.Element // the node in the linked list that is currently being edited
	EditLen          int           // the length of the edited string
}
type Line struct {
	LineY      int
	LineLength int
}

// prints a string starting at the given width and height
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
		tbPrint(bytesCountX+len(NoOfBytes)+3, Mheight-1, "EDIT!", termbox.ColorBlack|termbox.AttrBold, termbox.ColorWhite) // if in edit mode it will display it on the bar
	}
	for x := range Mwidth {
		termbox.SetBg(x, Mheight-1, termbox.ColorWhite) // paint the bottom bar white with black text
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
		value := reflect.ValueOf(startElement.Value) // reflection because im using the built in linked list module
		line := value.String()                       // converting the "value" type to a string then it will be iterated through in the for loop
		width = 0
		for _, ch := range line {
			termbox.SetChar(width, height, ch)
			if ch == 0x09 {
				for i := width; i < width+4; i++ {
					termbox.SetChar(i, height, ' ')
				}
				width += 4
			} else {
				width += runewidth.RuneWidth(ch)
			}
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

// builds a string from the screen buffer from a line with a specfied length
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

func HighlightLine(line Line) {
	for x := range line.LineLength {
		termbox.SetBg(x, line.LineY, termbox.ColorWhite)
	}
	termbox.Flush()
}
