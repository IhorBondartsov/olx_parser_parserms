package alarm_clock

import (
	"sort"
)

// Item - is struct which save:
// Id - object which we keep
// Time - "wake up time" time when this obj must be send. Using as sorting value. Must be on UNIX Time (second)
type Item struct {
	Id   int
	Time int64
}

// list - keep items, have Sort() and the sorted addition
type List struct {
	list []Item
}

func NewList() *List {
	return &List{
		list: make([]Item, 0, defaultCapacityForList),
	}
}

// Sort - sorted by time
func (l *List) Sort() {
	if len(l.list) == 0 {
		return
	}
	sort.Slice(l.list, func(i, j int) bool {
		return l.list[i].Time < l.list[j].Time
	})
}

// SortedAddByTime - added element, and keep slice sorted
func (l *List) SortedAddByTime(qi Item) int {
	for k, v := range l.list {
		if qi.Time < v.Time {
			l.list = append(l.list[:k], append([]Item{qi}, l.list[k:]...)...)
			return k
		}
	}
	l.list = append(l.list, qi)
	return len(l.list) - 1
}

// DeleteByIndex - delete element from slice by index
func (l *List) DeleteByIndex(number int) {
	l.list = append(l.list[:number], l.list[number+1:]...)
}

// Delete - find element in slice end delete it
func (l *List) Delete(li Item) {
	for k, v := range l.list {
		if li == v {
			l.list = append(l.list[:k], l.list[k+1:]...)
			return
		}
	}
}

// Len - return len items array
func (l *List) Len() int {
	return len(l.list)
}

// GetItem - get Item
func (l *List) GetItem(index int) Item {
	return l.list[index]
}
