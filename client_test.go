package moonshot_test

import (
	"errors"
	"os"

	"github.com/northes/go-moonshot"
)

func NewTestClient() (*moonshot.Client, error) {
	key, ok := os.LookupEnv("moonshot_key")
	if !ok {
		return nil, errors.New("missing environment variable: moonshot_key")
	}
	return moonshot.NewClient(moonshot.NewConfig(
		moonshot.SetAPIKey(key),
	))
}
