package moonshot_test

import (
	"errors"
	"os"
	"testing"

	"github.com/northes/go-moonshot"
	"github.com/stretchr/testify/require"
)

func NewTestClient() (*moonshot.Client, error) {
	key, ok := os.LookupEnv("MOONSHOT_KEY")
	if !ok {
		return nil, errors.New("missing environment variable: MOONSHOT_KEY")
	}
	return moonshot.NewClient(moonshot.NewConfig(
		moonshot.SetAPIKey(key),
	))
}

func TestConfig(t *testing.T) {
	tt := require.New(t)

	_, err := moonshot.NewClient(nil)
	tt.NotNil(err, "must got a required api key error")

	_, err = moonshot.NewClient(moonshot.NewConfig())
	tt.NotNil(err, "must got a required api key error")
}
