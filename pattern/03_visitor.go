package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Visitor_pattern
*/
// статья по паттерну visitor - https://refactoring.guru/ru/design-patterns/visitor
type Visitor interface {
	VisitA(elemA *ConcreteElementA)
	VisitB(elemB *ConcreteElementB)
}
type Element interface {
	Accept(Visitor)
}
type ConcreteElementA struct{}
type ConcreteElementB struct{}
type ConcreteVisitors struct{}

func (elA *ConcreteElementA) Accept(v Visitor) {
	fmt.Println("ConcreteElementA.Accept()")
	v.VisitA(elA)
}
func (elB *ConcreteElementB) Accept(v Visitor) {
	fmt.Println("ConcreteElementB.Accept()")
	v.VisitB(elB)
}
func (v *ConcreteVisitors) VisitA(elemA *ConcreteElementA) {
	fmt.Println("ConcreteVisitor.VisitA()")
}
func (v *ConcreteVisitors) VisitB(elemB *ConcreteElementB) {
	fmt.Println("ConcreteVisitor.VisitB()")
}
