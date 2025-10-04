package utils

type Set[T comparable] = map[T]struct{}

func Add[T comparable](s Set[T], v T) {
	s[v] = struct{}{}
}

func Contains[T comparable](s Set[T], v T) bool {
	_, ok := s[v]
	return ok
}
