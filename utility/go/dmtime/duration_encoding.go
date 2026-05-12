package dmtime

import (
	"encoding/json"
	"fmt"
	"math"

	yaml_v3 "gopkg.in/yaml.v3"
)

func (cdur ConstDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(cdur.CloneToMutable())
}

func (cdur *ConstDuration) UnmarshalJSON(b []byte) error {
	var dur Duration
	err := json.Unmarshal(b, &dur)
	if err != nil {
		return err
	}
	*cdur = dur.CloneToConst()
	return nil
}

func (cdur ConstDuration) MarshalYAML() (any, error) {
	return cdur.CloneToMutable(), nil
}

func (cdur *ConstDuration) UnmarshalYAML(node *yaml_v3.Node) error {
	var dur Duration
	err := node.Decode(&dur)
	if err != nil {
		return err
	}
	*cdur = dur.CloneToConst()
	return nil
}

func safeFloor(v float64) (int64, error) {
	vf := math.Floor(v)
	if vf <= math.MinInt64 || math.MaxInt64 <= vf {
		return 0, fmt.Errorf("float64: overflowed int64 (%f)", vf)
	}
	return int64(vf), nil
}
