package client

import (
	"fmt"
	"testing"

	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"
)

func Test_FiltredAdv(t *testing.T) {

	a1 := entities.Advertisement{
		ID:    1,
		URL:   "1",
		Title: "1",
	}
	a2 := entities.Advertisement{
		ID:    2,
		URL:   "2",
		Title: "2",
	}
	a3 := entities.Advertisement{
		ID:    3,
		URL:   "3",
		Title: "3",
	}
	a4 := entities.Advertisement{
		ID:    4,
		URL:   "4",
		Title: "4",
	}
	a5 := entities.Advertisement{
		ID:    5,
		URL:   "5",
		Title: "5",
	}

	savedAdv := []entities.Advertisement{a1, a2, a3}
	newAdv := []entities.Advertisement{a1, a2, a3, a4, a5}
	expected := []entities.Advertisement{a4, a5}

	var actual []entities.Advertisement
	for _, v := range newAdv {
		has := false
		for _, sv := range savedAdv {
			if sv.URL == v.URL {
				has = true
				break
			}
		}
		if !has {
			actual = append(actual, v)
		}
	}

	fmt.Println(actual)
	fmt.Println(expected)

}
