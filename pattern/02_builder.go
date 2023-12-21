package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Car struct {
	Wheel  string
	Engine string
	Color  string
}

type CarBuilderI interface {
	Wheel(firm string) CarBuilderI
	Engine(v string) CarBuilderI
	Color(color string) CarBuilderI
	Build() Car
}
type CarBuilder struct {
	Car
}

func (c *CarBuilder) Wheel(firm string) CarBuilderI {
	c.Car.Wheel = firm
	return c
}
func (c *CarBuilder) Engine(v string) CarBuilderI {
	c.Car.Engine = v
	return c
}
func (c *CarBuilder) Color(color string) CarBuilderI {
	c.Car.Color = color
	return c
}
func (c *CarBuilder) Build() Car {
	return Car{
		Wheel:  c.Car.Wheel,
		Engine: c.Car.Engine,
		Color:  c.Car.Color,
	}
}
