package dmutil

import "github.com/Deimvis/go-ext/go1.25/xptr"

func (p *EQBoolAsOption) Enabled() bool {
	return p != nil && p.EqEnabled != nil && *p.EqEnabled
}

func (p *EQBoolAsOption) Enable() {
	p.EqEnabled = xptr.T(true)
}

func (p *EQBoolAsOption) Disable() {
	p.EqEnabled = xptr.T(false)
}

func (p *EQBoolAsConstOption) Enabled() bool {
	return p != nil && p.EqEnabled != nil && *p.EqEnabled
}

func (p *EQOptionalBoolAsOption) Enabled() bool {
	return p != nil && p.EqEnabled != nil && *p.EqEnabled
}

func (p *EQOptionalBoolAsOption) Enable() {
	p.EqEnabled = xptr.T(true)
}

func (p *EQOptionalBoolAsOption) Disable() {
	p.EqEnabled = xptr.T(false)
}

func (p *EQOptionalBoolAsConstOption) Enabled() bool {
	return p != nil && p.EqEnabled != nil && *p.EqEnabled
}

func (p *EQStringAsOption) Enabled() bool {
	return p != nil && p.EqEnabled != nil && *p.EqEnabled
}

func (p *EQStringAsOption) Enable() {
	p.EqEnabled = xptr.T(true)
}

func (p *EQStringAsOption) Disable() {
	p.EqEnabled = xptr.T(false)
}

func (p *EQOptionalStringAsOption) Enabled() bool {
	return p != nil && p.EqEnabled != nil && *p.EqEnabled
}

func (p *EQOptionalStringAsOption) Enable() {
	p.EqEnabled = xptr.T(true)
}

func (p *EQOptionalStringAsOption) Disable() {
	p.EqEnabled = xptr.T(false)
}
