package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
			name: "ok positive",
			args: args{
				nums:   []int{4, 5, 10, 2, 6},
				target: 16,
			},
			exp: exp{
				res: []int{2, 4},
			},
		},
		{
			name: "ok negative",
			args: args{
				nums:   []int{-1, -6, 0},
				target: -7,
			},
			exp: exp{
				res: []int{1, 0},
			},
		},
		{
			name: "no result",
			args: args{
				nums:   []int{1, 2},
				target: 2,
			},
			exp: exp{
				res: []int{},
			},
		},
		{
			name: "invalid entries",
			args: args{
				nums:   []int{},
				target: 0,
			},
			exp: exp{
				res: []int{},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.exp.res, twoSum(tc.args.nums, tc.args.target))
		})
	}
}

func Test_validEntries(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
			name: "valid - edge case 1",
			args: args{
				nums:   make([]int, int(math.Pow10(4))),
				target: int(math.Pow10(9)),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - edge case 2",
			args: args{
				nums:   []int{int(-math.Pow10(9)), int(math.Pow10(9))},
				target: int(-math.Pow10(9)),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "not valid - nums too short",
			args: args{
				nums:   []int{},
				target: 0,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "not valid - nums too long",
			args: args{
				nums:   make([]int, int(math.Pow10(4)+1)),
				target: 0,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "not valid - num in nums too small",
			args: args{
				nums:   []int{int(-math.Pow10(9) - 1), 2},
				target: 0,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "not valid - num in nums too big",
			args: args{
				nums:   []int{int(math.Pow10(9) + 1), 2},
				target: 0,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "not valid - target too small",
			args: args{
				nums:   []int{1, 2},
				target: int(-math.Pow10(9) - 1),
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "not valid - target too big",
			args: args{
				nums:   []int{1, 2},
				target: int(math.Pow10(9) + 1),
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, validEntries(tc.args.nums, tc.args.target))
		})
	}
}
