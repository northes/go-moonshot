package moonshot_test

import (
	"context"
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
	resp, err := cli.Files().Upload(context.Background(), &moonshot.FilesUploadRequest{
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
	uploadResp, err := cli.Files().UploadBytes(context.Background(), &moonshot.FilesUploadBytesRequest{
		Name:    "byteFile.txt",
		Bytes:   content,
		Purpose: moonshot.FilePurposeExtract,
	})
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(uploadResp.Bytes, len(content))

	// check content
	contentResp, err := cli.Files().Content(context.Background(), uploadResp.ID)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(string(content), contentResp.Content)

	// remove
	deleteResp, err := cli.Files().Delete(context.Background(), uploadResp.ID)
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
	resp, err := cli.Files().List(context.Background())
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
	resp, err := cli.Files().Info(context.Background(), fileID)
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
	resp, err := cli.Files().Content(context.Background(), fileID)
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
	resp, err := cli.Files().Delete(context.Background(), fileID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", test.MarshalJsonToStringX(resp))
}

func TestFilesBatchDelete(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}

	fileIdList := make([]string, 0)

	for i := 0; i < 10; i++ {
		fp, _ := test.GenerateTestFile(test.GenerateTestContent())
		uploadResp, err := cli.Files().Upload(context.Background(), &moonshot.FilesUploadRequest{
			Name:    filepath.Base(fp),
			Path:    fp,
			Purpose: moonshot.FilePurposeExtract,
		})
		if err != nil {
			t.Logf("upload file err: %v", err)
			continue
		}
		fileIdList = append(fileIdList, uploadResp.ID)
	}

	t.Logf("file id to delete: %v", fileIdList)

	deleteAllResp, err := cli.Files().BatchDelete(context.Background(), &moonshot.FilesBatchDeleteRequest{
		FileIDList: fileIdList,
	})
	if err != nil {
		t.Fatal(err)
	}

	require.EqualValues(t, func(in []string) (ls []string) {
		for _, v := range in {
			ls = append(ls, v)
		}
		return
	}(fileIdList), func(in []*moonshot.FilesDeleteResponse) (ls []string) {
		for _, resp := range in {
			ls = append(ls, resp.ID)
		}
		return
	}(deleteAllResp.RespList), "must delete all files")

	t.Logf("deleted Files ID List: %v", func(ls []*moonshot.FilesDeleteResponse) (l []string) {
		for _, v := range ls {
			l = append(l, v.ID)
		}
		return l
	}(deleteAllResp.RespList))
}
