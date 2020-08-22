package stack

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

func (l *List) Push(value interface{}) {
	l.slice = append(l.slice, value)
}

func (l *List) Pop() interface{} {
	lastIndex := len(l.slice) - 1
	lastValue := l.slice[lastIndex]
	l.slice = l.slice[:lastIndex]
	return lastValue
}
