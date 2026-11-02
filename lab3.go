package main

import (
	"fmt"
	"sync"
)

// ==================================================
// 1. SINGLETON
// ==================================================

// GameManager - объект, который должен существовать
// только в одном экземпляре.
type GameManager struct {
	Name string
}

var instance *GameManager
var once sync.Once

// GetGameManager возвращает единственный экземпляр GameManager.
func GetGameManager() *GameManager {
	once.Do(func() {
		instance = &GameManager{
			Name: "Strategy Game",
		}
	})

	return instance
}

func TestSingleton() {
	fmt.Println("=== Singleton ===")

	manager1 := GetGameManager()
	manager2 := GetGameManager()

	fmt.Println("Manager 1:", manager1.Name)
	fmt.Println("Manager 2:", manager2.Name)

	// Проверяем, что это один и тот же объект.
	fmt.Println("Один объект:", manager1 == manager2)
}

// ==================================================
// 2. FACTORY METHOD
// ==================================================

// Fighter - общий интерфейс игровых юнитов.
type Fighter interface {
	Attack()
}

// ArcherUnit - лучник.
type ArcherUnit struct{}

func (a *ArcherUnit) Attack() {
	fmt.Println("Лучник стреляет из лука")
}

// WarriorUnit - воин.
type WarriorUnit struct{}

func (w *WarriorUnit) Attack() {
	fmt.Println("Воин атакует мечом")
}

// UnitFactory - интерфейс фабрики.
type UnitFactory interface {
	CreateUnit() Fighter
}

// ArcherFactory создаёт лучников.
type ArcherFactory struct{}

func (f *ArcherFactory) CreateUnit() Fighter {
	return &ArcherUnit{}
}

// WarriorFactory создаёт воинов.
type WarriorFactory struct{}

func (f *WarriorFactory) CreateUnit() Fighter {
	return &WarriorUnit{}
}

func TestFactoryMethod() {
	fmt.Println("\n=== Factory Method ===")

	var factory UnitFactory

	factory = &ArcherFactory{}
	archer := factory.CreateUnit()
	archer.Attack()

	factory = &WarriorFactory{}
	warrior := factory.CreateUnit()
	warrior.Attack()
}

// ==================================================
// 3. ABSTRACT FACTORY
// ==================================================

// Soldier - интерфейс солдата.
type Soldier interface {
	Info()
}

// Base - интерфейс базы.
type Base interface {
	Info()
}

// -------------------- Human --------------------

type HumanSoldier struct{}

func (h *HumanSoldier) Info() {
	fmt.Println("Создан солдат людей")
}

type HumanBase struct{}

func (h *HumanBase) Info() {
	fmt.Println("Создана база людей")
}

// -------------------- Orc --------------------

type OrcSoldier struct{}

func (o *OrcSoldier) Info() {
	fmt.Println("Создан солдат орков")
}

type OrcBase struct{}

func (o *OrcBase) Info() {
	fmt.Println("Создана база орков")
}

// GameFactory - абстрактная фабрика.
// Создаёт семейство связанных объектов.
type GameFactory interface {
	CreateSoldier() Soldier
	CreateBase() Base
}

// HumanFactory создаёт объекты людей.
type HumanFactory struct{}

func (f *HumanFactory) CreateSoldier() Soldier {
	return &HumanSoldier{}
}

func (f *HumanFactory) CreateBase() Base {
	return &HumanBase{}
}

// OrcFactory создаёт объекты орков.
type OrcFactory struct{}

func (f *OrcFactory) CreateSoldier() Soldier {
	return &OrcSoldier{}
}

func (f *OrcFactory) CreateBase() Base {
	return &OrcBase{}
}

func TestAbstractFactory() {
	fmt.Println("\n=== Abstract Factory ===")

	var factory GameFactory

	fmt.Println("Люди:")

	factory = &HumanFactory{}

	humanSoldier := factory.CreateSoldier()
	humanBase := factory.CreateBase()

	humanSoldier.Info()
	humanBase.Info()

	fmt.Println("Орки:")

	factory = &OrcFactory{}

	orcSoldier := factory.CreateSoldier()
	orcBase := factory.CreateBase()

	orcSoldier.Info()
	orcBase.Info()
}

// ==================================================
// 4. BUILDER
// ==================================================

// Character - сложный игровой объект.
type Character struct {
	Name   string
	Class  string
	Weapon string
	Armor  string
	HP     int
}

// CharacterBuilder - строитель персонажа.
type CharacterBuilder struct {
	character Character
}

// NewCharacterBuilder создаёт нового строителя.
func NewCharacterBuilder() *CharacterBuilder {
	return &CharacterBuilder{
		character: Character{
			HP: 100,
		},
	}
}

func (b *CharacterBuilder) SetName(name string) *CharacterBuilder {
	b.character.Name = name
	return b
}

func (b *CharacterBuilder) SetClass(class string) *CharacterBuilder {
	b.character.Class = class
	return b
}

func (b *CharacterBuilder) SetWeapon(weapon string) *CharacterBuilder {
	b.character.Weapon = weapon
	return b
}

func (b *CharacterBuilder) SetArmor(armor string) *CharacterBuilder {
	b.character.Armor = armor
	return b
}

func (b *CharacterBuilder) SetHP(hp int) *CharacterBuilder {
	b.character.HP = hp
	return b
}

// Build возвращает готового персонажа.
func (b *CharacterBuilder) Build() Character {
	return b.character
}

func TestBuilder() {
	fmt.Println("\n=== Builder ===")

	character := NewCharacterBuilder().
		SetName("Aragorn").
		SetClass("Warrior").
		SetWeapon("Sword").
		SetArmor("Heavy Armor").
		SetHP(150).
		Build()

	fmt.Println("Name:", character.Name)
	fmt.Println("Class:", character.Class)
	fmt.Println("Weapon:", character.Weapon)
	fmt.Println("Armor:", character.Armor)
	fmt.Println("HP:", character.HP)
}

func main() {
	TestSingleton()
	TestFactoryMethod()
	TestAbstractFactory()
	TestBuilder()
}
