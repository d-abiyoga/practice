package main

import "testing"

func TestPow(t *testing.T) {
	testCases := []struct {
		desc string
		a    float32
		b    float32
		want float32
	}{
		{
			desc: "normal",
			a:    3,
			b:    4,
			want: 81,
		},
		{
			desc: "by zero",
			a:    3,
			b:    0,
			want: 1,
		},
		{
			desc: "negative",
			a:    3,
			b:    -4,
			want: float32(1) / float32(81),
		},
		{
			desc: "odd",
			a:    3,
			b:    5,
			want: 243,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Pow(tC.a, tC.b)
			if got != tC.want {
				t.Errorf("Pow(%f, %f) = %f; want %f", tC.a, tC.b, got, tC.want)
			}

		})
	}
}
