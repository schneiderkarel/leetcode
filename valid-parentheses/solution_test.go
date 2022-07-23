package valid_parentheses

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
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
			name: "ok",
			args: args{
				str: "{()[]}",
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid",
			args: args{
				str: "{([)]}",
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - empty queue",
			args: args{
				str: "})",
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - one char",
			args: args{
				str: "{",
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid entry",
			args: args{
				str: "",
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, isValid(tc.args.str))
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
				str: "{",
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - long",
			args: args{
				str: string(make([]byte, 0, int(math.Pow10(4)))),
			},
			exp: exp{
				res: false,
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
				str: string(make([]byte, 0, int(math.Pow10(4)+1))),
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
