package best_time_to_buy_and_sell_stock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	type exp struct {
		res []int
	}
	tcs := []struct {
		name string
		args args
		exp  exp
	}{
		{
			name: "ok - 1",
			args: args{
				digits: []int{1, 9, 9},
			},
			exp: exp{
				res: []int{2, 0, 0},
			},
		},
		{
			name: "invalid entry",
			args: args{
				digits: []int{},
			},
			exp: exp{
				res: []int{},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, plusOne(tc.args.digits))
		})
	}
}

func Test_recursiveCallsOverDigits(t *testing.T) {
	type args struct {
		index  int
		digit  int
		digits []int
	}
	type exp struct {
		res []int
	}
	tcs := []struct {
		name string
		args args
		exp  exp
	}{
		{
			name: "ok - 1",
			args: args{
				index:  2,
				digit:  9,
				digits: []int{1, 9, 9},
			},
			exp: exp{
				res: []int{2, 0, 0},
			},
		},
		{
			name: "ok - 3",
			args: args{
				index:  0,
				digit:  9,
				digits: []int{9},
			},
			exp: exp{
				res: []int{1, 0},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, recursiveCallsOverDigits(tc.args.index, tc.args.digit, tc.args.digits))
		})
	}
}

func Test_areEntryDigitsValid(t *testing.T) {
	type args struct {
		digits []int
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
			name: "valid - short",
			args: args{
				digits: []int{1},
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - long",
			args: args{
				digits: make([]int, 100),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - too short",
			args: args{
				digits: []int{},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - too long",
			args: args{
				digits: make([]int, 101),
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - digits digit too low",
			args: args{
				digits: []int{-1},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - digits digit too low",
			args: args{
				digits: []int{10},
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, areEntryDigitsValid(tc.args.digits))
		})
	}
}
