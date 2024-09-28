package pkg

import (
	"fmt"

	"github.com/nbcx/hi"
)

func Run() {

	var hello = func(c *hi.Context) {
		fmt.Println("hello")
	}
	var hello2 = func(c *hi.Context) {
		fmt.Println("hello2")
	}
	rr := hi.Default()
	rr.GET("hh", hello, func(c *hi.Context) {
		c.JSON(200, map[string]any{"hello": "world"})
	}, hello2)

	rr.Run(":8989")
}

// type Context interface{}

// type Gin[T Context] struct {
// }

// type HandlerFunc[T Context] func(T)

// func (p *Gin[T]) Get(path string, f HandlerFunc[T]) {

// }

// func New[T Context](t T) *Gin[T] {
// 	return &Gin[T]{}
// }

// type Cx struct{}

// func Run() {

// 	g := New(&Cx{})

// 	g.Get("dd", func(c *Cx) {

// 	})
// }
