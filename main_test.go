package main 

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, World!"
	actual := Hello()

	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}