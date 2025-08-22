package search

import (
	"gotext/termfunc"
	"reflect"
)

func FindLine(screen *termfuc.Screen, line, Mheight int) termfuc.Line {
	var retLine termfuc.Line
	node := screen.StrList.Front()
	tabs := 0 // keep track of tabs to have proper length
	for y := 0; node != nil; y++ {
		if y == line {
			if y > Mheight-1 {
				screen.Offset = +uint32(Mheight) - 1
				y -= Mheight
				retLine.LineY = y
			}
			break
		}
		node = node.Next()
	}
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
