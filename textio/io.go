package textio

import (
	"bufio"
	"container/list"
	"io"
	"log"
	"os"
	"reflect"
)

func LoadFileData(file *os.File) (bytes int64, rlist *list.List) {
	retList := list.New()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		retList.PushBack(str)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	fileStat, _ := file.Stat()
	bytes = fileStat.Size()

	return bytes, retList
}

func WriteFileData(file *os.File, list *list.List) (bytes int64) {
	writer := bufio.NewWriter(file)
	node := list.Front()
	for node != nil {
		value := reflect.ValueOf(node.Value)
		str := value.String()
		bytesWrote, err := writer.WriteString(str)
		bytes += int64(bytesWrote)
		if err != nil {
			log.Fatal(err)
		}
		node = node.Next()
	}
	writer.Flush()
	return bytes
}
