// THIS CODE IS GENERATED - DO NOT CHANGE IT (source: utility/definitions/predicates_as_options.yaml)
package dmutil

type EQBoolAsOption struct {
	EqEnabled *bool `json:"eq_enabled"`
	EQ        bool  `json:"eq"`
}

type EQBoolAsConstOption struct {
	EqEnabled *bool `json:"eq_enabled"`
	EQ        bool  `json:"eq"`
}

type EQOptionalBoolAsOption EQOptionalBool

type EQOptionalBoolAsConstOption struct {
	EqEnabled *bool `json:"eq_enabled"`
	EQ        *bool `json:"eq"`
}

type EQStringAsOption struct {
	EqEnabled *bool  `json:"eq_enabled"`
	EQ        string `json:"eq"`
}

type EQOptionalStringAsOption EQOptionalString
