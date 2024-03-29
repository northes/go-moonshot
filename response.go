package moonshot

type CommonResponse struct {
	Code    int    `json:"code,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Method  string `json:"method,omitempty"`
	Scode   string `json:"scode,omitempty"`
	Status  bool   `json:"status,omitempty"`
	UA      string `json:"ua,omitempty"`
	URL     string `json:"url,omitempty"`
}

type CommonAPIResponse struct {
	Error *CommonAPIResponseError `json:"error"`
}

type CommonAPIResponseError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}
