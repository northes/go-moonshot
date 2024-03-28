package moonshot_test

import (
	"context"
	"testing"

	"github.com/northes/gox"
)

func TestListModels(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.Models().List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", gox.JsonMarshalToStringX(resp))
}
