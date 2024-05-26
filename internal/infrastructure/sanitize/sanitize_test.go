package sanitize

import (
	"fmt"
	"testing"
)

type Bar struct {
	BarInt int
	BarStr string
}

type BarItem struct {
	BarSlice []Bar
}

type ForSanitize struct {
	Foo     string
	BarItem BarItem
	//Bar Bar
}

func TestSanitize(t *testing.T) {
	s := New().SetPolicy(StrictPolicy)

	forSanitize := ForSanitize{
		Foo: "<b>FOO</b>",
		//barItem: Bar{
		//	BarInt: 0,
		//	BarStr: "<a>javascript mo href=\"http://localhost\"</a><a>",
		//},
		BarItem: BarItem{
			BarSlice: []Bar{{
				BarInt: 0,
				BarStr: "<a>BAR</a>",
			}, {
				BarInt: 0,
				BarStr: "<a>javascript mo href=\"http://localhost\"</a><a>",
			}},
		}}
	fmt.Println(forSanitize)

	any, err := s.Any(forSanitize)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("any: ", any)

	s.Struct(&forSanitize)

	fmt.Println(forSanitize)
}
