package ctntype

type Either[A, B any] struct {
	value any
}

func (e *Either[A, B]) SetA(a A) {
	e.value = a
}

func (e *Either[A, B]) SetB(b B) {
	e.value = b
}

func Switch[A, B, R any](e *Either[A, B], onA func(a A) R, onB func(b B) R) R {
	switch v := e.value.(type) {
	case A:
		return onA(v)
	case B:
		return onB(v)
	}
	return *new(R)
}
