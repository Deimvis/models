package dmutil

type pagination interface {
	applyRules() paginationApplyRules
}

func Paginate[T any, S ~[]T](s S, p pagination) S {
	sz := int64(len(s))
	rules := p.applyRules()
	return s[min(rules.begin, sz):min(rules.end, sz)]
}

type paginationApplyRules struct {
	begin int64
	end   int64
}

func (olp OffsetLimitPagination) applyRules() paginationApplyRules {
	return paginationApplyRules{begin: olp.Offset, end: olp.Offset + olp.Limit}
}

// func Paginate()
