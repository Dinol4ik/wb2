package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type Handler interface {
	execute(isHandle bool)
}

type HandlerOne struct {
	next Handler
}

func (r *HandlerOne) execute(isHandle bool) {
	// logic
	if isHandle {
		r.next.execute(isHandle)
	}
}

type HandlerTwo struct {
	next Handler
}

func (h *HandlerTwo) execute(isHandle bool) {
	// logic
	if isHandle {
		h.next.execute(isHandle)
	}
}

// например:
//first := &HandlerOne{next:&HandlerTwo{}}
//first.execute(true)
