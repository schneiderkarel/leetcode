package climbing_stairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStairs(t *testing.T) {
	type args struct {
		num int
	}
	type exp struct {
		res int
	}
	tcs := []struct {
		name string
		args args
		exp  exp
	}{
		{
			name: "ok",
			args: args{
				num: 8,
			},
			exp: exp{
				res: 34,
			},
		},
		{
			name: "invalid entry",
			args: args{
				num: -1,
			},
			exp: exp{
				res: 0,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, climbStairs(tc.args.num))
		})
	}
}

func Test_isEntryNumberValid(t *testing.T) {
	type args struct {
		num int
	}
	type exp struct {
		res bool
	}
	tcs := []struct {
		name string
		args args
		exp  exp
	}{
		{
			name: "valid - low",
			args: args{
				num: 1,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - high",
			args: args{
				num: 45,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - too low",
			args: args{
				num: 0,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - too high",
			args: args{
				num: 46,
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, isEntryNumberValid(tc.args.num))
		})
	}
}
