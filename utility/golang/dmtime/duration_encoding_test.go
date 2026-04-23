package dmtime

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	yaml_v3 "gopkg.in/yaml.v3"
)

func TestConstDuration_Encoding_JsonMarshal(t *testing.T) {
	tcs := []struct {
		title string
		d     time.Duration
		exp   string
	}{
		{
			"5ns",
			5 * time.Nanosecond,
			`{"d":0, "h":0, "m":0, "s": 0, "ms":0, "mcs":0, "ns":5}`,
		},
		{
			"5mcs",
			5 * time.Microsecond,
			`{"d":0, "h":0, "m":0, "s": 0, "ms":0, "mcs":5, "ns":0}`,
		},
		{
			"5ms",
			5 * time.Millisecond,
			`{"d":0, "h":0, "m":0, "s": 0, "ms":5, "mcs":0, "ns":0}`,
		},
		{
			"5s",
			5 * time.Second,
			`{"d":0, "h":0, "m":0, "s": 5, "ms":0, "mcs":0, "ns":0}`,
		},
		{
			"5m",
			5 * time.Minute,
			`{"d":0, "h":0, "m":5, "s": 0, "ms":0, "mcs":0, "ns":0}`,
		},
		{
			"5h",
			5 * time.Hour,
			`{"d":0, "h":5, "m":0, "s": 0, "ms":0, "mcs":0, "ns":0}`,
		},
		{
			"5d",
			5 * 24 * time.Hour,
			`{"d":5, "h":0, "m":0, "s": 0, "ms":0, "mcs":0, "ns":0}`,
		},
		{
			"59s",
			59 * time.Second,
			`{"d":0, "h":0, "m":0, "s": 59, "ms":0, "mcs":0, "ns":0}`,
		},
		{
			"60s",
			60 * time.Second,
			`{"d":0, "h":0, "m":1, "s": 0, "ms":0, "mcs":0, "ns":0}`,
		},
		{
			"61s",
			61 * time.Second,
			`{"d":0, "h":0, "m":1, "s": 1, "ms":0, "mcs":0, "ns":0}`,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.title, func(t *testing.T) {
			dur := ConstDuration{d: tc.d}
			data, err := json.Marshal(dur)
			require.NoError(t, err)
			require.JSONEq(t, tc.exp, string(data))
		})
	}
}

func TestConstDuration_Encoding_JsonUnmarshal(t *testing.T) {
	var dur ConstDuration
	data := `{"d":0, "h":0, "m":1, "ms":0, "ns":0, "mcs":0, "s":0}`
	err := json.Unmarshal([]byte(data), &dur)
	require.NoError(t, err)
	require.Equal(t, time.Minute, dur.d)
}

func TestConstDuration_Encoding_YamlV3(t *testing.T) {
	tcs := []struct {
		title string
		d     time.Duration
		data  string
	}{
		{
			"5ns",
			5 * time.Nanosecond,
			"d: 0\nh: 0\nm: 0\ns: 0\nms: 0\nmcs: 0\nns: 5\n",
		},
		{
			"5mcs",
			5 * time.Microsecond,
			"d: 0\nh: 0\nm: 0\ns: 0\nms: 0\nmcs: 5\nns: 0\n",
		},
		{
			"5ms",
			5 * time.Millisecond,
			"d: 0\nh: 0\nm: 0\ns: 0\nms: 5\nmcs: 0\nns: 0\n",
		},
		{
			"5s",
			5 * time.Second,
			"d: 0\nh: 0\nm: 0\ns: 5\nms: 0\nmcs: 0\nns: 0\n",
		},
		{
			"5m",
			5 * time.Minute,
			"d: 0\nh: 0\nm: 5\ns: 0\nms: 0\nmcs: 0\nns: 0\n",
		},
		{
			"5h",
			5 * time.Hour,
			"d: 0\nh: 5\nm: 0\ns: 0\nms: 0\nmcs: 0\nns: 0\n",
		},
		{
			"5d",
			5 * 24 * time.Hour,
			"d: 5\nh: 0\nm: 0\ns: 0\nms: 0\nmcs: 0\nns: 0\n",
		},
		{
			"59s",
			59 * time.Second,
			"d: 0\nh: 0\nm: 0\ns: 59\nms: 0\nmcs: 0\nns: 0\n",
		},
		{
			"60s",
			60 * time.Second,
			"d: 0\nh: 0\nm: 1\ns: 0\nms: 0\nmcs: 0\nns: 0\n",
		},
		{
			"61s",
			61 * time.Second,
			"d: 0\nh: 0\nm: 1\ns: 1\nms: 0\nmcs: 0\nns: 0\n",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.title, func(t *testing.T) {
			t.Run("Marshal", func(t *testing.T) {
				dur := ConstDuration{d: tc.d}
				data, err := yaml_v3.Marshal(dur)
				require.NoError(t, err)
				require.Equal(t, tc.data, string(data))
			})
			t.Run("Unmarshal", func(t *testing.T) {
				var dur ConstDuration
				err := yaml_v3.Unmarshal([]byte(tc.data), &dur)
				require.NoError(t, err)
				require.Equal(t, tc.d, dur.d)
			})
		})
	}
}
