package alarm_clock

import (
	"time"

	"github.com/powerman/narada-go/narada"
)

var log = narada.NewLog("alarm_clock: ")

const (
	defaultCapacityForList = 16
)

// AlarmClock - is alarm clock for item where,
// List - is list item which will have to wake up
// addChan - this chan used for adding items into the list
// stopChan - stopped alarm clock work()
// ownTimer - this timer has is occupy flag which says that timer has already worked
type AlarmClock struct {
	List     *List
	addChan  chan Item
	stopChan chan struct{}
	To       chan int
	ownTimer specialTimer
}

// specialTimer - its struct have timerIsOccupy.
// timerIsOccupy - if timer is working value is true, if timer is stop or expired value is false
type specialTimer struct {
	timer         *time.Timer
	timerIsOccupy bool
}

func NewAlarmClock(to chan int, sizeBufferAddChan int) *AlarmClock {
	return &AlarmClock{
		List:     NewList(),
		addChan:  make(chan Item, sizeBufferAddChan),
		stopChan: make(chan struct{}),
		ownTimer: specialTimer{
			timerIsOccupy: false,
		},
		To: to,
	}
}

func (ac *AlarmClock) Start() {
	log.DEBUG("AlarmClock started...")
	ac.ownTimer.timer = time.NewTimer(0)
	go ac.work()

}

// work - must work as goroutine.
// When into addChan come a value we check timer occupation if timer is free we switch on timer.
// When timer is ending we check all element in list
func (ac *AlarmClock) work() {
	log.DEBUG("work started...")
	for {
		select {
		case item := <-ac.addChan:
			i := ac.List.SortedAddByTime(item)
			if i == 0 {
				ac.ownTimer.timer.Reset(time.Second * 1)
			}
			if !ac.ownTimer.timerIsOccupy {
				sleepTime := time.Duration(item.Time - time.Now().Unix())
				ac.ownTimer.timer = time.NewTimer(time.Second * time.Duration(sleepTime))
				ac.ownTimer.timerIsOccupy = true
			}
		case <-ac.ownTimer.timer.C:
			ac.checkList()

		case <-ac.stopChan:
			return
		}
	}
}

// checkList - if element in list is expired we send this element into channel, and delete from array.
// If not expired we create new timer and interrupt execution. After checking list we set timer occupation as false
func (ac *AlarmClock) checkList() {
	for i := 0; i < ac.List.Len(); i++ {

		now := time.Now().Unix()
		item := ac.List.GetItem(i)

		if item.Time > now {
			sleepTime := time.Duration(item.Time - now)
			ac.ownTimer.timer.Reset(time.Second * sleepTime)
			return
		}

		go func(item Item) { ac.To <- item.Id }(item)
		ac.List.DeleteByIndex(i)
		i--
	}

	if ac.List.Len() != 0 {
		ac.ownTimer.timer.Reset(time.Second * 2)
	} else {
		ac.ownTimer.timerIsOccupy = false
	}
}

func (ac *AlarmClock) GetAddChan() chan Item {
	return ac.addChan
}

func (ac *AlarmClock) GetStopChan() chan struct{} {
	return ac.stopChan
}
