package main

import "fmt"

type printer interface {
	printInfo()
}

type Book struct {
	Title string
	Price float64
}

type Drink struct {
	Name  string
	Price float64
}

type Toy struct {
	Name  string
	Price float64
}

func (d Drink) printInfo() {
	fmt.Printf("Drink: %s Price: %.2f\n", d.Name, d.Price)
}
func (b Book) printInfo() {
	fmt.Printf("Book: %s Price: %.2f\n", b.Title, b.Price)
}
func (b Toy) printInfo() {
	fmt.Printf("Toy: %s Price: %.2f\n", b.Name, b.Price)
}

func main() {
	gunslinger := Book{
		Title: "The Gunslinger",
		Price: 4.75,
	}

	coffe := Drink{
		Name:  "Coffee",
		Price: 2.50,
	}
	rubbix := Toy{
		Name:  "Rubix Cube",
		Price: 5.00,
	}
	gunslinger.printInfo()
	coffe.printInfo()

	info := []printer{gunslinger, coffe, rubbix}
	info[0].printInfo()
	info[1].printInfo()
	info[2].printInfo()
}
