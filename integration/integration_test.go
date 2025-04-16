package integration

import (
	"bytes"
	"os"
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

	t.Log("Running TestIntegrationOutboxHandler (placeholder)")
}

func TestHelloWorldOutput(t *testing.T) {
	if testutil.ShouldSkipIntegrationTests() {
		t.SkipNow()
		return
	}

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

func TestAreIntegrationTestsEnabled(t *testing.T) {
	// Case 1: When integration=true
	os.Setenv("integration", "true")
	if !testutil.AreIntegrationTestsEnabled() {
		t.Errorf("Expected integration tests to be enabled")
	}

	// Case 2: When integration=enabled
	os.Setenv("integration", "enabled")
	if !testutil.AreIntegrationTestsEnabled() {
		t.Errorf("Expected integration tests to be enabled")
	}

	// Case 3: When integration=enable
	os.Setenv("integration", "enable")
	if !testutil.AreIntegrationTestsEnabled() {
		t.Errorf("Expected integration tests to be enabled")
	}

	// Case 4: When integration is not set or has an invalid value
	os.Unsetenv("integration")
	if testutil.AreIntegrationTestsEnabled() {
		t.Errorf("Expected integration tests to be disabled")
	}
}
