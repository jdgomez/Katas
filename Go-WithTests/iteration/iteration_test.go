package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Got %s, expected %s", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}

func BenchmarkBenchedRepeat(b *testing.B) {
	for b.Loop() {
		BenchedRepeat("a")
	}
}
