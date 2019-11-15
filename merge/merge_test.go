package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		name string
		a1   []int
		a2   []int
		want []int
	}{
		{
			name: "basic",
			a1:   []int{1, 5, 9, 9, 10},
			a2:   []int{1, 3, 8, 9},
			want: []int{1, 3, 5, 8, 9, 10},
		},
		{
			name: "empty",
			a1:   nil,
			a2:   []int{},
			want: []int{},
		},
		{
			name: "repeated values",
			a1:   []int{1, 1, 1, 1, 1, 1, 1},
			a2:   []int{1},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c1 := make(chan int)
			c2 := make(chan int)
			c := merge(c1, c2)
			go write(c1, tt.a1)
			go write(c2, tt.a2)
			got := []int{}
			for v := range c {
				got = append(got, v)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
