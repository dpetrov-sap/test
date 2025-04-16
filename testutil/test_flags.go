package testutil

import (
	"os"
	"strings"
	

)


func ShouldSkipIntegrationTests() bool {
	return !AreIntegrationTestsEnabled()
}


func AreIntegrationTestsEnabled() bool {
	for _, env := range os.Environ() {
		lowerEnv := strings.ToLower(env)
		if strings.HasPrefix(lowerEnv, "integration=") {
			val := strings.TrimPrefix(lowerEnv, "integration=")
			if val == "true" || val == "enabled" || val == "enable" {
				return true
			}
		}
	}
	return false
}
