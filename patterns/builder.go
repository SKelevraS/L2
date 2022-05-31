package builder

type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

// Director implements a manager
type Director struct {
	builder Builder
}

// Construct tells the builder what to do and in what order.
func (d *Director) Construct() {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

// ConcreteBuilder implements Builder interface.
type ConcreteBuilder struct {
	product *Product
}

// MakeHeader builds a header of document..
func (b *ConcreteBuilder) MakeHeader(str string) {
	b.product.Content += "<header>" + str + "</header>"
}

// MakeBody builds a body of document.
func (b *ConcreteBuilder) MakeBody(str string) {
	b.product.Content += "<article>" + str + "</article>"
}

// MakeFooter builds a footer of document.
func (b *ConcreteBuilder) MakeFooter(str string) {
	b.product.Content += "<footer>" + str + "</footer>"
}

// Product implementation.
type Product struct {
	Content string
}

// Show returns product.
func (p *Product) Show() string {
	return p.Content
}
/*
Паттерн Builder относится к порождающим паттернам уровня объекта.
Он определяет процесс поэтапного построения сложного продукта. После того как будет построена последняя его часть, продукт можно использовать.
*/