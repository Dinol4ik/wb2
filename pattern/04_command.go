package pattern

/*
	Реализовать паттерн «комманда».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Command_pattern
*/
// Статья по паттерну command - https://refactoring.guru/ru/design-patterns/command/go/example
import (
	"fmt"
)

type Command interface {
	Execute()
}
type ReceiverI interface {
	Action(msg string)
}
type CommandOne struct {
	receiver *Receiver
}

func (c *CommandOne) Execute() {
	c.receiver.Action("CommandOne")
}

type CommandTwo struct {
	receiver *Receiver
}

func (c *CommandTwo) Execute() {
	c.receiver.Action("CommandTwo")
}

type Receiver struct{}

func (r *Receiver) Action(msg string) {
	fmt.Println(msg)
}

type Sequence struct {
	history []Command
}

func (i *Sequence) StoreAndExecute(command Command) {
	i.history = append(i.history, command)
	for i, command := range i.history {
		fmt.Printf("history%d: ", i)
		command.Execute()
	}
}
