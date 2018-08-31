package alarm_clock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Sort(t *testing.T) {
	a := assert.New(t)

	expectedItemTime := []int64{1, 2, 3, 4, 5}

	qi1 := Item{
		Time: 1,
		Id:   5,
	}
	qi2 := Item{
		Time: 2,
		Id:   4,
	}
	qi3 := Item{
		Time: 3,
		Id:   3,
	}
	qi4 := Item{
		Time: 4,
		Id:   2,
	}
	qi5 := Item{
		Time: 5,
		Id:   1,
	}

	queue := List{
		list: []Item{
			qi2, qi4, qi1, qi5, qi3,
		},
	}

	queue.Sort()

	for k, v := range queue.list {
		a.Equal(expectedItemTime[k], v.Time)
	}
}

func TestList_SortedAddByTime(t *testing.T) {
	a := assert.New(t)

	qi1 := Item{
		Time: 1,
		Id:   5,
	}
	qi2 := Item{
		Time: 2,
		Id:   4,
	}
	qi3 := Item{
		Time: 3,
		Id:   3,
	}
	qi4 := Item{
		Time: 4,
		Id:   2,
	}
	qi5 := Item{
		Time: 5,
		Id:   1,
	}
	queue := List{}
	queue.SortedAddByTime(qi1)
	queue.SortedAddByTime(qi3)
	queue.SortedAddByTime(qi5)

	queue.SortedAddByTime(qi4)
	queue.SortedAddByTime(qi2)

	a.Equal(qi2, queue.list[1])
	a.Equal(qi4, queue.list[3])
}

func TestList_SortedAddByTime2(t *testing.T) {
	a := assert.New(t)

	queue := List{}
	time := []int64{3, 7, 8, 2, 4, 6, 5, 1, 9, 0} // when we add to alarm_clock, last elem must be first

	for _, v := range time {
		queue.SortedAddByTime(Item{Time: v})
	}

	for i := 0; i < 9; i++ {
		a.Equal(int64(i), queue.list[i].Time)
	}
}

func TestList_SortedAddMatchTime(t *testing.T) {
	a := assert.New(t)

	queue := List{}
	time := []int64{3, 7, 8, 1, 4, 4, 5, 1, 9, 1} // when we add to alarm_clock, last elem must be first

	expected := []int64{1, 1, 1, 3, 4, 4, 5, 7, 8, 9}

	for _, v := range time {
		queue.SortedAddByTime(Item{Time: v})
	}

	for i := 0; i < 9; i++ {
		a.Equal(expected[i], queue.list[i].Time)
	}
}

func TestQueue_Delete(t *testing.T) {
	a := assert.New(t)

	expectedItemTime := []int64{1, 2, 3, 4}

	qi1 := Item{
		Time: 1,
		Id:   5,
	}
	qi2 := Item{
		Time: 2,
		Id:   4,
	}
	qi3 := Item{
		Time: 3,
		Id:   3,
	}
	qi4 := Item{
		Time: 4,
		Id:   2,
	}
	qi5 := Item{
		Time: 5,
		Id:   1,
	}

	queue := List{}
	queue.SortedAddByTime(qi1)
	queue.SortedAddByTime(qi2)
	queue.SortedAddByTime(qi3)
	queue.SortedAddByTime(qi5)
	queue.SortedAddByTime(qi4)

	queue.Delete(qi5)

	for k, v := range queue.list {
		a.Equal(expectedItemTime[k], v.Time)
	}
}
