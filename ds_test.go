package ds

import (
	"fmt"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	ss := NewSearchServer([]ISearch{
		&Godaddy{},
	}, []rune{'1', '2', '3'}, 2)

	v := ss.Do()
	fmt.Printf("%v", v)

	for i := 0; i < 100; i++ {
		time.Sleep(1e9)
	}
}

func TestOne(t *testing.T) {
	ss := NewSearchServer([]ISearch{
		&Godaddy{},
	}, []rune{'1', '2', '3', '4', '5', '6'}, 6)

	fmt.Printf("%v", ss.One("1145599"))
}

func TestCombine(t *testing.T) {
	ret := GetCombineMatch([]rune{'1', '2', '3'}, 2)
	fmt.Printf("ret %v", ret)
}
