package main

import "log"

type Code uint32

// func (c Code) Uint32() uint32 {
// 	return uint32(c)
// }

const (
	OK Code = iota + 1
	Error
)

type TT[T interface{}] interface {
	Status(status T)
}

type Foo struct {
}

func (f Foo) Status(status Code) {
	log.Println(status)
}

func main() {
	foo := Foo{}

	foo.Status(Error)
}
