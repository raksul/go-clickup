package clickup

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_UnmarshalJSON(t *testing.T) {
	r := require.New(t)

	p := Point{}

	table := []struct {
		Val string
		Exp string
	}{
		{"1", "1"},
		{"0.5", "0.5"},
		{"0.000001", "0.000001"},
	}

	for _, tt := range table {
		err := p.UnmarshalJSON([]byte(tt.Val))
		r.NoError(err)
		r.Equal(tt.Exp, p.Value.String())
	}
}

func Test_MarshalJSON(t *testing.T) {
	r := require.New(t)

	table := []struct {
		P   Point
		Exp string
	}{
		{Point{IntVal: Int64(9)}, "9"},
		{Point{FloatVal: Float64(0.5)}, "0.5"},
	}

	for _, tt := range table {
		v, err := tt.P.MarshalJSON()
		r.NoError(err)
		r.Equal(tt.Exp, string(v))
	}

}
