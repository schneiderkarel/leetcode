package number_of_one_bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
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
				num: 00000000000000000000000000001011,
			},
			exp: exp{
				res: 3,
			},
		},
		{
			name: "empty",
			args: args{
				num: 0,
			},
			exp: exp{
				res: 0,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, hammingWeight(tc.args.num))
		})
	}
}
