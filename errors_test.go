package moonshot_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/northes/go-moonshot"
	"github.com/stretchr/testify/require"
)

func TestErrors(t *testing.T) {
	err1 := moonshot.ErrorInvalidRequest
	codeErr1 := moonshot.StatusCodeToError(400)
	require.ErrorIs(t, codeErr1, err1)
}

func TestErrorsCodeNotExist(t *testing.T) {
	code := 200
	err := moonshot.StatusCodeToError(code)
	require.Equal(t, err, fmt.Errorf("[%d] %s", code, http.StatusText(code)))
}
