package circular

type Circular[T any] struct {
	items []T
	start int
}

func (buffer *Circular[T]) Push(value T) {
	buffer.start++
	if buffer.start == len(buffer.items) {
		buffer.start = 0
	}
	buffer.items[buffer.start] = value
}

func (buffer *Circular[T]) Do(function func(*T)) {
	for f := buffer.start; f >= 0; f-- {
		function(&buffer.items[f])
	}
	for f := len(buffer.items) - 1; f > buffer.start; f-- {
		function(&buffer.items[f])
	}
}

func (buffer *Circular[T]) DoReverse(function func(*T)) {
	for f := buffer.start + 1; f < len(buffer.items); f++ {
		function(&buffer.items[f])
	}
	for f := 0; f <= buffer.start; f++ {
		function(&buffer.items[f])
	}
}

func New[T any](size int, initialValue T) Circular[T] {
	items := make([]T, size)
	for f := 0; f < size; f++ {
		items[f] = initialValue
	}

	return Circular[T]{
		items: items,
		start: 0,
	}
}
