package main

import "fmt"

/*
代理模式（Proxy Pattern）是一种结构型设计模式，它为其他对象提供一种代理以控制对这个对象的访问。代理对象通常会在客户端和目标对象之间起到一个中介的作用，可以在不改变目标对象的前提下扩展其功能。

代理模式的主要特点：

	1.	控制访问：代理可以控制对目标对象的访问，可以在访问前后进行预处理或后处理操作。
	2.	延迟实例化：代理可以延迟实例化目标对象，直到需要时才进行实例化，这对于资源密集型对象非常有用。
	3.	日志记录或缓存：代理可以记录访问日志或者实现缓存机制，以提高系统性能。

主要参与者：

	1.	Subject：定义了目标对象和代理对象的共同接口。
	2.	RealSubject：实现了Subject接口的具体对象，是真正的业务对象。
	3.	Proxy：实现了Subject接口，并包含对RealSubject的引用，可以控制对RealSubject的访问。
*/

func main() {
	proxy := &Proxy{}
	fmt.Println(proxy.Request())
}

// Subject 接口
type Subject interface {
	Request() string
}

// RealSubject 结构体
type RealSubject struct{}

func (r *RealSubject) Request() string {
	return "RealSubject: Handling request."
}

// Proxy 结构体
type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) Request() string {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	// 可以在这里添加额外的操作，如权限检查、日志记录等
	fmt.Println("Proxy: Logging the request.")
	return p.realSubject.Request()
}
