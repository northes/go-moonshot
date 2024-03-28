package moonshot_test

import (
	"testing"

	"github.com/northes/gox"
)

func TestUpload(t *testing.T) {
	filePath := "./README.md"
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.Files().Upload(filePath)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", gox.JsonMarshalToStringX(resp))
}

func TestList(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.Files().Lists()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("file count: %d", len(resp.Data))
	t.Logf("%+v", gox.JsonMarshalToStringX(resp))
}

func TestDelete(t *testing.T) {
	fileID := "co2etto3r07e3eeneiig"
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.Files().Delete(fileID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", gox.JsonMarshalToStringX(resp))
}

func TestInfo(t *testing.T) {
	fileID := "co2eraqlnl9coc92ho2g"
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.Files().Info(fileID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", gox.JsonMarshalToStringX(resp))
}

func TestContent(t *testing.T) {
	fileID := "co2eraqlnl9coc92ho2g"
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.Files().Content(fileID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", gox.JsonMarshalToStringX(resp))
}
