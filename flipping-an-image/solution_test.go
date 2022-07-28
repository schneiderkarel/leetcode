package flipping_an_image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_flipAndInvertImage(t *testing.T) {
	type args struct {
		image [][]int
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
				image: [][]int{
					{1, 0, 1},
					{0, 1, 0},
					{1, 0, 1},
				},
			},
			exp: exp{
				res: [][]int{
					{0, 1, 0},
					{1, 0, 1},
					{0, 1, 0},
				},
			},
		},
		{
			name: "invalid entry",
			args: args{
				image: [][]int{},
			},
			exp: exp{
				res: [][]int{},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, flipAndInvertImage(tc.args.image))
		})
	}
}

func Test_isEntryImageValid(t *testing.T) {
	var validHighImage [][]int
	for i := 0; i < 20; i++ {
		validHighImage = append(validHighImage, make([]int, 20))
	}

	type args struct {
		image [][]int
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
				image: [][]int{
					{1},
				},
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "valid - high",
			args: args{
				image: validHighImage,
			},
			exp: exp{
				res: true,
			},
		},
		{
			name: "invalid - too low",
			args: args{
				image: [][]int{},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - too high",
			args: args{
				image: make([][]int, 21),
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - image pixels shorter than image",
			args: args{
				image: [][]int{
					{1, 0, 1},
					{0, 1},
				},
			},
			exp: exp{
				res: false,
			},
		},
		{
			name: "invalid - image pixels contain not allowed digit",
			args: args{
				image: [][]int{
					{10},
				},
			},
			exp: exp{
				res: false,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp.res, isEntryImageValid(tc.args.image))
		})
	}
}
