package moonshot

import (
	"bytes"
	"fmt"
	"os"

	utils "github.com/northes/go-moonshot/internal/builder"
)

type files struct {
	client *Client
}

func (c *Client) Files() *files {
	return &files{
		client: c,
	}
}

type FilesUploadRequest struct {
	Name    string
	Path    string
	Purpose FilesPurpose
}
type FilesUploadResponse struct {
	ID            string `json:"id"`
	Object        string `json:"object"`
	Bytes         int    `json:"bytes"`
	CreatedAt     int    `json:"created_at"`
	Filename      string `json:"filename"`
	Purpose       string `json:"purpose"`
	Status        string `json:"status"`
	StatusDetails string `json:"status_details"`
}

func (f *files) Upload(req *FilesUploadRequest) (*FilesUploadResponse, error) {
	const path = "/v1/files"

	var b bytes.Buffer

	builder := utils.NewFormBuilder(&b)
	err := builder.WriteField("purpose", FilePurposeExtract.String())
	if err != nil {
		return nil, err
	}
	fileData, err := os.Open(req.Path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(fileData)
	err = builder.CreateFormFile("file", fileData)
	if err != nil {
		return nil, err
	}
	err = builder.Close()
	if err != nil {
		return nil, err
	}

	resp, err := f.client.HTTPClient().
		SetPath(path).
		SetBody(&b).
		SetContentType(builder.FormDataContentType()).
		Post()
	if err != nil {
		return nil, err
	}

	if !resp.StatusOK() {
		return nil, StatusCodeToError(resp.Raw().StatusCode)
	}

	uploadResponse := new(FilesUploadResponse)
	err = resp.Unmarshal(uploadResponse)
	if err != nil {
		return nil, err
	}

	return uploadResponse, nil
}

type FilesUploadBytesRequest struct {
	Name    string
	Bytes   []byte
	Purpose FilesPurpose
}
type FilesUploadBytesResponse struct {
	ID            string `json:"id"`
	Object        string `json:"object"`
	Bytes         int    `json:"bytes"`
	CreatedAt     int    `json:"created_at"`
	Filename      string `json:"filename"`
	Purpose       string `json:"purpose"`
	Status        string `json:"status"`
	StatusDetails string `json:"status_details"`
}

func (f *files) UploadBytes(req *FilesUploadBytesRequest) (*FilesUploadBytesResponse, error) {
	const path = "/v1/files"

	var b bytes.Buffer
	reader := bytes.NewReader(req.Bytes)

	builder := utils.NewFormBuilder(&b)
	err := builder.WriteField("purpose ", FilePurposeExtract.String())
	if err != nil {
		return nil, err
	}
	err = builder.CreateFormFileReader("file", reader, req.Name)
	if err != nil {
		return nil, err
	}
	err = builder.Close()
	if err != nil {
		return nil, err
	}

	resp, err := f.client.HTTPClient().
		SetPath(path).
		SetBody(&b).
		SetContentType(builder.FormDataContentType()).
		Post()
	if err != nil {
		return nil, err
	}

	if !resp.StatusOK() {
		return nil, StatusCodeToError(resp.Raw().StatusCode)
	}

	uploadResponse := new(FilesUploadBytesResponse)
	err = resp.Unmarshal(uploadResponse)
	if err != nil {
		return nil, err
	}

	return uploadResponse, nil
}

type FilesListRequest struct {
}
type FilesListResponse struct {
	Object string                   `json:"object"`
	Data   []*FilesListResponseData `json:"data"`
}
type FilesListResponseData struct {
	ID           string       `json:"id"`
	Object       string       `json:"object"`
	Bytes        int64        `json:"bytes"`
	CreatedAt    int64        `json:"created_at"`
	Filename     string       `json:"filename"`
	Purpose      FilesPurpose `json:"purpose"`
	Status       string       `json:"status"`
	StatusDetail string       `json:"status_detail"`
}

func (f *files) Lists() (*FilesListResponse, error) {
	const path = "/v1/files"
	resp, err := f.client.HTTPClient().SetPath(path).Get()
	if err != nil {
		return nil, err
	}
	if !resp.StatusOK() {
		return nil, StatusCodeToError(resp.Raw().StatusCode)
	}
	listResponse := new(FilesListResponse)
	err = resp.Unmarshal(listResponse)
	if err != nil {
		return nil, err
	}
	return listResponse, nil
}

type FilesDeleteResponse struct {
	CommonAPIResponse
	Deleted bool   `json:"deleted"`
	ID      string `json:"id"`
	Object  string `json:"object"`
}

func (f *files) Delete(fileID string) (*FilesDeleteResponse, error) {
	const path = "/v1/files/%s"
	fullPath := fmt.Sprintf(path, fileID)
	resp, err := f.client.HTTPClient().SetPath(fullPath).Delete()
	if err != nil {
		return nil, err
	}
	if !resp.StatusOK() {
		return nil, StatusCodeToError(resp.Raw().StatusCode)
	}
	deleteResponse := new(FilesDeleteResponse)
	err = resp.Unmarshal(deleteResponse)
	if err != nil {
		return nil, err
	}
	return deleteResponse, nil
}

type FilesBatchDeleteRequest struct {
	FileIDList []string `json:"file_ids"`
}
type FilesBatchDeleteResponse struct {
	RespList []*FilesDeleteResponse `json:"resp_list"`
}

func (f *files) BatchDelete(req *FilesBatchDeleteRequest) (*FilesBatchDeleteResponse, error) {
	if req == nil || len(req.FileIDList) == 0 {
		return nil, fmt.Errorf("batch delete request must contain at least one file id")
	}

	deleteAllResp := &FilesBatchDeleteResponse{
		RespList: make([]*FilesDeleteResponse, 0),
	}

	for _, fileID := range req.FileIDList {
		deleteResp, err := f.Delete(fileID)
		if err != nil {
			return nil, err
		}
		deleteAllResp.RespList = append(deleteAllResp.RespList, deleteResp)
	}

	return deleteAllResp, nil
}

type FilesInfoResponse struct {
	Id            string `json:"id"`
	Object        string `json:"object"`
	Bytes         int    `json:"bytes"`
	CreatedAt     int    `json:"created_at"`
	Filename      string `json:"filename"`
	Purpose       string `json:"purpose"`
	Status        string `json:"status"`
	StatusDetails string `json:"status_details"`
}

func (f *files) Info(fileID string) (*FilesInfoResponse, error) {
	const path = "/v1/files/%s"
	fullPath := fmt.Sprintf(path, fileID)
	resp, err := f.client.HTTPClient().SetPath(fullPath).Get()
	if err != nil {
		return nil, err
	}
	if !resp.StatusOK() {
		return nil, StatusCodeToError(resp.Raw().StatusCode)
	}
	infoResponse := new(FilesInfoResponse)
	err = resp.Unmarshal(infoResponse)
	if err != nil {
		return nil, err
	}
	return infoResponse, nil
}

type FileContentResponse struct {
	Content  string `json:"content"`
	FileType string `json:"file_type"`
	Filename string `json:"filename"`
	Title    string `json:"title"`
	Type     string `json:"type"`
}

func (f *files) Content(fileID string) (*FileContentResponse, error) {
	const path = "/v1/files/%s/content"
	fullPath := fmt.Sprintf(path, fileID)
	resp, err := f.client.HTTPClient().SetPath(fullPath).Get()
	if err != nil {
		return nil, err
	}
	if !resp.StatusOK() {
		return nil, StatusCodeToError(resp.Raw().StatusCode)
	}
	contentResponse := new(FileContentResponse)
	err = resp.Unmarshal(contentResponse)
	if err != nil {
		return nil, err
	}
	return contentResponse, nil
}
