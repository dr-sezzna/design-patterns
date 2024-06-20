package main

import "fmt"

/*
外观模式（Facade Pattern）是一种结构型设计模式，它为子系统中的一组接口提供了一个统一的接口。
外观模式定义了一个高层接口，使得这一子系统更加容易使用。
该模式旨在简化客户端与复杂子系统之间的交互，使客户端无需了解子系统的内部细节。
特别适合重构以及维护老的有很多子系统的工程项目.

外观模式的主要特点：

	1.	简化接口：通过提供一个统一的接口，简化子系统的使用。
	2.	解耦：客户端代码与子系统的实现细节解耦，降低了客户端与子系统之间的依赖。
	3.	提高灵活性：可以在不影响客户端代码的情况下更改子系统。

主要参与者：

	1.	Facade：提供一个高层接口，简化客户端对子系统的访问。
	2.	Subsystem Classes：实现子系统的功能，处理Facade对象指派的任务。客户端通过Facade与子系统进行交互。
*/

func main() {
	facade := NewFacade()
	facade.Operation1() // 集合了子系统操作的接口
	fmt.Println()
	facade.Operation2()
}

// 子系统 A
type SubSystemA struct{}

func (s *SubSystemA) OperationA() {
	fmt.Println("SubSystem A, Method A")
}

// 子系统 B
type SubSystemB struct{}

func (s *SubSystemB) OperationB() {
	fmt.Println("SubSystem B, Method B")
}

// 子系统 C
type SubSystemC struct{}

func (s *SubSystemC) OperationC() {
	fmt.Println("SubSystem C, Method C")
}

// Facade 提供了一个统一的接口来访问子系统
type Facade struct {
	subsystemA *SubSystemA
	subsystemB *SubSystemB
	subsystemC *SubSystemC
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubSystemA{},
		subsystemB: &SubSystemB{},
		subsystemC: &SubSystemC{},
	}
}

func (f *Facade) Operation1() {
	fmt.Println("Facade Operation 1:")
	f.subsystemA.OperationA()
	f.subsystemB.OperationB()
}

func (f *Facade) Operation2() {
	fmt.Println("Facade Operation 2:")
	f.subsystemB.OperationB()
	f.subsystemC.OperationC()
}
