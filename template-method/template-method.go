package main

import "fmt"

/*
模板方法模式（Template Method Pattern）是一种行为设计模式，它在基类中定义一个算法的骨架，并允许子类在不改变算法结构的情况下重新定义算法的某些步骤。
模板方法模式通过将不变的行为提取到基类中，减少代码重复，并提供一个统一的算法框架。

模板方法模式的主要特点：

	1.	算法框架：在基类中定义一个算法的骨架，将算法的具体步骤延迟到子类中实现。
	2.	代码复用：通过在基类中实现不变的部分，减少代码重复，提高代码复用性。
	3.	灵活性：允许子类在不改变算法结构的情况下重新定义算法的某些步骤，实现算法的不同变体。

主要参与者：

	1.	AbstractClass：定义算法骨架，并包含一个或多个模板方法。
	2.	ConcreteClass：实现 AbstractClass 中的抽象方法，完成具体步骤的定义。
*/

func main() {
	var concreteClassA *BaseClass = &BaseClass{&ConcreteClassA{}}
	var concreteClassB *BaseClass = &BaseClass{&ConcreteClassB{}}

	fmt.Println("Executing TemplateMethod for ConcreteClassA:")
	concreteClassA.TemplateMethod()

	fmt.Println("\nExecuting TemplateMethod for ConcreteClassB:")
	concreteClassB.TemplateMethod()
}

// AbstractClass 定义了算法的骨架步骤
type AbstractClass interface {
	Step1()
	Step2()
	Step3()
}

// BaseClass 提供了模板方法的默认实现
type BaseClass struct {
	AbstractClass
}

// TemplateMethod 模版方法实现了具体的逻辑骨架,子步骤的执行顺序等
func (b *BaseClass) TemplateMethod() {
	b.Step1()
	b.Step2()
	b.Step3()
}

// ConcreteClassA 实现了具体步骤
type ConcreteClassA struct {
}

func (c *ConcreteClassA) Step1() {
	fmt.Println("ConcreteClassA: Step 1")
}

func (c *ConcreteClassA) Step2() {
	fmt.Println("ConcreteClassA: Step 2")
}

func (c *ConcreteClassA) Step3() {
	fmt.Println("ConcreteClassA: Step 3")
}

// ConcreteClassB 实现了具体步骤
type ConcreteClassB struct {
}

func (c *ConcreteClassB) Step1() {
	fmt.Println("ConcreteClassB: Step 1")
}

func (c *ConcreteClassB) Step2() {
	fmt.Println("ConcreteClassB: Step 2")
}

func (c *ConcreteClassB) Step3() {
	fmt.Println("ConcreteClassB: Step 3")
}
