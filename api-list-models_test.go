package moonshot_test

import (
	"context"
	"testing"

	"github.com/northes/go-moonshot"
)

func TestListModels(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.ListModels(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", moonshot.MarshalToStringX(resp))
}
