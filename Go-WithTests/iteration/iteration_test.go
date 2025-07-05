package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("Got %s, expected %s", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func BenchmarkBenchedRepeat(b *testing.B) {
	for b.Loop() {
		BenchedRepeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 6)
	fmt.Println(repeated)
	// Output: aaaaaa
}
func ExampleBenchedRepeat() {
	repeated := BenchedRepeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
