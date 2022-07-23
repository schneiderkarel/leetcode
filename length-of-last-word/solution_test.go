package length_of_last_word

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLastWord(t *testing.T) {
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
			name: "ok - prefix space",
			args: args{
				str: "   pear apple",
			},
			exp: exp{
				res: 5,
			},
		},
		{
			name: "ok - suffix space",
			args: args{
				str: "pear apple   ",
			},
			exp: exp{
				res: 5,
			},
		},
		{
			name: "ok - short",
			args: args{
				str: "a ",
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
			assert.Equal(t, tc.exp.res, lengthOfLastWord(tc.args.str))
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
				str: string(make([]byte, int(math.Pow10(4)))),
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
				str: string(make([]byte, int(math.Pow10(4)+1))),
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
