/*
Тип интерфейса определяет множество методов, называющееся интерфейс.
Переменная типа интерфейса может хранить значение любого типа с набором методов,
который является любым надмножеством интерфейса.
Говорят, что такой тип реализует интерфейс.
Значение неинициализированной переменной типа интерфейса равно nil
*/

package main

import "fmt"

type actions interface {
	print()
	calc()
}

func run(a actions) {
	a.print()
	a.calc()
}

type col float64

func (c col) print() {
	fmt.Printf("количество = %0.2f шт\n", c)
}
func (c col) calc() {
	c *= 0.5
	fmt.Printf("вес = %0.2f кг\n", c)
}

// func colres(c col){
// 	c.print()
// 	c.calc()
// }

type price float64

func (p price) print() {
	fmt.Printf("цена 1ед = %0.2f руб\n", p)
}
func (p price) calc() {
	p *= 100
	fmt.Printf("ИТОГО = %0.2f руб\n", p)
}

// func priceres(p price){
// 	p.print()
// 	p.calc()
// }

func main() {
	var maslo actions = col(55)
	run(maslo)
	//var maslo col = 55
	//colres(maslo)

	var baton actions = price(5)
	run(baton)
	//var baton price = 5
	//priceres(baton)

}
