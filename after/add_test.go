package after

import "testing"

func TestAdd(t *testing.T) {
	var tests = []struct {
		inputs []int
		ans    int
	}{
		{
			[]int{2, 2},
			4,
		},
		{
			[]int{2, 10},
			12,
		},
		{
			[]int{-1, 3},
			2,
		},
	}
	for _, tt := range tests {
		t.Run("returns correct integer", func(t *testing.T) {
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
		inputs []string
		ans    string
	}{
		{
			[]string{"test", "ing"},
			"testing",
		},
		{
			[]string{"foo", "bar"},
			"foobar",
		},
	}
	for _, tt := range tests {
		t.Run("return correct string", func(t *testing.T) {
			got := CombineString(tt.inputs[0], tt.inputs[1])
			exp := tt.ans
			if got != exp {
				t.Errorf("got %s, expected %s", got, exp)
			}
		})

	}
}
