// THIS CODE IS GENERATED - DO NOT CHANGE IT (source: utility/definitions/predicates.yaml)
package dmutil

type EQBool struct {
	EQ bool `json:"eq"`
}

type EQOptionalBool struct {
	EqEnabled *bool `json:"eq_enabled"`
	EQ        *bool `json:"eq"`
}

type EQInt64 struct {
	EQ int64 `json:"eq"`
}

type EQOptionalInt64 struct {
	EqEnabled *bool  `json:"eq_enabled"`
	EQ        *int64 `json:"eq"`
}

type LTInt64 struct {
	Lt int64 `json:"lt"`
}

type LTOptionalInt64 struct {
	LtEnabled *bool  `json:"lt_enabled"`
	Lt        *int64 `json:"lt"`
}

type GTInt64 struct {
	Gt int64 `json:"gt"`
}

type GTOptionalInt64 struct {
	GtEnabled *bool  `json:"gt_enabled"`
	Gt        *int64 `json:"gt"`
}

type LEInt64 struct {
	LE int64 `json:"le"`
}

type GEInt64 struct {
	GE int64 `json:"ge"`
}

type EQString struct {
	EQ string `json:"eq"`
}

type EQOptionalString struct {
	EqEnabled *bool   `json:"eq_enabled"`
	EQ        *string `json:"eq"`
}

type AnyOfString struct {
	AnyOfEnabled *bool    `json:"any_of_enabled"`
	AnyOf        []string `json:"any_of"`
}

type AnyOfOptionalString struct {
	AnyOfEnabled *bool    `json:"any_of_enabled"`
	AnyOf        []string `json:"any_of"`
}

type RegexpFullmatch struct {
	RegexpFullmatch *string `json:"regexp_fullmatch"`
}
