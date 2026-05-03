package utils

import "container/list"

func GetNthElement(l *list.List, n int) *list.Element {
	i := -1
	for e := l.Front(); e != nil && i < n; e = e.Next() {
		if i+1 == n {
			return e
		}
		i++
	}
	return nil
}
