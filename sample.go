package sample

import "fmt"

type Item interface {
	GetKey() int
}

func Hello(item Item) {
	fmt.Println(item.GetKey())
}
