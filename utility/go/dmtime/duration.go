package dmtime

import (
	"fmt"
	"math"
	"time"

	"github.com/Deimvis/go-ext/go1.25/xcheck/xmust"
	"github.com/Deimvis/go-ext/go1.25/xcheck/xshould"
	"github.com/Deimvis/go-ext/go1.25/xptr"
)

func (dur Duration) Duration() time.Duration {
	if dur.d == nil {
		d := xmust.Do(dur.calcDuration())
		dur.d = &d
	}
	return *dur.d
}

func (dur Duration) CloneToConst() ConstDuration {
	return ConstDuration{d: dur.Duration()}
}

func (cdur ConstDuration) Duration() time.Duration {
	return cdur.d
}

func (cdur ConstDuration) CloneToMutable() Duration {
	remain := uint64(cdur.d.Nanoseconds())
	ns := remain % 1e3
	remain /= 1e3
	mcs := remain % 1e3
	remain /= 1e3
	ms := remain % 1e3
	remain /= 1e3
	s := remain % 60
	remain /= 60
	m := remain % 60
	remain /= 60
	h := remain % 24
	remain /= 24
	d := remain

	dur_ := Duration{
		D:   xptr.T(d),
		H:   xptr.T(h),
		M:   xptr.T(m),
		S:   xptr.T(s),
		Ms:  xptr.T(ms),
		Mcs: xptr.T(mcs),
		Ns:  xptr.T(ns),
	}
	return dur_
}

func (dur Duration) calcDuration() (time.Duration, error) {
	var err error = nil
	v := func(vptr *uint64) int64 {
		if vptr == nil {
			return 0
		}
		if *vptr > math.MaxInt64 {
			err = fmt.Errorf("uint64: overflowed int64 (%v)", *vptr)
		}
		return int64(*vptr)
	}
	rawDur := time.Duration(v(dur.D)*24)*time.Hour +
		time.Duration(v(dur.H))*time.Hour +
		time.Duration(v(dur.M))*time.Minute +
		time.Duration(v(dur.S))*time.Second +
		time.Duration(v(dur.Ms))*time.Millisecond +
		time.Duration(v(dur.Mcs))*time.Microsecond +
		time.Duration(v(dur.Ns))*time.Nanosecond
	return rawDur, err
}

func (dur Duration) ValidateSelf() error {
	// TODO: implement validation tag for that and disable by default
	return xshould.Lt(0, dur.Duration(), "zero duration")
}

func (cdur ConstDuration) ValidateSelf() error {
	// TODO: implement validation tag for that and disable by default
	return xshould.Lt(0, cdur.Duration(), "zero duration")
}
