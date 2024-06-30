package main

import "fmt"

/*
建造者模式（Builder Pattern）是一种创建型设计模式，它将一个复杂对象的构建过程与其表示分离，使得同样的构建过程可以创建不同的表示。
通过使用建造者模式，可以一步一步地构建一个复杂的对象，而不必知道具体的构建细节。

建造者模式的主要特点：

	1.	将构建过程与表示分离：建造者模式将对象的构建过程封装在不同的建造者类中，与表示分离。
	2.	步骤灵活：可以通过不同的建造者实现不同的构建步骤，从而构建出不同的对象。
	3.	逐步构建：允许逐步构建对象，并在构建过程中提供更好的控制。

主要参与者：

	1.	Builder：为创建产品对象定义一个抽象接口，包含构建产品的各个部件的抽象方法。
	2.	ConcreteBuilder：实现Builder接口，具体构建并装配各个部件。
	3.	Director：构建一个使用Builder接口的对象，主要负责安排构建过程。
	4.	Product：表示被构建的复杂对象，包含各个部件的定义。
*/

func main() {
	builder := NewConcreteBuilder()
	director := NewDirector(builder)

	director.ConstructMethodA() // 第一种组装方法
	productA := builder.GetResult()
	productA.Show()

	director.ConstructMethodB() // 第二种组装方法
	productB := builder.GetResult()
	productB.Show()
}

// Product 产品类,表示一个复杂的对象
type Product struct {
	partA string
	partB string
	partC string
}

func (p *Product) Show() {
	fmt.Println("Product Parts:")
	fmt.Println("PartA:", p.partA)
	fmt.Println("PartB:", p.partB)
	fmt.Println("PartC:", p.partC)
}

// Builder 为创建产品对象定义抽象接口, 产品的每一个部件都有不同的组装方法
type Builder interface {
	BuildPartA1()
	BuildPartA2()
	BuildPartB1()
	BuildPartB2()
	BuildPartC1()
	BuildPartC2()
	GetResult() *Product
	Reset()
}

// ConcreteBuilder 实现了 Builder 接口,实现了具体的创建过程
type ConcreteBuilder struct {
	product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{product: &Product{}}
}

func (c *ConcreteBuilder) BuildPartA1() {
	c.product.partA = "Part A1"
}

func (c *ConcreteBuilder) BuildPartA2() {
	c.product.partA = "Part A2"
}

func (c *ConcreteBuilder) BuildPartB1() {
	c.product.partB = "Part B1"
}

func (c *ConcreteBuilder) BuildPartB2() {
	c.product.partB = "Part B2"
}

func (c *ConcreteBuilder) BuildPartC1() {
	c.product.partC = "Part C1"
}

func (c *ConcreteBuilder) BuildPartC2() {
	c.product.partC = "Part C2"
}

func (c *ConcreteBuilder) GetResult() *Product {
	return c.product
}

func (c *ConcreteBuilder) Reset() {
	c.product = &Product{}
}

// Director 构建一个使用 Builder 接口的对象
type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{builder: b}
}

// ConstructMethodA 建造的过程顺序可以随意交换或固定
func (d *Director) ConstructMethodA() {
	d.builder.Reset() // 每一次构建都是一个新的产品
	d.builder.BuildPartA1()
	d.builder.BuildPartB1()
	d.builder.BuildPartC1()
}

// ConstructMethodB 建造的过程顺序可以随意交换或固定
func (d *Director) ConstructMethodB() {
	d.builder.Reset()
	d.builder.BuildPartC2()
	d.builder.BuildPartB2()
	d.builder.BuildPartA2()
}
