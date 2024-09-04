package test

import (
	"os"
)

func IsGithubActions() bool {
	if val, ok := os.LookupEnv("GITHUB_ACTIONS"); !ok || val != "true" {
		return false
	}
	return true
}
