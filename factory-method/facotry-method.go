/*
工厂方法模式（Factory Method Pattern）是一种创建型设计模式，它定义了一个创建对象的接口，但由子类决定要实例化的类是哪一个。
工厂方法让一个类的实例化延迟到其子类。这样做的目的是让客户端代码与特定类的实例化解耦，便于将来扩展和维护。

工厂方法模式的主要特点：

 1. 解耦：工厂方法模式将对象的创建和使用分离，使得客户端代码不需要直接实例化对象。
 2. 扩展性：可以通过子类化工厂类来创建不同的对象，便于系统的扩展。
 3. 灵活性：允许客户端使用不同的工厂来创建对象，提供了更大的灵活性。

主要参与者：

 1. Product：定义了工厂方法所创建对象的接口。
 2. ConcreteProduct：实现了Product接口的具体类。
 3. Creator：声明了工厂方法，返回一个Product类型的对象。Creator可以是一个抽象类，也可以是一个接口。
 4. ConcreteCreator：实现了工厂方法，返回具体的Product实例。
*/
package main

import "fmt"

func main() {
	var creator Creator

	// 使用 ConcreteCreatorA 来创建 Product A
	creator = &ConcreteCreatorA{}
	productA := creator.CreateProduct()
	fmt.Println(productA.Use())

	// 使用 ConcreteCreatorB 来创建 Product B
	creator = &ConcreteCreatorB{}
	productB := creator.CreateProduct()
	fmt.Println(productB.Use())
}

// Product 接口
type Product interface {
	Use() string
}

// ConcreteProductA 结构体
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Using Product A"
}

// ConcreteProductB 结构体
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Using Product B"
}

// Creator 接口
type Creator interface {
	CreateProduct() Product
}

// ConcreteCreatorA 结构体
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteCreatorB 结构体
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{}
}
