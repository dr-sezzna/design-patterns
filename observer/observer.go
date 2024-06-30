package main

import "fmt"

/*
观察者模式（Observer Pattern）是一种行为设计模式，它定义了一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都会得到通知并自动更新。
观察者模式用于实现对象间的解耦，使得一个对象的改变能够通知到所有相关的对象，而不需要直接依赖它们。

观察者模式的主要特点：

	1.	一对多依赖：一个被观察的对象（主题）维护一组观察者对象，当该对象的状态发生改变时，自动通知所有依赖于它的观察者对象。
	2.	解耦：观察者模式将观察者和被观察者解耦，方便各自的独立复用和维护。
	3.	动态性：观察者可以在运行时动态添加或删除，不需要修改主题对象的代码。

主要参与者：

	1.	Subject：主题，维护一组观察者对象并通知它们状态的改变。
	2.	Observer：观察者，定义一个接口，用于接收主题对象的通知。
	3.	ConcreteSubject：具体主题，实现了Subject接口，维护具体的状态。
	4.	ConcreteObserver：具体观察者，实现了Observer接口，更新自身状态以响应主题的改变。
*/

func main() {
	subject := NewConcreteSubject()

	observer1 := NewConcreteObserver("A")
	observer2 := NewConcreteObserver("B")

	// 项目添加观察者
	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)

	// 改变项目状态,自动同志观察者调用 Update
	subject.SetState("State 1")
	subject.SetState("State 2")

	// 移除一个观察者
	subject.RemoveObserver(observer1)
	subject.SetState("State 3")
}

var _ Subject = NewConcreteSubject()

// Observer 定义了接收通知的方法
type Observer interface {
	Update(string)
}

// Subject (主题) 定义了注册、注销和通知观察者的方法
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// ConcreteSubject 实现了 Subject 接口
type ConcreteSubject struct {
	observers []Observer
	state     string
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{
		observers: []Observer{},
	}
}

func (c *ConcreteSubject) RegisterObserver(observer Observer) {
	c.observers = append(c.observers, observer)
}

func (c *ConcreteSubject) RemoveObserver(observer Observer) {
	for i, o := range c.observers {
		if o == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
		}
	}
}

func (c *ConcreteSubject) NotifyObservers() {
	for _, o := range c.observers {
		o.Update(c.state)
	}
}

func (c *ConcreteSubject) SetState(state string) {
	c.state = state
	c.NotifyObservers()
}

// ConcreteObserver 实现了 Observer 接口
type ConcreteObserver struct {
	name  string
	state string
}

func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name: name}
}

func (c *ConcreteObserver) Update(state string) {
	c.state = state
	fmt.Printf("Observer %s: State updated to %s\n", c.name, c.state)
}
