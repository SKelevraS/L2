package chain

// Handler provides a handler interface.
type Handler interface {
	SendRequest(message int) string
}

// ConcreteHandlerA implements handler "A".
type ConcreteHandlerA struct {
	next Handler
}

// SendRequest implementation.
func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerB implements handler "B".
type ConcreteHandlerB struct {
	next Handler
}

// SendRequest implementation.
func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerC implements handler "C".
type ConcreteHandlerC struct {
	next Handler
}

// SendRequest implementation.
func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

/*
Паттерн Chain относится к поведенческим паттернам уровня объекта.

Паттерн Chain позволяет избежать привязки объекта-отправителя запроса к объекту-получателю запроса, при этом давая шанс обработать этот запрос нескольким объектам. Получатели связываются в цепочку, и запрос передается по цепочке, пока не будет обработан каким-то объектом.

По сути это цепочка обработчиков, которые по очереди получают запрос, а затем решают, обрабатывать его или нет. Если запрос не обработан, то он передается дальше по цепочке. Если же он обработан, то паттерн сам решает передавать его дальше или нет. Если запрос не обработан ни одним обработчиком, то он просто теряется.
*/