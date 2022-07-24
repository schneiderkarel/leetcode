package longest_common_prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
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
			name: "ok - prefix",
			args: args{
				strs: []string{"app", "apple", "application"},
			},
			exp: exp{
				res: "app",
			},
		},
		{
			name: "ok - no prefix",
			args: args{
				strs: []string{"apple", "pear", "cherry"},
			},
			exp: exp{
				res: "",
			},
		},
		{
			name: "invalid entry",
			args: args{
				strs: []string{},
			},
			exp: exp{
				res: "",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, longestCommonPrefix(tc.args.strs))
		})
	}
}

func Test_areEntryStringsValid(t *testing.T) {
	var validMaxLenStrs []string
	for i := 0; i < 200; i++ {
		validMaxLenStrs = append(validMaxLenStrs, "a")
	}

	type args struct {
		strs []string
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
				strs: []string{"apple"},
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - long",
			args: args{
				strs: validMaxLenStrs,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - too short",
			args: args{
				strs: []string{},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - too long",
			args: args{
				strs: make([]string, 201),
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - strs str too long",
			args: args{
				strs: []string{string(make([]byte, 201))},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - invalid char inside strs",
			args: args{
				strs: []string{"A"},
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, areEntryStringsValid(tc.args.strs))
		})
	}
}
