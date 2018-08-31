package alarm_clock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAlarmClock_Work(t *testing.T) {
	a := assert.New(t)

	to := make(chan int)

	ac := NewAlarmClock(to, 17)
	ac.Start()

	addChan := ac.GetAddChan()
	stopChan := ac.GetStopChan()

	now := time.Now().Unix()

	ecpected := []int{1, 2, 3, 4, 5, 6}

	item := Item{
		Time: now + int64(2),
		Id:   1,
	}
	item1 := Item{
		Time: now + int64(3),
		Id:   2,
	}
	item2 := Item{
		Time: now + int64(4),
		Id:   3,
	}
	item3 := Item{
		Time: now + int64(5),
		Id:   4,
	}
	item4 := Item{
		Time: now + int64(6),
		Id:   5,
	}
	item5 := Item{
		Time: now + int64(7),
		Id:   6,
	}

	go func() { addChan <- item }()
	go func() { addChan <- item1 }()
	go func() { addChan <- item2 }()
	go func() { addChan <- item3 }()
	go func() { addChan <- item4 }()
	go func() { addChan <- item5 }()

	res := <-ac.To
	res1 := <-ac.To
	res2 := <-ac.To
	res3 := <-ac.To
	res4 := <-ac.To
	res5 := <-ac.To

	time.Sleep(2 * time.Second)

	stopChan <- struct{}{}

	result := []int{res, res1, res2, res3, res4, res5}

	for _, v := range ecpected {
		a.Contains(result, v)
	}
}

func TestAlarmClock_WorkDublicateTime(t *testing.T) {
	a := assert.New(t)

	to := make(chan int)

	ac := NewAlarmClock(to, 32)
	ac.Start()

	addChan := ac.GetAddChan()
	stopChan := ac.GetStopChan()

	now := time.Now().Unix()

	ecpected := []int{1, 2, 3, 4, 5, 6}

	item := Item{
		Time: now + int64(2),
		Id:   1,
	}
	item1 := Item{
		Time: now + int64(2),
		Id:   2,
	}
	item2 := Item{
		Time: now + int64(2),
		Id:   3,
	}
	item3 := Item{
		Time: now + int64(2),
		Id:   4,
	}
	item4 := Item{
		Time: now + int64(2),
		Id:   5,
	}
	item5 := Item{
		Time: now + int64(2),
		Id:   6,
	}

	go func() { addChan <- item }()
	go func() { addChan <- item1 }()
	go func() { addChan <- item2 }()
	go func() { addChan <- item3 }()
	go func() { addChan <- item4 }()
	go func() { addChan <- item5 }()

	res := <-ac.To
	res1 := <-ac.To
	res2 := <-ac.To
	res3 := <-ac.To
	res4 := <-ac.To
	res5 := <-ac.To

	time.Sleep(2 * time.Second)

	stopChan <- struct{}{}

	result := []int{res, res1, res2, res3, res4, res5}

	for _, v := range ecpected {
		a.Contains(result, v)
	}
}

func TestAlarmClock_WorkFirstBigAfterSmall(t *testing.T) {
	a := assert.New(t)

	to := make(chan int)

	ac := NewAlarmClock(to, 1)
	ac.Start()

	addChan := ac.GetAddChan()
	stopChan := ac.GetStopChan()

	now := time.Now().Unix()

	ecpected := []int{1, 2, 3, 4, 5, 6}

	item := Item{
		Time: now + int64(10),
		Id:   1,
	}
	item1 := Item{
		Time: now + int64(9),
		Id:   2,
	}
	item2 := Item{
		Time: now + int64(8),
		Id:   3,
	}
	item3 := Item{
		Time: now + int64(0),
		Id:   4,
	}
	item4 := Item{
		Time: now + int64(7),
		Id:   5,
	}
	item5 := Item{
		Time: now + int64(0),
		Id:   6,
	}

	go func() { addChan <- item }()
	go func() { addChan <- item1 }()
	go func() { addChan <- item2 }()
	go func() { addChan <- item3 }()
	go func() { addChan <- item4 }()
	go func() { addChan <- item5 }()

	res := <-ac.To
	res1 := <-ac.To
	res2 := <-ac.To
	res3 := <-ac.To
	res4 := <-ac.To
	res5 := <-ac.To

	time.Sleep(2 * time.Second)

	stopChan <- struct{}{}

	result := []int{res, res1, res2, res3, res4, res5}

	for _, v := range ecpected {
		a.Contains(result, v)
	}
}
