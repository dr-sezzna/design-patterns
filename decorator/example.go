package main

import "fmt"

/*
在游戏开发中，装饰模式（Decorator Pattern）可以用于为游戏对象动态地添加功能，而无需修改对象的代码。这个设计模式允许行为的灵活组合，使代码更加可扩展和维护。以下是一个使用装饰模式的例子：

例子：角色属性增强系统

假设我们在开发一个角色扮演游戏（RPG），玩家的角色可以通过不同的装备来增强属性，如增加攻击力、防御力、速度等。我们可以使用装饰模式来实现这个功能。
*/

/*
基本结构
	1.	角色接口（Character Interface）：
		•	定义角色的基本功能，如获取攻击力、获取防御力等。
*/

// Character 角色接口
type Character interface {
	GetAttack() int
	GetDefense() int
}

/*
	2.	具体角色类（Concrete Character Class）：
	•	实现基本角色功能。
*/
// BasicCharacter 具体角色类
type BasicCharacter struct {
	attack  int
	defense int
}

func (c *BasicCharacter) GetAttack() int {
	return c.attack
}

func (c *BasicCharacter) GetDefense() int {
	return c.defense
}

/*
	3.	装饰器基类（Decorator Base Class）：
	•	持有一个角色对象，并实现角色接口，具体装饰器类会继承该类。
*/
// CharacterDecorator 装饰器基类
type CharacterDecorator struct {
	character Character
}

func (d *CharacterDecorator) GetAttack() int {
	return d.character.GetAttack()
}

func (d *CharacterDecorator) GetDefense() int {
	return d.character.GetDefense()
}

/*
	4.	具体装饰器类（Concrete Decorator Classes）：
	•	为角色动态添加功能。
*/
// SwordDecorator 具体装饰器类，增加攻击力
type SwordDecorator struct {
	CharacterDecorator
}

func (d *SwordDecorator) GetAttack() int {
	return d.character.GetAttack() + 5
}

// ShieldDecorator 具体装饰器类，增加防御力
type ShieldDecorator struct {
	CharacterDecorator
}

func (d *ShieldDecorator) GetDefense() int {
	return d.character.GetDefense() + 3
}

/*
使用装饰器,通过装饰器模式，我们可以动态地为角色添加功能
*/

func main() {
	// 创建一个基础角色
	basicCharacter := &BasicCharacter{attack: 10, defense: 5}
	fmt.Printf("Basic Attack: %d, Basic Defense: %d\n", basicCharacter.GetAttack(), basicCharacter.GetDefense())

	// 为角色添加一把剑
	swordCharacter := &SwordDecorator{CharacterDecorator{character: basicCharacter}}
	fmt.Printf("Attack with Sword: %d, Defense with Sword: %d\n", swordCharacter.GetAttack(), swordCharacter.GetDefense())

	// 为角色添加一个盾牌
	shieldCharacter := &ShieldDecorator{CharacterDecorator{character: swordCharacter}}
	fmt.Printf("Attack with Sword and Shield: %d, Defense with Sword and Shield: %d\n", shieldCharacter.GetAttack(), shieldCharacter.GetDefense())
}
