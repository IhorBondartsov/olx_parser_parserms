package dateParser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	today             = "Сегодня"
	yestarday         = "Вчера"
	secondInDay int64 = 86400
)

var month = []string{"янв", "фев", "март", "апр", "май", "июнь", "июль", "авг", "сен", "окт", "ноя", "дек"}

func ParseTime(date string) int64 {
	if p := parseDay(date); p != 0 {
		t := parseTime(date)
		return p + int64(t)
	}

	if p := parseMonth(date); p != 0 {
		d := parseNumberDay(date)
		return p + d
	}
	return 0
}

// return UNIX
func parseDay(date string) int64 {
	days := time.Now().Unix() / secondInDay
	if strings.Contains(date, today) {
		return days * secondInDay
	}
	if strings.Contains(date, yestarday) {
		return days*secondInDay - secondInDay
	}
	return 0
}

func parseMonth(date string) int64 {
	for k, v := range month {
		if strings.Contains(date, v) {
			var month string
			if k+1 < 10 {
				month = fmt.Sprintf("0%d", k+1)
			} else {
				month = fmt.Sprintf("%d", k+1)
			}
			year := time.Now().Year()
			str := fmt.Sprintf("%d-%s-%s", year, month, "01")
			t, _ := time.Parse("2006-01-02", str)
			return t.Unix() - secondInDay
		}
	}
	return 0
}

// return UNIX time
func parseNumberDay(date string) int64 {
	var valNumDay = regexp.MustCompile(`[0-9]+`)
	numDay := valNumDay.FindString(date)
	day, _ := strconv.ParseInt(numDay, 10, 10)
	return day * secondInDay
}

// return Unix
func parseTime(strTime string) int {
	var valTimeStr = regexp.MustCompile(`[0-9]{2}:[0-9]{2}`)
	parsed := valTimeStr.FindString(strTime)
	val := strings.Split(parsed, ":")
	valInt := make([]int, len(val))

	for k, v := range val {
		valInt[k], _ = strconv.Atoi(v)
	}
	return valInt[0]*3600 + valInt[1]*60

}
