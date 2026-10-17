package main

import "fmt"

// Attacker - интерфейс для объектов, которые могут атаковать юнитов.
type Attacker interface {
	Attack(unit *Unit)
}

// Moveable - интерфейс для объектов, которые могут перемещаться.
type Moveable interface {
	Move(x, y int)
}

// GameObject - базовый тип для всех игровых объектов.
type GameObject struct {
	id   int
	name string
	x    int
	y    int
}

func NewGameObject(id int, name string, x, y int) GameObject {
	return GameObject{
		id:   id,
		name: name,
		x:    x,
		y:    y,
	}
}

func (g *GameObject) GetId() int {
	return g.id
}

func (g *GameObject) GetName() string {
	return g.name
}

func (g *GameObject) GetX() int {
	return g.x
}

func (g *GameObject) GetY() int {
	return g.y
}

// Unit - базовый тип для управляемых юнитов.
type Unit struct {
	GameObject
	hp float32
}

func NewUnit(id int, name string, x, y int, hp float32) Unit {
	return Unit{
		GameObject: NewGameObject(id, name, x, y),
		hp:         hp,
	}
}

func (u *Unit) IsAlive() bool {
	return u.hp > 0
}

func (u *Unit) GetHp() float32 {
	return u.hp
}

func (u *Unit) ReceiveDamage(damage float32) {
	if damage <= 0 {
		return
	}

	u.hp -= damage

	if u.hp < 0 {
		u.hp = 0
	}
}

// Building - базовый тип для построек.
type Building struct {
	GameObject
	built bool
}

func NewBuilding(id int, name string, x, y int, built bool) Building {
	return Building{
		GameObject: NewGameObject(id, name, x, y),
		built:      built,
	}
}

func (b *Building) IsBuilt() bool {
	return b.built
}

// Archer наследует Unit.
// Реализует интерфейсы Attacker и Moveable.
type Archer struct {
	Unit
	damage float32
}

func NewArcher(
	id int,
	name string,
	x, y int,
	hp float32,
	damage float32,
) *Archer {
	return &Archer{
		Unit:   NewUnit(id, name, x, y, hp),
		damage: damage,
	}
}

func (a *Archer) Attack(unit *Unit) {
	if !a.IsAlive() {
		fmt.Printf("%s не может атаковать, потому что он мёртв\n", a.GetName())
		return
	}

	if !unit.IsAlive() {
		fmt.Printf("%s уже мёртв\n", unit.GetName())
		return
	}

	fmt.Printf(
		"%s атакует %s и наносит %.1f урона\n",
		a.GetName(),
		unit.GetName(),
		a.damage,
	)

	unit.ReceiveDamage(a.damage)
}

func (a *Archer) Move(x, y int) {
	if !a.IsAlive() {
		fmt.Printf("%s не может двигаться, потому что он мёртв\n", a.GetName())
		return
	}

	a.x = x
	a.y = y

	fmt.Printf(
		"%s переместился в точку (%d, %d)\n",
		a.GetName(),
		x,
		y,
	)
}

// Fort наследует Building.
// Реализует интерфейс Attacker.
type Fort struct {
	Building
	damage float32
}

func NewFort(
	id int,
	name string,
	x, y int,
	built bool,
	damage float32,
) *Fort {
	return &Fort{
		Building: NewBuilding(id, name, x, y, built),
		damage:   damage,
	}
}

func (f *Fort) Attack(unit *Unit) {
	if !f.IsBuilt() {
		fmt.Printf(
			"%s не может атаковать, потому что крепость ещё не построена\n",
			f.GetName(),
		)
		return
	}

	if !unit.IsAlive() {
		fmt.Printf("%s уже мёртв\n", unit.GetName())
		return
	}

	fmt.Printf(
		"%s атакует %s и наносит %.1f урона\n",
		f.GetName(),
		unit.GetName(),
		f.damage,
	)

	unit.ReceiveDamage(f.damage)
}

// MobileHome наследует Building.
// Реализует интерфейс Moveable.
type MobileHome struct {
	Building
}

func NewMobileHome(
	id int,
	name string,
	x, y int,
	built bool,
) *MobileHome {
	return &MobileHome{
		Building: NewBuilding(id, name, x, y, built),
	}
}

func (m *MobileHome) Move(x, y int) {
	if !m.IsBuilt() {
		fmt.Printf(
			"%s не может двигаться, потому что он ещё не построен\n",
			m.GetName(),
		)
		return
	}

	m.x = x
	m.y = y

	fmt.Printf(
		"%s переместился в точку (%d, %d)\n",
		m.GetName(),
		x,
		y,
	)
}

func main() {
	archer := NewArcher(
		1,
		"Archer",
		0,
		0,
		100,
		25,
	)

	enemy := NewArcher(
		2,
		"Enemy Archer",
		10,
		10,
		80,
		20,
	)

	fort := NewFort(
		3,
		"Fort",
		5,
		5,
		true,
		40,
	)

	mobileHome := NewMobileHome(
		4,
		"Mobile Home",
		2,
		2,
		true,
	)

	fmt.Println("=== GameObject data ===")
	fmt.Println("ID:", archer.GetId())
	fmt.Println("Name:", archer.GetName())
	fmt.Println("X:", archer.GetX())
	fmt.Println("Y:", archer.GetY())

	fmt.Println("\n=== Unit data ===")
	fmt.Println("HP:", archer.GetHp())
	fmt.Println("Alive:", archer.IsAlive())

	fmt.Println("\n=== Archer attack ===")
	archer.Attack(&enemy.Unit)

	fmt.Println("Enemy HP:", enemy.GetHp())

	fmt.Println("\n=== Archer movement ===")
	archer.Move(3, 4)

	fmt.Printf(
		"Archer position: (%d, %d)\n",
		archer.GetX(),
		archer.GetY(),
	)

	fmt.Println("\n=== Fort attack ===")
	fort.Attack(&enemy.Unit)

	fmt.Println("Enemy HP:", enemy.GetHp())

	fmt.Println("\n=== MobileHome movement ===")
	mobileHome.Move(7, 8)

	fmt.Printf(
		"MobileHome position: (%d, %d)\n",
		mobileHome.GetX(),
		mobileHome.GetY(),
	)

	fmt.Println("\n=== Building state ===")
	fmt.Println("Fort built:", fort.IsBuilt())
	fmt.Println("MobileHome built:", mobileHome.IsBuilt())
}
