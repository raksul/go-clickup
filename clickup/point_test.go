package clickup

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_String(t *testing.T) {
	r := require.New(t)

	table := []struct {
		P   Point
		Exp string
	}{
		{Point{IntVal: Int64(9)}, "9"},
		{Point{FloatVal: Float64(0.5)}, "0.5"},
	}

	for _, tt := range table {
		r.Equal(tt.Exp, tt.P.String())
	}
}

func Test_Int(t *testing.T) {
	r := require.New(t)

	table := []struct {
		P   Point
		Exp int64
	}{
		{Point{IntVal: Int64(9)}, 9},
		{Point{FloatVal: Float64(0.5)}, 0},
		{Point{FloatVal: Float64(1.5)}, 1},
	}

	for _, tt := range table {
		r.Equal(tt.Exp, tt.P.Int())
	}
}

func Test_Float(t *testing.T) {
	r := require.New(t)

	table := []struct {
		P   Point
		Exp float64
	}{
		{Point{IntVal: Int64(9)}, 9.0},
		{Point{FloatVal: Float64(0.5)}, 0.5},
		{Point{FloatVal: Float64(1.5)}, 1.5},
	}

	for _, tt := range table {
		r.Equal(tt.Exp, tt.P.Float())
	}
}

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
		r.Equal(tt.Exp, p.String())
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
