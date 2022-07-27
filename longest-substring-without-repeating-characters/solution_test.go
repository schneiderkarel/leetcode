package longest_substring_without_repeating_characters

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		str string
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
				str: "abcabcbb",
			},
			exp: exp{
				res: 3,
			},
		},
		{
			name: "ok - only whitespace",
			args: args{
				str: " ",
			},
			exp: exp{
				res: 1,
			},
		},
		{
			name: "invalid entry",
			args: args{
				str: "",
			},
			exp: exp{
				res: 0,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, lengthOfLongestSubstring(tc.args.str))
		})
	}
}

func Test_isEntryStringValid(t *testing.T) {
	type args struct {
		str string
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
				str: "a",
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - high",
			args: args{
				str: string(make([]byte, int(math.Pow(5*10, 4)))),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - too low",
			args: args{
				str: "",
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - too high",
			args: args{
				str: string(make([]byte, int(math.Pow(5*10, 4))+1)),
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - string contains not ascii char",
			args: args{
				str: "héllo",
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, isEntryStringValid(tc.args.str))
		})
	}
}
