package dmutil

type AsOption[T any] struct {
	Option `yaml:",inline`
	value  T `yaml:",inline"`
}

func (opt AsOption[T]) Value() T {
	return opt.value
}
