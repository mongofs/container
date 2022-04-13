package list

import (
	"fmt"
	"testing"
)

func TestList_Push(t *testing.T) {
	list := New()

	fmt.Println(list.Push("你好吗"))
	list.Push("我很好")
	list.Push("你好我好大家好")
}
