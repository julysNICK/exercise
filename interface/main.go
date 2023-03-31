package main

import "fmt"

type showInformation interface {
	showInformation()
}

type person struct {
	name string
	age  int
}

func show(s showInformation) {
	fmt.Println(s)
}

func (p person) showInformation() {
	fmt.Println(p.name)
}

type calculateProduct interface {
	calculateImpost()
	calculateTotal()
}

type productIceCream struct {
	name  string
	price float64
}

func calculate(p calculateProduct) {
	fmt.Println(p)
	p.calculateImpost()
	p.calculateTotal()
}

func (p productIceCream) calculateImpost() {

	fmt.Println(p.price + 10)
}

func (p productIceCream) calculateTotal() {

	fmt.Println(p.price + 10)
}

//3

type calculateImpostBr interface {
	calculateImpost()
	calculatePriceBr()
}

type calculateImpostUsa interface {
	calculateImpost()
	calculatePriceUsa()
}

type product1 struct {
	name  string
	price float64
}

type product2 struct {
	name  string
	price float64
}

func (p product1) calculateImpost() {
	fmt.Println(p.price + 10)
}

func (p product1) calculatePriceBr() {
	fmt.Println(p.price + 10)
}

func (p product1) calculatePriceUsa() {
	fmt.Println(p.price + 10)
}

func (p product2) calculateImpost() {
	fmt.Println(p.price + 30)
}

func (p product2) calculatePriceBr() {
	fmt.Println(p.price + 30)
}

func (p product2) calculatePriceUsa() {
	fmt.Println(p.price + 30)
}

func calculateProductBr(p calculateImpostBr) {
	fmt.Println(p)
	p.calculateImpost()
	p.calculatePriceBr()
}

func calculateProductUsa(p calculateImpostUsa) {
	fmt.Println(p)
	p.calculateImpost()
	p.calculatePriceUsa()
}
func main() {
	//1
	// p := person{name: "João", age: 20}
	// show(p)

	//2
	// p := productIceCream{name: "Cone", price: 10.00}
	// calculate(p)

	//3
	p1 := product1{name: "Cone", price: 10.00}
	p2 := product2{name: "p2", price: 20.00}

	calculateProductBr(p1)
	calculateProductUsa(p2)
}
