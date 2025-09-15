package search

import (
	"gotext/termfunc"
	"reflect"
)

func FindLine(screen *termfuc.Screen, line, Mheight int) termfuc.Line {
	var retLine termfuc.Line
	var y int
	node := screen.StrList.Front()
	tabs := 0 // keep track of tabs to have proper length
	for y = 0; y < line && node != nil; y++ {
		node = node.Next()
	}
	if y < int(screen.Offset) {
		screen.Offset = uint32(y)
		y = 0

	}
	if y > Mheight {
		screen.Offset = uint32(y) - 3
		y = 3
	}
	retLine.LineY = y
	nodeValue := reflect.ValueOf(node.Value)
	nodeStr := nodeValue.String()
	for _, ch := range nodeStr { // find tabs to compensate for extra length
		if ch == 0x09 {
			tabs++
		}
	}
	retLine.LineLength = len(nodeStr) + (4 * tabs)
	return retLine
}
