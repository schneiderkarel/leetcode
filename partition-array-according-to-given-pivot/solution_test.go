package partition_array_according_to_given_pivot

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pivotArray(t *testing.T) {
	type args struct {
		nums  []int
		pivot int
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
			name: "ok",
			args: args{
				nums:  []int{9, 12, 5, 10, 14, 3, 10},
				pivot: 10,
			},
			exp: exp{
				res: []int{9, 5, 3, 10, 10, 12, 14},
			},
		},
		{
			name: "invalid entries",
			args: args{
				nums:  []int{},
				pivot: 1,
			},
			exp: exp{
				res: []int{},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, pivotArray(tc.args.nums, tc.args.pivot))
		})
	}
}

func Test_areEntriesValid(t *testing.T) {
	type args struct {
		nums  []int
		pivot int
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
			name: "valid - nums short",
			args: args{
				nums:  []int{1},
				pivot: 1,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - nums long",
			args: args{
				nums:  make([]int, int(math.Pow10(5))),
				pivot: 0,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - nums too short",
			args: args{
				nums:  []int{},
				pivot: 0,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - nums too long",
			args: args{
				nums:  make([]int, int(math.Pow10(5)+1)),
				pivot: 0,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - nums num too low",
			args: args{
				nums:  []int{-int(math.Pow10(6) - 1)},
				pivot: 0,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - nums num too high",
			args: args{
				nums:  []int{-int(math.Pow10(6) + 1)},
				pivot: 0,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - no pivot hit",
			args: args{
				nums:  []int{1},
				pivot: 10,
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, areEntriesValid(tc.args.nums, tc.args.pivot))
		})
	}
}
