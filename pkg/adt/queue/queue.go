package queue

type List struct {
	slice []interface{}
}

func New() List {
	slice := make([]interface{}, 0, 0)
	return List{
		slice: slice,
	}
}

func (l *List) Len() int {
	return len(l.slice)
}

func (l *List) Enqueue(value interface{}) {
	l.slice = append(l.slice, value)
}

func (l *List) Dequeue() interface{} {
	lastValue := l.slice[0]
	l.slice = l.slice[1:]
	return lastValue
}
