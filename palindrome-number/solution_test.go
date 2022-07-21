package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
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
			name: "positive palindrome x",
			args: args{
				x: 1001,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "negative x",
			args: args{
				x: -1001,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "positive x",
			args: args{
				x: 100,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid x",
			args: args{
				x: int(math.Pow(-2, 31) - 1),
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, isPalindrome(tc.args.x))
		})
	}
}

func Test_validEntry(t *testing.T) {
	type args struct {
		x int
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
			name: "valid positive x",
			args: args{
				x: int(math.Pow(2, 31) - 1),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid negative x",
			args: args{
				x: int(math.Pow(-2, 31)),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid positive x",
			args: args{
				x: int(math.Pow(2, 31)),
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid negative x",
			args: args{
				x: int(math.Pow(-2, 31) - 1),
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, validEntry(tc.args.x))
		})
	}
}
