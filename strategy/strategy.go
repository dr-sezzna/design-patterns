package main

import "fmt"

/*
策略模式（Strategy Pattern）是一种行为设计模式，它定义了一系列算法，并将每一个算法封装起来，使得它们可以互相替换。策略模式使得算法可以独立于使用它的客户端变化，从而实现了算法的解耦和灵活性。

策略模式的主要特点：

	1.	算法封装：将每一个算法封装到独立的策略类中，使得算法可以互相替换。
	2.	独立变化：算法可以独立于客户端变化，客户端可以动态选择和切换不同的算法。
	3.	开闭原则：通过引入新的策略类，可以在不修改现有代码的情况下扩展新的算法。

主要参与者：

	1.	Strategy：策略接口，定义了算法的接口。
	2.	ConcreteStrategy：具体策略类，实现了策略接口的具体算法。
	3.	Context：上下文类，维护对策略对象的引用，提供给客户端使用的接口。
*/

// Strategy 定义了算法的接口
type Strategy interface {
	Execute(a, b int) int
}

// ConcreteStrategyAdd 实现了加法策略
type ConcreteStrategyAdd struct {
}

func (c *ConcreteStrategyAdd) Execute(a, b int) int {
	return a + b
}

// ConcreteStrategySubtract 实现了减法策略
type ConcreteStrategySubtract struct {
}

func (c *ConcreteStrategySubtract) Execute(a, b int) int {
	return a - b
}

// ConcreteStrategyMultiply 实现了乘法策略
type ConcreteStrategyMultiply struct {
}

func (c *ConcreteStrategyMultiply) Execute(a, b int) int {
	return a * b
}

// Context 上下文类，维护对策略对象的引用
type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

func main() {
	context := NewContext(&ConcreteStrategyAdd{})
	fmt.Println("Add Strategy: 3 + 4 =", context.ExecuteStrategy(3, 4))

	context.SetStrategy(&ConcreteStrategySubtract{})
	fmt.Println("Subtract Strategy: 3 - 4 =", context.ExecuteStrategy(3, 4))

	context.SetStrategy(&ConcreteStrategyMultiply{})
	fmt.Println("Multiply Strategy: 3 * 4 =", context.ExecuteStrategy(3, 4))
}
