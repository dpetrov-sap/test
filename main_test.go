package main

import (
	"testing"

	"github.com/dpetrov-sap/test/hello"
)

func TestHello(t *testing.T) {
	expected := "Hello, World1!"
	actual := hello.Hello()

	if actual != expected {
		t.Errorf("expected1 %q but got %q", expected, actual)
	}
}
