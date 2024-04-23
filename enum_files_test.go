package moonshot_test

import (
	"testing"

	"github.com/northes/go-moonshot"
	"github.com/stretchr/testify/require"
)

func TestEnumFiles(t *testing.T) {
	tt := require.New(t)

	tt.EqualValues(moonshot.FilePurposeExtract, moonshot.FilePurposeExtract.String())
}
