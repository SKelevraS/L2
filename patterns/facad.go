package facade

import (
	"strings"
)

// NewMan creates man.
func NewMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

// Man implements man and facade.
type Man struct {
	house *House
	tree  *Tree
	child *Child
}

// Todo returns that man must do.
func (m *Man) Todo() string {
	result := []string{
		m.house.Build(),
		m.tree.Grow(),
		m.child.Born(),
	}
	return strings.Join(result, "\n")
}

// House implements a subsystem "House"
type House struct {
}

// Build implementation.
func (h *House) Build() string {
	return "Build house"
}

// Tree implements a subsystem "Tree"
type Tree struct {
}

// Grow implementation.
func (t *Tree) Grow() string {
	return "Tree grow"
}

// Child implements a subsystem "Child"
type Child struct {
}

// Born implementation.
func (c *Child) Born() string {
	return "Child born"
}

/*
Паттерн Facade относится к структурным паттернам уровня объекта.

Паттерн Facade предоставляет высокоуровневый унифицированный интерфейс в виде набора имен методов к набору взаимосвязанных классов или объектов некоторой подсистемы, что облегчает ее использование.

Разбиение сложной системы на подсистемы позволяет упростить процесс разработки, а также помогает максимально снизить зависимости одной подсистемы от другой. Однако использовать такие подсистемы становиться довольно сложно. Один из способов решения этой проблемы является паттерн Facade. Наша задача, сделать простой, единый интерфейс, через который можно было бы взаимодействовать с подсистемами.
*/