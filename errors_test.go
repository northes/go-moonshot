package moonshot_test

import (
	"context"
	"testing"

	"github.com/northes/go-moonshot"
	"github.com/stretchr/testify/require"
)

func TestErrors(t *testing.T) {
	tt := require.New(t)
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	_, err = cli.Chat().Completions(context.Background(), moonshot.NewChatCompletionsBuilder().SetModel("xxxx").ToRequest())
	tt.NotNil(err)
	t.Log(err)
}
