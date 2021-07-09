package hello

import "testing"

func TestHello(t *testing.T) {
	if got, want := Hello(), "Hello, World!"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
