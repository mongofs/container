package list

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestList_Push(t *testing.T) {
	list := New()
	Convey("test Push ",t, func() {
		list.Push("你好吗")
		list.Push("我很好")
		list.Push("你好我好大家好")
		list.Push("你好我好大家好好好好好")

		fmt.Println()
	})

}
