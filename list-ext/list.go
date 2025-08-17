package listext

import (
	"container/list"
	termfuc "gotext/termfunc"
)

func FindNodeAt(line int, screen termfuc.Screen) (node *list.Element) {
	index := line + int(screen.Offset)
	node = screen.StrList.Front()
	for i := 0; i < index; i++ {
		node = node.Next()
	}
	return node
}

/*
unused function may be useful somewhere else though
func InsertValueAfter(index int, str string, list *list.List) {
	node := list.Front()
	for i := 0; i < index; i++ {
		node = node.Next()
	}
	list.InsertAfter(str, node)
}
*/
