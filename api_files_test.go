package moonshot_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/northes/go-moonshot"
	"github.com/northes/go-moonshot/test"
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
	resp, err := cli.Files().Upload(&moonshot.FilesUploadRequest{
		Name:    filepath.Base(filePath),
		Path:    filePath,
		Purpose: moonshot.FilePurposeExtract,
	})
	if err != nil {
		t.Fatal(err)
	}
	fileID = resp.ID
	t.Logf("%+v", test.MarshalJsonToStringX(resp))
	err = os.Remove(filePath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFilesUploadBytes(t *testing.T) {
	content := []byte("hello")
	require := require.New(t)
	// upload
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	uploadResp, err := cli.Files().UploadBytes(&moonshot.FilesUploadBytesRequest{
		Name:    "byteFile.txt",
		Bytes:   content,
		Purpose: moonshot.FilePurposeExtract,
	})
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(uploadResp.Bytes, len(content))

	// check content
	contentResp, err := cli.Files().Content(uploadResp.ID)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(string(content), contentResp.Content)

	// remove
	deleteResp, err := cli.Files().Delete(uploadResp.ID)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(deleteResp.ID, uploadResp.ID)
	require.Equal(deleteResp.Deleted, true)
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
	t.Logf("%+v", test.MarshalJsonToStringX(resp))
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
	t.Logf("%+v", test.MarshalJsonToStringX(resp))
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

	t.Logf("%+v", test.MarshalJsonToStringX(resp))
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
	t.Logf("%+v", test.MarshalJsonToStringX(resp))
}

/*
❗ This test may lead to unexpected results
*/
func TestFilesDeleteAll(t *testing.T) {

	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}

	for range 10 {
		fp, _ := test.GenerateTestFile(test.GenerateTestContent())
		_, _ = cli.Files().Upload(&moonshot.FilesUploadRequest{
			Name:    filepath.Base(fp),
			Path:    fp,
			Purpose: moonshot.FilePurposeExtract,
		})
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

	t.Logf("Deleted Files ID List: %v", func(ls []*moonshot.FilesDeleteResponse) (l []string) {
		for _, v := range ls {
			l = append(l, v.ID)
		}
		return l
	}(deleteAllResp.RespList))
}
