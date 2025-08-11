package listext

import "container/list"

func FindNodeAt(index int, list *list.List) (node *list.Element) {
	node = list.Front()
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
