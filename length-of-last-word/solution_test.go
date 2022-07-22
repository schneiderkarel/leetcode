package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
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
			name: "ok - prefix space",
			args: args{
				s: "   pear apple",
			},
			exp: exp{
				res: 5,
			},
		},
		{
			name: "ok - suffix space",
			args: args{
				s: "pear apple   ",
			},
			exp: exp{
				res: 5,
			},
		},
		{
			name: "invalid entry",
			args: args{
				s: "",
			},
			exp: exp{
				res: 0,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, lengthOfLastWord(tc.args.s))
		})
	}
}

func Test_validEntries(t *testing.T) {
	type args struct {
		s string
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
			name: "ok - short",
			args: args{
				s: "a",
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "ok - long",
			args: args{
				s: string(make([]byte, int(math.Pow10(4)))),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid entry - too short",
			args: args{
				s: "",
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid entry - too long",
			args: args{
				s: string(make([]byte, int(math.Pow10(4)+1))),
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, validEntries(tc.args.s))
		})
	}
}
