package minimum_consecutive_cards_to_pick_up

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumCardPickup(t *testing.T) {
	type args struct {
		cards []int
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
				cards: []int{1, 2, 3, 4, 5, 1, 1, 1},
			},
			exp: exp{
				res: 2,
			},
		},
		{
			name: "ok - impossible matching",
			args: args{
				cards: []int{1, 2, 3, 4, 5, 6},
			},
			exp: exp{
				res: -1,
			},
		},
		{
			name: "???",
			args: args{
				cards: []int{3, 4, 2, 3, 4, 7},
			},
			exp: exp{
				res: 4,
			},
		},
		{
			name: "invalid entry",
			args: args{
				cards: []int{},
			},
			exp: exp{
				res: -1,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, minimumCardPickup(tc.args.cards))
		})
	}
}

func Test_areEntryNumbersValid(t *testing.T) {
	type args struct {
		cards []int
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
				cards: []int{0, int(math.Pow10(6))},
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - high",
			args: args{
				cards: make([]int, int(math.Pow10(5))),
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - too low",
			args: args{
				cards: []int{},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - too high",
			args: args{
				cards: make([]int, int(math.Pow10(5))+1),
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - cards card too low",
			args: args{
				cards: []int{-1},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - cards card too high",
			args: args{
				cards: []int{int(math.Pow10(6) + 1)},
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, areEntryNumbersValid(tc.args.cards))
		})
	}
}
