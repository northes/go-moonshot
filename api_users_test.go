package moonshot_test

import (
	"context"
	"testing"

	"github.com/northes/go-moonshot/test"
)

func TestUserBalance(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.Users().Balance(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("User Balance: %+v", test.MarshalJsonToStringX(resp))
}
