package climbing_stairs

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
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
			name: "ok - palindrome",
			args: args{
				str: "A man, a plan, a canal: Panama",
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "ok - not palindrome",
			args: args{
				str: "A man, a planz, a canal: Panama",
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "ok - empty",
			args: args{
				str: " ",
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid entry",
			args: args{
				str: "á",
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, isPalindrome(tc.args.str))
		})
	}
}

func Test_normalizeString(t *testing.T) {
	type args struct {
		str string
	}
	type exp struct {
		res string
	}
	tcs := []struct {
		name string
		args args
		exp  exp
	}{
		{
			name: "ok",
			args: args{
				str: "A man, a plan, a canal: Panama",
			},
			exp: exp{
				res: "amanaplanacanalpanama",
			},
		},
		{
			name: "empty",
			args: args{
				str: "  ",
			},
			exp: exp{
				res: "",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, normalizeString(tc.args.str))
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
			name: "valid - short",
			args: args{
				str: "a",
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - long",
			args: args{
				str: string(make([]byte, int(2*math.Pow10(5)))),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - too short",
			args: args{
				str: "",
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - too long",
			args: args{
				str: string(make([]byte, int(2*math.Pow10(5))+1)),
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
