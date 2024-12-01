package int

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("We got %d but we want %d", got, want)
	}
}

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}

func ExampleAdd_negative() {
	sum := Add(-4, -6)
	fmt.Println(sum)
	// Output: -10
}
