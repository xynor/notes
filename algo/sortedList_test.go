package algo

import (
	"container/list"
	"fmt"
	"math/rand"
	"testing"
)

func TestSortedList_Insert(t *testing.T) {
	sl := SortedList{}
	sl.List = list.New()
	for i := 0; i < 20; i++ {
		v := rand.Intn(99) + 1
		sl.Insert(v, func(v, current interface{}) bool {
			return v.(int) <= current.(int) //ascend
		})
	}
	for el := sl.Front(); el != nil; el = el.Next() {
		fmt.Println("v:", el.Value.(int))
	}
}
