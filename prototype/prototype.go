package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

/*
原型模式（Prototype Pattern）是一种创建型设计模式，通过复制现有的对象来创建新的对象，而不是通过实例化类。这个模式可以用来避免创建对象时的复杂过程，提高性能，尤其是在创建对象代价很高的情况下。

原型模式的主要特点：

	1.	克隆：通过复制现有对象来创建新对象，而不必依赖于类的实例化。
	2.	性能：在需要大量创建对象时，通过克隆来避免重复创建，可以提高性能。
	3.	动态：可以在运行时动态地改变克隆对象的结构和状态。

主要参与者：

	1.	Prototype：声明一个克隆自身的接口。
	2.	ConcretePrototype：实现克隆接口，提供克隆自身的功能。
	3.	Client：通过调用原型对象的克隆方法来创建新的对象。

注意: 本列使用 encoding/gob 来实现深拷贝
*/

func main() {
	p1 := &Person{
		Name:    "小明",
		Age:     18,
		Address: &Address{City: "加尼福尼亚"},
	}

	p2 := p1.Clone()
	p2.SetAddress("华盛顿")

	fmt.Printf("p1 %+v, addr : %v\n", p1, p1.Address.City)
	fmt.Printf("p2 %+v, addr : %v\n", p2, p2.Address.City)
}

// Address 是引用类型
type Address struct {
	City string
}

// Person 是需要实现深拷贝的类型
type Person struct {
	Name    string
	Age     int
	Address *Address
}

func (p *Person) SetAddress(addr string) {
	p.Address.City = addr
}

// Clone 方法使用 gob 包实现深拷贝
func (p *Person) Clone() *Person {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	// 序列化对象
	if err := enc.Encode(p); err != nil {
		panic(err)
	}

	// 反序列化到新对象
	var newPerson Person
	if err := dec.Decode(&newPerson); err != nil {
		panic(err)
	}

	return &newPerson
}
