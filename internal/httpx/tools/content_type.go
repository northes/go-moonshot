package tools

import (
	"strings"
)

type ContentType string

const (
	BaseApplication = "application"

	TextHtml  ContentType = "text/html"
	TextPlain ContentType = "text/plain"
	TextXml   ContentType = "text/xml"

	ImageGif  ContentType = "image/gif"
	ImageJpeg ContentType = "image/jpeg"
	ImagePng  ContentType = "image/png"

	ApplicationXhtmlXml           ContentType = "application/xhtml+xml"
	ApplicationXml                ContentType = "application/xml"
	ApplicationAtomXml            ContentType = "application/atom+xml"
	ApplicationJson               ContentType = "application/json"
	ApplicationPDF                ContentType = "application/pdf"
	ApplicationMsWord             ContentType = "application/msword"
	ApplicationOctetStream        ContentType = "application/octet-stream"
	ApplicationXWwwFormUrlencoded ContentType = "application/x-www-form-urlencoded"

	MultipartFormData ContentType = "multipart/form-data"
)

func (c ContentType) String() string {
	return string(c)
}

// NewContentType 换取ContentType
func NewContentType(subType string) string {
	return strings.Join([]string{BaseApplication, subType}, "/")
}

// ContentSubType 解析ContentType
func ContentSubType(contentType ContentType) string {
	contentTypeStr := contentType.String()
	left := strings.Index(contentTypeStr, "/")
	if left == -1 {
		return ""
	}
	right := strings.Index(contentTypeStr, ";")
	if right == -1 {
		right = len(contentTypeStr)
	}
	if right < left {
		return ""
	}
	return contentTypeStr[left+1 : right]
}
