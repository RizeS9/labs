package main

import "fmt"

// ==================================================
// 1. PROXY
// ==================================================

// GameServer описывает общий интерфейс игрового сервера.
type GameServer interface {
	Connect(playerName string)
}

// RealGameServer - настоящий игровой сервер.
type RealGameServer struct{}

func (s *RealGameServer) Connect(playerName string) {
	fmt.Printf("%s подключился к игровому серверу\n", playerName)
}

// GameServerProxy контролирует доступ к настоящему серверу.
type GameServerProxy struct {
	server *RealGameServer
	banned map[string]bool
}

func NewGameServerProxy() *GameServerProxy {
	return &GameServerProxy{
		server: &RealGameServer{},
		banned: map[string]bool{
			"Cheater": true,
		},
	}
}

func (p *GameServerProxy) Connect(playerName string) {
	if p.banned[playerName] {
		fmt.Printf("Доступ запрещён: игрок %s заблокирован\n", playerName)
		return
	}

	fmt.Println("Proxy: проверка доступа выполнена")
	p.server.Connect(playerName)
}

func TestProxy() {
	fmt.Println("=== Proxy ===")

	server := NewGameServerProxy()

	server.Connect("Archer")
	server.Connect("Cheater")
}

// ==================================================
// 2. BRIDGE
// ==================================================

// Weapon - реализация, которую можно использовать
// независимо от типа персонажа.
type Weapon interface {
	Use()
}

// Sword - конкретная реализация оружия.
type Sword struct{}

func (s *Sword) Use() {
	fmt.Println("атака мечом")
}

// Bow - конкретная реализация оружия.
type Bow struct{}

func (b *Bow) Use() {
	fmt.Println("выстрел из лука")
}

// Hero - абстракция, которая содержит ссылку
// на реализацию Weapon.
type Hero struct {
	name   string
	weapon Weapon
}

func NewHero(name string, weapon Weapon) *Hero {
	return &Hero{
		name:   name,
		weapon: weapon,
	}
}

func (h *Hero) Attack() {
	fmt.Printf("%s: ", h.name)
	h.weapon.Use()
}

func (h *Hero) SetWeapon(weapon Weapon) {
	h.weapon = weapon
}

func TestBridge() {
	fmt.Println("\n=== Bridge ===")

	hero := NewHero("Hero", &Sword{})
	hero.Attack()

	hero.SetWeapon(&Bow{})
	hero.Attack()
}

// ==================================================
// 3. ADAPTER
// ==================================================

// Attacker - интерфейс, который ожидает наша программа.
type Attacker interface {
	Attack()
}

// Archer - обычный объект, совместимый с Attacker.
type Archer struct{}

func (a *Archer) Attack() {
	fmt.Println("Лучник атакует")
}

// OldMage - старый класс с несовместимым интерфейсом.
// Вместо Attack у него есть CastSpell.
type OldMage struct{}

func (m *OldMage) CastSpell() {
	fmt.Println("Старый маг использует заклинание")
}

// MageAdapter адаптирует OldMage к интерфейсу Attacker.
type MageAdapter struct {
	mage *OldMage
}

func NewMageAdapter(mage *OldMage) *MageAdapter {
	return &MageAdapter{
		mage: mage,
	}
}

func (a *MageAdapter) Attack() {
	a.mage.CastSpell()
}

// PerformAttack работает с любым объектом,
// который реализует интерфейс Attacker.
func PerformAttack(attacker Attacker) {
	attacker.Attack()
}

func TestAdapter() {
	fmt.Println("\n=== Adapter ===")

	archer := &Archer{}
	PerformAttack(archer)

	oldMage := &OldMage{}
	adapter := NewMageAdapter(oldMage)

	PerformAttack(adapter)
}

func main() {
	TestProxy()
	TestBridge()
	TestAdapter()
}
