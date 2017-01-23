package ds

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	ss := &SearchServer{
		Apis: []ISearch{
			&Godaddy{},
		},
		Chars:   []rune{'1', '2', '3', '4', '5', '6'},
		CharNum: 6,
	}
	fmt.Printf("%v", ss.One("1145599"))

	// fmt.Printf("%v", ss.Do())
}

func TestCombine(t *testing.T) {
	ret := GetCombineMatch([]rune{'1', '2', '3', '4', '5'}, 5)
	fmt.Printf("ret %v", ret)
}
