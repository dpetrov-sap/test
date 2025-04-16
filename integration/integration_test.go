package integration

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"

	"github.com/dpetrov-sap/test/testutil"
)

func TestIntegrationOutboxHandler(t *testing.T) {
    if testutil.ShouldSkipIntegrationTests() {
        t.SkipNow()
        return
    }
}


func TestHelloWorldOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run main.go: %v", err)
	}
	output := strings.TrimSpace(out.String())
	expected := "Hello, World!"

	if output != expected {
		t.Errorf("expected %q but got %q", expected, output)
	}
}
