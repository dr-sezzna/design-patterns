package main

import "fmt"

func main() {
	coffee := &Coffee{}
	fmt.Printf("Plain Coffee: $%.2f\n", coffee.Cost())

	milkCoffee := &MilkDecorator{
		BeverageDecorator: BeverageDecorator{
			beverage: coffee,
		},
	}
	fmt.Printf("Coffee with Milk: $%.2f\n", milkCoffee.Cost())

	sugarMilkCoffee := &SugarDecorator{
		BeverageDecorator: BeverageDecorator{
			beverage: milkCoffee,
		},
	}
	fmt.Printf("Coffee with Milk and Sugar: $%.2f\n", sugarMilkCoffee.Cost())
}

// Component 接口
type Beverage interface {
	Cost() float64
}

// ConcreteComponent 结构体
type Coffee struct{}

func (c *Coffee) Cost() float64 {
	return 5.0
}

// Decorator 结构体
type BeverageDecorator struct {
	beverage Beverage
}

func (d *BeverageDecorator) Cost() float64 {
	return d.beverage.Cost()
}

// ConcreteDecorator 结构体
type MilkDecorator struct {
	BeverageDecorator
}

func (m *MilkDecorator) Cost() float64 {
	return m.beverage.Cost() + 1.0
}

type SugarDecorator struct {
	BeverageDecorator
}

func (s *SugarDecorator) Cost() float64 {
	return s.beverage.Cost() + 0.5
}
