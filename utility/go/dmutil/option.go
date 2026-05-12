package dmutil

// I am sorry for this name.
// Disableable wouldn't be better...
type Enableable interface {
	Enable()
	Disable()
	ConstEnableable
}

// I am sorry for this name too.
type ConstEnableable interface {
	IsEnabled() bool
}

func (opt *Option) IsEnabled() bool {
	return opt != nil && opt.Enabled != nil && *opt.Enabled
}

func (opt *ConstOption) IsEnabled() bool {
	return opt != nil && opt.Enabled != nil && *opt.Enabled
}

func NewEnabledOption() Option {
	b := true
	return Option{Enabled: &b}
}

func NewDisabledOption() Option {
	b := false
	return Option{Enabled: &b}
}
