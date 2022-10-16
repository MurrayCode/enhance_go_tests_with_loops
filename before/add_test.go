package before

import "testing"

func TestAdd(t *testing.T) {
	t.Run("returns correct integer", func(t *testing.T) {
		got := AddNums(2, 2)
		exp := 4
		if got != exp {
			t.Errorf("got %d, expected %d", got, exp)
		}
	})
	t.Run("returns correct integer", func(t *testing.T) {
		got := AddNums(2, 10)
		exp := 12
		if got != exp {
			t.Errorf("got %d, expected %d", got, exp)
		}
	})
	t.Run("returns correct integer", func(t *testing.T) {
		got := AddNums(-1, 3)
		exp := 2
		if got != exp {
			t.Errorf("got %d, expected %d", got, exp)
		}
	})
}

func TestCombineString(t *testing.T) {
	t.Run("return correct string", func(t *testing.T) {
		got := CombineString("test", "ing")
		exp := "testing"
		if got != exp {
			t.Errorf("got %s, expected %s", got, exp)
		}
	})
	t.Run("return correct string", func(t *testing.T) {
		got := CombineString("foo", "bar")
		exp := "foobar"
		if got != exp {
			t.Errorf("got %s, expected %s", got, exp)
		}
	})
}
