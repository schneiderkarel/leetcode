package pascals_triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate(t *testing.T) {
	type args struct {
		numRows int
	}
	type exp struct {
		res [][]int
	}
	tcs := []struct {
		name string
		args args
		exp  exp
	}{
		{
			name: "ok",
			args: args{
				numRows: 5,
			},
			exp: exp{
				res: [][]int{
					{
						1,
					},
					{
						1, 1,
					},
					{
						1, 2, 1,
					},
					{
						1, 3, 3, 1,
					},
					{
						1, 4, 6, 4, 1,
					},
				},
			},
		},
		{
			name: "invalid entry",
			args: args{
				numRows: 0,
			},
			exp: exp{
				res: [][]int{},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, generate(tc.args.numRows))
		})
	}
}

func Test_isEntryNumberValid(t *testing.T) {
	type args struct {
		numRows int
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
			name: "valid - low",
			args: args{
				numRows: 1,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - high",
			args: args{
				numRows: 30,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - too low",
			args: args{
				numRows: 0,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - too high",
			args: args{
				numRows: 31,
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, isEntryNumberValid(tc.args.numRows))
		})
	}
}
