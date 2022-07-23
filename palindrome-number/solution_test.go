package palindrome_number

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		num int
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
			name: "positive palindrome num",
			args: args{
				num: 1001,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "negative non-palindrome num",
			args: args{
				num: -1001,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "positive non-palindrome num",
			args: args{
				num: 100,
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid entry",
			args: args{
				num: int(math.Pow(-2, 31) - 1),
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, isPalindrome(tc.args.num))
		})
	}
}

func Test_isEntryNumberValid(t *testing.T) {
	type args struct {
		num int
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
				num: int(-math.Pow(2, 31) + 1),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - high",
			args: args{
				num: int(math.Pow(2, 31) - 1),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - too low",
			args: args{
				num: int(-math.Pow(2, 31) - 1),
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - too high",
			args: args{
				num: int(math.Pow(2, 31) + 1),
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, isEntryNumberValid(tc.args.num))
		})
	}
}
