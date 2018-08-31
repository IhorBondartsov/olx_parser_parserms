package dateParser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTime(t *testing.T) {
	a := assert.New(t)

	cases := map[string]int64{
		"Сегодня 19:56": 1527796560,
		"Сегодня 00:00": 1527724800,
		"Вчера 19:56":   1527710160,
		"27 май":        1527379200,
		"2 май ":        1525219200,
		"Fail value":    0,
	}

	for k, v := range cases {
		ti := ParseTime(k)
		a.Equal(v, ti)
	}
}
