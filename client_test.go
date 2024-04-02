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
		moonshot.WithAPIKey(key),
		moonshot.WithHost(moonshot.DefaultHost),
	)

	isDebug, err := strconv.ParseBool(debug)
	if err != nil {
		return nil, err
	}
	if isDebug {
		cfg.Debug = isDebug
	}

	return moonshot.NewClientWithConfig(cfg)
}

func TestNewClient(t *testing.T) {
	tt := require.New(t)

	cli, err := moonshot.NewClient("")
	tt.NotNil(err)
	tt.Nil(cli)

	cli, err = moonshot.NewClient("xxxx")
	tt.Nil(err)
	tt.NotNil(cli)
}

func TestNewClientWithConfig(t *testing.T) {
	tt := require.New(t)

	cli, err := moonshot.NewClientWithConfig(nil)
	tt.NotNil(err, "must got a required api key error")
	tt.Nil(cli)

	cli, err = moonshot.NewClientWithConfig(moonshot.NewConfig())
	tt.NotNil(err, "must got a required api key error")
	tt.Nil(cli)

	cli, err = moonshot.NewClientWithConfig(
		moonshot.NewConfig(
			moonshot.WithAPIKey("xxxx"),
		),
	)
	tt.Nil(err)
	tt.NotNil(cli)
}
