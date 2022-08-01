package contains_duplicate

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsDuplicate(t *testing.T) {
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
			name: "ok - duplicate",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 1},
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "ok - no duplicate",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid entry",
			args: args{
				nums: []int{},
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, containsDuplicate(tc.args.nums))
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
			name: "ok - long",
			args: args{
				nums: make([]int, int(math.Pow10(5))),
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
				nums: make([]int, int(math.Pow10(5))+1),
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - nums num too low",
			args: args{
				nums: []int{int(-math.Pow10(9) - 1)},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - nums num too high",
			args: args{
				nums: []int{int(math.Pow10(9) + 1)},
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
