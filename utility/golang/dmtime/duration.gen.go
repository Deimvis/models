// THIS CODE IS GENERATED - DO NOT CHANGE IT (source: utility/definitions/time.yaml)
package dmtime

import "time"

type Duration struct {
	D   *uint64        `json:"d" validate:"gt=0"`
	H   *uint64        `json:"h" validate:"gt=0"`
	M   *uint64        `json:"m" validate:"gt=0"`
	S   *uint64        `json:"s" validate:"gt=0"`
	Ms  *uint64        `json:"ms" validate:"gt=0"`
	Mcs *uint64        `json:"mcs" validate:"gt=0"`
	Ns  *uint64        `json:"ns" validate:"gt=0"`
	d   *time.Duration // cached
}

type ConstDuration struct{ d time.Duration }
