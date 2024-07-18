package moonshot_test

import (
	"context"
	"github.com/northes/go-moonshot"
	"github.com/northes/go-moonshot/test"
	"testing"
)

func TestContextCache_Create(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	response, err := cli.ContextCache().Create(ctx, &moonshot.ContextCacheCreateRequest{})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(test.MarshalJsonToStringX(response))
}
