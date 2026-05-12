package dmutil

import "github.com/Deimvis/go-ext/go1.25/xptr"

func (p EQOptionalBool) Enabled() bool {
	return p.EqEnabled != nil && *p.EqEnabled
}

func (p *EQOptionalBool) Enable() {
	p.EqEnabled = xptr.T(true)
}

func (p *EQOptionalBool) Disable() {
	p.EqEnabled = xptr.T(false)
}
