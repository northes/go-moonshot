package moonshot

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrorInvalidAuthentication                  = errors.New("鉴权失败请确认")
	ErrorInvalidRequest                         = errors.New("这个通常意味着您输入格式有误，包括使用了预期外的参数，比如过大的 temperature，或者 messages 的大小超过了限制。在 message 字段通常会有更多解释")
	ErrorRateLimitReachedOrExceededCurrentQuota = errors.New("您超速了。我们设置了最大并发上限和分钟为单位的次数限制。如果在 429 后立即重试，可能会遇到罚时建议控制并发大小，并且在 429 后 sleep 3 秒。或者是Quota 不够了，请联系管理员加量")
)

var statusCodeErrorMap map[int]error

func init() {
	statusCodeErrorMap = map[int]error{
		400: ErrorInvalidRequest,
		401: ErrorInvalidAuthentication,
		429: ErrorRateLimitReachedOrExceededCurrentQuota,
	}
}

func StatusCodeToError(code int) error {
	if err, ok := statusCodeErrorMap[code]; ok {
		return err
	}
	return fmt.Errorf("[%d] %s", code, http.StatusText(code))
}
