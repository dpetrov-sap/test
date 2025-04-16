package main

import (
	"testing"

	"github.com/dpetrov-sap/test/hello"
)

func TestHello(t *testing.T) {
	expected := "Hello, World!"
	actual := hello.Hello()

	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}
