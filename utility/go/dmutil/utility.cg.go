// THIS CODE IS GENERATED - DO NOT CHANGE IT (source: utility/definitions/utility.yaml)
package dmutil

type UpdateCommand struct {
	Enabled *bool `json:"_enabled"`
}

type SortOrder string

type SearchResultPagination struct {
	TotalCount int64 `json:"total_count"`
}

type LimitPagination struct {
	Limit int64 `json:"limit"`
}

type OffsetLimitPagination struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type SortKey struct {
	Index int64     `json:"index" validate:"gte=0"`
	Order SortOrder `json:"order"  validate:"oneof=asc desc"`
}

type SearchOptionContainsAllString struct {
	ContainsAll []string `json:"contains_all"`
}

type SearchOptionAnyOfString struct {
	AnyOf []string `json:"any_of"`
}

type SearchOptionEQBool struct {
	EQ *bool `json:"eq"`
}

type SearchOptionEQString struct {
	EQ *string `json:"eq"`
}

type SearchOptionGE struct {
	GE *int64 `json:"ge"`
}

type SearchOptionLE struct {
	LE *int64 `json:"le"`
}

type SearchOptionGELE struct {
	SearchOptionGE
	SearchOptionLE
}

var (
	SO_Asc  SortOrder = "asc"
	SO_Desc SortOrder = "desc"
)
