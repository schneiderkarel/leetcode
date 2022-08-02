package contains_duplicate_ii

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums   []int
		offset int
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
			name: "ok - nearby",
			args: args{
				nums:   []int{1, 2, 3, 1},
				offset: 3,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "ok - nothing nearby",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 1},
				offset: 1,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid entries",
			args: args{
				nums:   []int{},
				offset: -1,
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, containsNearbyDuplicate(tc.args.nums, tc.args.offset))
		})
	}
}

func Test_areEntriesValid(t *testing.T) {
	type args struct {
		nums   []int
		offset int
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
			name: "ok",
			args: args{
				nums:   []int{1},
				offset: 0,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid nums",
			args: args{
				nums:   []int{},
				offset: 0,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid offset",
			args: args{
				nums:   []int{1},
				offset: -1,
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, areEntriesValid(tc.args.nums, tc.args.offset))
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

func Test_isEntryOffsetValid(t *testing.T) {
	type args struct {
		offset int
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
			name: "ok - low",
			args: args{
				offset: 0,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "ok - high",
			args: args{
				offset: int(math.Pow10(5)),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - too low",
			args: args{
				offset: -1,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - too high",
			args: args{
				offset: int(math.Pow10(5)) + 1,
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, isEntryOffsetValid(tc.args.offset))
		})
	}
}
