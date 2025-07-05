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

func TestMyCut(t *testing.T) {
	t.Run("cut found", func(t *testing.T) {
		before, after, found := MyCut("cut", "hellocutworld")
		expectedBefore := "hello"
		expectedAfter := "world"

		assertCutExpectations(t, expectedBefore, before, expectedAfter, after, true, found)
	})

	t.Run("cut not found", func(t *testing.T) {
		before, after, found := MyCut("cut", "helloworld")
		expectedBefore := "helloworld"
		expectedAfter := ""

		assertCutExpectations(t, expectedBefore, before, expectedAfter, after, false, found)
	})
}

func assertCutExpectations(t *testing.T, expectedBefore string, before string, expectedAfter string, after string, expectedFound bool, found bool) {
	t.Helper()
	if expectedBefore != before {
		t.Errorf("failed before, got %s, expected %s", before, expectedBefore)
	}
	if expectedAfter != after {
		t.Errorf("failed after, got %s, expected %s", after, expectedAfter)
	}
	if expectedFound != found {
		t.Errorf("failed found, got %t, expected %t", found, expectedFound)
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
