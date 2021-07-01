package algo

import "container/list"

type SortedList struct {
	*list.List
}

func (sl *SortedList) Insert(v interface{}, f func(v, current interface{}) bool) {
	for el := sl.Front(); el != nil; el = el.Next() {
		if f(v, el.Value) { //v >= el.Value
			sl.InsertBefore(v, el)
			return
		}
	}
	sl.PushBack(v)
}
