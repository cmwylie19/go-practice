package hello

import "testing"

func TestHello(t *testing.T) {
	expect := "Hello, world."
	if value := Hello(); value != expect {
		t.Errorf("Hello() = %q, want %q", value, expect)
	}
}

func TestProverb(t *testing.T) {
	expect := "Concurrency is not parallelism."
	if value := Proverb(); value != expect {
		t.Errorf("Proverb() = %q, expect %q", value, expect)
	}
}
