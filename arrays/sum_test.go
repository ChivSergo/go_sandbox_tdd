package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("arrays", func(t *testing.T) {
		got := Sum([]int{1, 2, 3, 4, 5})
		want := 15
		if want != got {
			t.Errorf("We got %d, but want %d", got, want)
		}
	})
	t.Run("slices", func(t *testing.T) {
		got := Sum([]int{1, 2, 3})
		want := 15
		if want != got {
			t.Errorf("We got %d, but want %d", got, want)
		}
	})
}
