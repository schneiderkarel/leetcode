package single_number

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{2, 2, 3, 3, 1},
			},
			exp: exp{
				res: 1,
			},
		},
		{
			name: "invalid entry",
			args: args{
				nums: []int{},
			},
			exp: exp{
				res: 0,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, singleNumber(tc.args.nums))
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
			name: "ok - short",
			args: args{
				nums: []int{1},
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - too short",
			args: args{
				nums: []int{},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - too long",
			args: args{
				nums: make([]int, int(3*math.Pow10(4)+1)),
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - nums num too low",
			args: args{
				nums: []int{-int(3*math.Pow10(4)) - 1},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - nums num too high",
			args: args{
				nums: []int{int(3*math.Pow10(4) + 1)},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - nums num multiple ones appearances",
			args: args{
				nums: []int{1, 2},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - nums num more than two appearances",
			args: args{
				nums: []int{2, 2, 2},
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
