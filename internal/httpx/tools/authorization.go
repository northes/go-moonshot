package tools

import (
	"errors"
	"fmt"
	"strings"
)

var ErrorNotABBearerToken = errors.New("错误的 Bearer Token 格式")

func ParseBearerToken(authorHead string) (token string, err error) {
	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrorNotABBearerToken
	}
	return parts[1], nil
}

func ToBearToken(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}
