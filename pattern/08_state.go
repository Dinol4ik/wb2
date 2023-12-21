package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

import (
	"fmt"
)

type State interface {
	Handle()
}

type Context struct {
	state State
}

func (c *Context) Request() {
	c.state.Handle()
}

func (c *Context) SetState(state State) {
	c.state = state
}

type StateA struct{}

func (s *StateA) Handle() {
	fmt.Println("ConcreteStateA.Handle()")
}

type StateB struct{}

func (s *StateB) Handle() {
	fmt.Println("ConcreteStateB.Handle()")
}

// Пример
//	ctx := &Context{state: &StateA{}}
//	ctx.Request()
//	ctx.SetState(&StateB{})
//	ctx.Request()
