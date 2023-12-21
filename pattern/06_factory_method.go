package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
type factory interface {
	GetFactory(firm string) (Product, error)
}

type Creator struct {
	factory factory
}

func (c *Creator) CreateFactory() {
	product, _ := c.factory.GetFactory("Adidas")
	product.SomeMethod()
}

type ConcreteCreator struct{}

func (c *ConcreteCreator) GetFactory(firm string) (Product, error) {
	if firm == "Adidas" {
		return &Adidas{}, nil
	}
	if firm == "Nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("wrong brand type passed")
}

type Product interface {
	SomeMethod()
}

type Adidas struct{}

func (a *Adidas) SomeMethod() {
	fmt.Println("Adidas.method()")
}

type Nike struct{}

func (n *Nike) SomeMethod() {
	fmt.Println("Nike.method()")
}
