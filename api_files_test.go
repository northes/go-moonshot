package moonshot_test

import (
	"os"
	"testing"

	"github.com/northes/go-moonshot"
	"github.com/northes/go-moonshot/test"
	"github.com/northes/gox"
	"github.com/stretchr/testify/require"
)

var (
	fileID      string
	filePath    string
	fileContent []byte
)

func init() {
	fileContent = test.GenerateTestContent()
	var err error
	filePath, err = test.GenerateTestFile(fileContent)
	if err != nil {
		panic(err)
	}
}

func TestFilesUpload(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.Files().Upload(filePath)
	if err != nil {
		t.Fatal(err)
	}
	fileID = resp.ID
	t.Logf("%+v", gox.JsonMarshalToStringX(resp))
	err = os.Remove(filePath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFilesList(t *testing.T) {
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

func TestFilesInfo(t *testing.T) {
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

func TestFilesContent(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.Files().Content(fileID)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, []byte(resp.Content), fileContent, "file content not match")

	t.Logf("%+v", gox.JsonMarshalToStringX(resp))
}

func TestFilesDelete(t *testing.T) {
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

func TestFilesDeleteAll(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}

	for range 10 {
		fp, _ := test.GenerateTestFile(test.GenerateTestContent())
		_, _ = cli.Files().Upload(fp)
	}

	listResp, err := cli.Files().Lists()
	if err != nil {
		t.Fatal(err)
	}

	deleteAllResp, err := cli.Files().DeleteAll()
	if err != nil {
		t.Fatal(err)
	}

	require.EqualValues(t, func(in []*moonshot.FilesListResponseData) (ls []string) {
		for _, v := range in {
			ls = append(ls, v.ID)
		}
		return
	}(listResp.Data), func(in []*moonshot.FilesDeleteResponse) (ls []string) {
		for _, resp := range in {
			ls = append(ls, resp.ID)
		}
		return
	}(deleteAllResp.RespList), "must delete all files")
}
