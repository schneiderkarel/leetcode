package best_time_to_buy_and_sell_stock

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
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
			name: "ok - profit",
			args: args{
				prices: []int{2, 4, 10, 9, 0},
			},
			exp: exp{
				res: 8,
			},
		},
		{
			name: "ok - no profit",
			args: args{
				prices: []int{5, 2, 1, 0},
			},
			exp: exp{
				res: 0,
			},
		},
		{
			name: "invalid entry",
			args: args{
				prices: []int{},
			},
			exp: exp{
				res: 0,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, maxProfit(tc.args.prices))
		})
	}
}

func Test_areEntryPricesValid(t *testing.T) {
	type args struct {
		prices []int
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
				prices: []int{1},
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - long",
			args: args{
				prices: make([]int, int(math.Pow10(5))),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - too short",
			args: args{
				prices: []int{},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - too long",
			args: args{
				prices: make([]int, int(math.Pow10(5))+1),
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - prices price too low",
			args: args{
				prices: []int{-1},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - prices price too low",
			args: args{
				prices: []int{int(math.Pow10(4) + 1)},
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, areEntryPricesValid(tc.args.prices))
		})
	}
}
