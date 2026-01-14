package internal

type Set[T ~string | int64 | int] map[T]struct{}

func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s Set[T]) Remove(item T) {
	delete(s, item)
}

func (s Set[T]) Has(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) Size() int {
	return len(map[T]struct{}(s))
}
