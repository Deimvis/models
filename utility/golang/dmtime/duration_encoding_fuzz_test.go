package dmtime

import (
	"encoding/json"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	yaml_v3 "gopkg.in/yaml.v3"
)

func FuzzConstDuration_Encoding_Json(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(1))
	f.Add(int64(59))
	f.Add(int64(60))
	f.Add(int64(61))
	f.Add(int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, d int64) {
		if d < 0 {
			t.Skip("negative duration")
		}
		dur := ConstDuration{d: time.Duration(d)}
		data, err := json.Marshal(dur)
		require.NoError(t, err)
		var durAct ConstDuration
		err = json.Unmarshal(data, &durAct)
		require.NoError(t, err)
		require.Equal(t, dur, durAct)
	})
}

func FuzzConstDuration_Encoding_YamlV3(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(1))
	f.Add(int64(59))
	f.Add(int64(60))
	f.Add(int64(61))
	f.Add(int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, d int64) {
		if d < 0 {
			t.Skip("negative duration")
		}
		dur := ConstDuration{d: time.Duration(d)}
		data, err := yaml_v3.Marshal(dur)
		require.NoError(t, err)
		var durAct ConstDuration
		err = yaml_v3.Unmarshal(data, &durAct)
		require.NoError(t, err)
		require.Equal(t, dur, durAct)
	})
}
