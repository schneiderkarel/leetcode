package two_sum

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums      []int
		targetNum int
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
				nums:      []int{4, 5, 10, 2, 6},
				targetNum: 16,
			},
			exp: exp{
				res: []int{2, 4},
			},
		},
		{
			name: "ok negative",
			args: args{
				nums:      []int{-1, -6, 0},
				targetNum: -7,
			},
			exp: exp{
				res: []int{1, 0},
			},
		},
		{
			name: "no result",
			args: args{
				nums:      []int{1, 2},
				targetNum: 2,
			},
			exp: exp{
				res: []int{},
			},
		},
		{
			name: "invalid entries",
			args: args{
				nums:      []int{},
				targetNum: 0,
			},
			exp: exp{
				res: []int{},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.exp.res, twoSum(tc.args.nums, tc.args.targetNum))
		})
	}
}

func Test_areEntriesValid(t *testing.T) {
	type args struct {
		nums      []int
		targetNum int
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
			name: "valid",
			args: args{
				nums:      make([]int, int(math.Pow10(4))),
				targetNum: int(math.Pow10(9)),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid nums",
			args: args{
				nums:      []int{},
				targetNum: 0,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid target",
			args: args{
				nums:      []int{1, 2},
				targetNum: int(-math.Pow10(9) - 1),
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, areEntriesValid(tc.args.nums, tc.args.targetNum))
		})
	}
}

func Test_areEntryNumbersValid(t *testing.T) {
	type args struct {
		nums []int
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
			name: "valid - long",
			args: args{
				nums: make([]int, int(math.Pow10(4))),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - nums",
			args: args{
				nums: []int{int(-math.Pow10(9)), int(math.Pow10(9))},
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - nums too short",
			args: args{
				nums: []int{},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - nums too long",
			args: args{
				nums: make([]int, int(math.Pow10(4)+1)),
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - num in nums too low",
			args: args{
				nums: []int{int(-math.Pow10(9) - 1), 2},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - num in nums too high",
			args: args{
				nums: []int{int(math.Pow10(9) + 1), 2},
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, areEntryNumbersValid(tc.args.nums))
		})
	}
}

func Test_isEntryTargetNumberValid(t *testing.T) {
	type args struct {
		targetNum int
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
			name: "valid - high",
			args: args{
				targetNum: int(math.Pow10(9)),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - low",
			args: args{
				targetNum: int(-math.Pow10(9)),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - target too low",
			args: args{
				targetNum: int(-math.Pow10(9) - 1),
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - target too high",
			args: args{
				targetNum: int(math.Pow10(9) + 1),
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, isEntryTargetNumberValid(tc.args.targetNum))
		})
	}
}
