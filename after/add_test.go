package after

import "testing"

func TestAddNums(t *testing.T) {
	var tests = []struct {
		name   string
		inputs []int
		ans    int
	}{
		{
			"Two positive integers",
			[]int{2, 2},
			4,
		},
		{
			"Two negative integers",
			[]int{-10, -10},
			-20,
		},
		{
			"One positive, one negative integer",
			[]int{-1, 3},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddNums(tt.inputs[0], tt.inputs[1])
			exp := tt.ans
			if got != exp {
				t.Errorf("got %d, expected %d", got, exp)
			}
		})

	}
}

func TestCombineString(t *testing.T) {
	var tests = []struct {
		name   string
		inputs []string
		ans    string
	}{
		{
			"test name one",
			[]string{"test", "ing"},
			"testing",
		},
		{
			"test name two",
			[]string{"foo", "bar"},
			"foobar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CombineString(tt.inputs[0], tt.inputs[1])
			exp := tt.ans
			if got != exp {
				t.Errorf("got %s, expected %s", got, exp)
			}
		})

	}
}
