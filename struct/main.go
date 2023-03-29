package main

type Person struct {
	name string
	age  int
}

func (p *Person) newPerson(name string, age int) *Person  {
	p.name = name
	p.age = age
	return p
}

func (p *Person) getName() string {
	return p.name
}

func (p *Person) getAge() int {
	return p.age
}

//rectangle
type Rectangle struct {
	width, height int
}

func (r *Rectangle) area() int{
	return r.width * r.height
}

// circle

type Circle struct {
	radius int
}

func (c *Circle) area() int{
	return c.radius * c.radius
}

// workers

type Worker struct {
	name string
	age int
	salary int
}

func (w *Worker) newWorker(name string, age int, salary int) *Worker {
	w.name = name
	w.age = age
	w.salary = salary
	return w
}

func (w *Worker) getName() string {
	return w.name
}

func (w *Worker) getAge() int {
	return w.age
}

func (w *Worker) getSalary() int {
	return w.salary
}

//Account the bank
type Account struct {
	name string
	number int
	balance float64
}

func (a *Account) newAccount(name string, number int, balance float64) *Account {
	a.name = name
	a.number = number
	a.balance = balance
	return a
}

func (a *Account) getName() string {
	return a.name
}

func (a *Account) getNumber() int {
	return a.number
}

func (a *Account) getBalance() float64 {
	return a.balance
}

func (a *Account) deposit(value float64) {
	a.balance += value
}

func (a *Account) withdraw(value float64) {
	a.balance -= value
}

func main(){
	// p := Person{}

	// p.newPerson("João", 20)
	// println(p.getName())
	// println(p.getAge())

	// r := Rectangle{10, 20}

	// println(r.area())

	// c := Circle{
	// 	10,
	// }

	// println(c.area())

	// w := []Worker{}

	

	// w = append(w, *new(Worker).newWorker("João", 20, 1000), *new(Worker).newWorker("Maria", 30, 2000), *new(Worker).newWorker("Pedro", 40, 3000))

	// for _, worker := range w {
	// 	println("Name: ", worker.getName())
	// 	println("Age: ", worker.getAge())
	// 	println("Salary: ", worker.getSalary())
	// 	println(" ")
	// }

	a := Account{}

	a.newAccount("João", 123, 1000)
	println("before deposit")
	println(a.getName())
	println(a.getNumber())
	println(a.getBalance())
	println(" ")
	println("After deposit")
	a.deposit(1000)
	println(a.getBalance())
	println(" ")
	println("after withdraw")
	a.withdraw(1000)

	println(a.getBalance())


	
}