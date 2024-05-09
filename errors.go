package moonshot

import (
	"fmt"

	"github.com/northes/go-moonshot/internal/httpx"
)

func ResponseToError(resp *httpx.Response) error {
	errResp := new(CommonAPIResponse)
	err := resp.Unmarshal(errResp)
	if err != nil {
		return err
	}
	return fmt.Errorf("[%s]%s", errResp.Error.Type, errResp.Error.Message)
}
