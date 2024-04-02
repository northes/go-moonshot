package moonshot_test

import (
	"errors"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/northes/go-moonshot"
	"github.com/stretchr/testify/require"
)

func NewTestClient() (*moonshot.Client, error) {
	key, ok := os.LookupEnv("MOONSHOT_KEY")
	if !ok {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
		key, ok = os.LookupEnv("MOONSHOT_KEY")
	}
	if !ok {
		return nil, errors.New("missing environment variable: MOONSHOT_KEY")
	}
	debug, ok := os.LookupEnv("MOONSHOT_DEBUG")
	if !ok {
		debug = "false"
	}

	cfg := moonshot.NewConfig(
		moonshot.SetAPIKey(key),
		moonshot.SetHost(moonshot.DefaultHost),
	)

	isDebug, err := strconv.ParseBool(debug)
	if err != nil {
		return nil, err
	}
	if isDebug {
		cfg.Debug = isDebug
	}

	return moonshot.NewClient(cfg)
}

func TestNewClient(t *testing.T) {
	tt := require.New(t)

	_, err := moonshot.NewClient(nil)
	tt.NotNil(err, "must got a required api key error")

	_, err = moonshot.NewClient(moonshot.NewConfig())
	tt.NotNil(err, "must got a required api key error")
}
