package main

import "fmt"

/*
装饰模式（Decorator Pattern）。装饰模式是一种结构型设计模式，允许向对象添加新功能，而不改变其结构。装饰模式通过使用装饰器类包装原始对象，从而动态地添加功能，这些装饰器类通常与原始对象共享相同的接口。

装饰模式的主要特点：

 1. 透明性：装饰器类与被装饰的对象实现相同的接口，客户端无需知道对象是否被装饰。
 2. 灵活性：可以动态地添加、删除功能，不需要修改类的定义。
 3. 组合性：可以使用多个装饰器类进行组合，形成复杂的功能。

主要参与者：

 1. Component：定义对象的接口，可以动态添加功能。
 2. ConcreteComponent：实现Component接口的具体类，是要被装饰的原始对象。
 3. Decorator：抽象装饰类，持有一个Component对象，定义与Component接口一致的方法。
 4. ConcreteDecorator：具体装饰类，扩展Decorator，添加额外的功能。
*/

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
