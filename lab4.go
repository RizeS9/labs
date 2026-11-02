package main

import "fmt"

// ==================================================
// 1. STRATEGY
// ==================================================

// AttackStrategy описывает стратегию атаки.
type AttackStrategy interface {
	Attack()
}

// SwordAttack - атака мечом.
type SwordAttack struct{}

func (s *SwordAttack) Attack() {
	fmt.Println("Персонаж атакует мечом")
}

// BowAttack - атака из лука.
type BowAttack struct{}

func (b *BowAttack) Attack() {
	fmt.Println("Персонаж стреляет из лука")
}

// MagicAttack - магическая атака.
type MagicAttack struct{}

func (m *MagicAttack) Attack() {
	fmt.Println("Персонаж использует магическую атаку")
}

// Character использует выбранную стратегию атаки.
type Character struct {
	name     string
	strategy AttackStrategy
}

func NewCharacter(name string, strategy AttackStrategy) *Character {
	return &Character{
		name:     name,
		strategy: strategy,
	}
}

// SetStrategy позволяет изменить стратегию во время работы программы.
func (c *Character) SetStrategy(strategy AttackStrategy) {
	c.strategy = strategy
}

func (c *Character) Attack() {
	fmt.Printf("%s: ", c.name)
	c.strategy.Attack()
}

func TestStrategy() {
	fmt.Println("=== Strategy ===")

	character := NewCharacter("Hero", &SwordAttack{})

	character.Attack()

	character.SetStrategy(&BowAttack{})
	character.Attack()

	character.SetStrategy(&MagicAttack{})
	character.Attack()
}

// ==================================================
// 2. CHAIN OF RESPONSIBILITY
// ==================================================

// Handler описывает обработчик запроса.
type Handler interface {
	SetNext(handler Handler)
	Handle(damage int)
}

// BaseHandler содержит ссылку на следующий обработчик.
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

func (b *BaseHandler) HandleNext(damage int) {
	if b.next != nil {
		b.next.Handle(damage)
	}
}

// LightArmor обрабатывает слабые атаки.
type LightArmor struct {
	BaseHandler
}

func (l *LightArmor) Handle(damage int) {
	if damage <= 20 {
		fmt.Println("Лёгкая броня заблокировала атаку")
		return
	}

	fmt.Println("Лёгкая броня не справилась")
	l.HandleNext(damage)
}

// HeavyArmor обрабатывает более сильные атаки.
type HeavyArmor struct {
	BaseHandler
}

func (h *HeavyArmor) Handle(damage int) {
	if damage <= 50 {
		fmt.Println("Тяжёлая броня заблокировала атаку")
		return
	}

	fmt.Println("Тяжёлая броня не справилась")
	h.HandleNext(damage)
}

// MagicShield обрабатывает самые сильные атаки.
type MagicShield struct {
	BaseHandler
}

func (m *MagicShield) Handle(damage int) {
	if damage <= 100 {
		fmt.Println("Магический щит заблокировал атаку")
		return
	}

	fmt.Println("Никто не смог заблокировать атаку")
}

func TestChainOfResponsibility() {
	fmt.Println("\n=== Chain of Responsibility ===")

	lightArmor := &LightArmor{}
	heavyArmor := &HeavyArmor{}
	magicShield := &MagicShield{}

	// Создаём цепочку:
	// LightArmor -> HeavyArmor -> MagicShield
	lightArmor.SetNext(heavyArmor)
	heavyArmor.SetNext(magicShield)

	fmt.Println("Урон: 10")
	lightArmor.Handle(10)

	fmt.Println("\nУрон: 40")
	lightArmor.Handle(40)

	fmt.Println("\nУрон: 80")
	lightArmor.Handle(80)

	fmt.Println("\nУрон: 150")
	lightArmor.Handle(150)
}

// ==================================================
// 3. ITERATOR
// ==================================================

// Unit - игровой юнит.
type Unit struct {
	Name string
}

// UnitCollection хранит коллекцию юнитов.
type UnitCollection struct {
	units []Unit
}

func (c *UnitCollection) Add(unit Unit) {
	c.units = append(c.units, unit)
}

// CreateIterator создаёт итератор для коллекции.
func (c *UnitCollection) CreateIterator() *UnitIterator {
	return &UnitIterator{
		collection: c,
		index:      0,
	}
}

// UnitIterator позволяет последовательно обходить коллекцию.
type UnitIterator struct {
	collection *UnitCollection
	index      int
}

// HasNext проверяет наличие следующего элемента.
func (i *UnitIterator) HasNext() bool {
	return i.index < len(i.collection.units)
}

// Next возвращает следующий элемент.
func (i *UnitIterator) Next() Unit {
	unit := i.collection.units[i.index]
	i.index++
	return unit
}

func TestIterator() {
	fmt.Println("\n=== Iterator ===")

	collection := &UnitCollection{}

	collection.Add(Unit{Name: "Archer"})
	collection.Add(Unit{Name: "Warrior"})
	collection.Add(Unit{Name: "Mage"})

	iterator := collection.CreateIterator()

	for iterator.HasNext() {
		unit := iterator.Next()
		fmt.Println("Юнит:", unit.Name)
	}
}

func main() {
	TestStrategy()
	TestChainOfResponsibility()
	TestIterator()
}
