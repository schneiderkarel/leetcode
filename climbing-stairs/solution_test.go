package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
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
				n: 8,
			},
			exp: exp{
				res: 34,
			},
		},
		{
			name: "invalid entry",
			args: args{
				n: -1,
			},
			exp: exp{
				res: 0,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, climbStairs(tc.args.n))
		})
	}
}

func Test_validEntry(t *testing.T) {
	type args struct {
		n int
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
			name: "valid - min",
			args: args{
				n: 1,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - max",
			args: args{
				n: 45,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - beyond min",
			args: args{
				n: 0,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - beyond max",
			args: args{
				n: 46,
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, validEntry(tc.args.n))
		})
	}
}
