package moonshot_test

import (
	"testing"

	"github.com/northes/go-moonshot"
	"github.com/stretchr/testify/require"
)

func TestEnumChatCompletions(t *testing.T) {
	tt := require.New(t)

	tt.EqualValues(moonshot.RoleSystem, moonshot.RoleSystem.String())
	tt.EqualValues(moonshot.RoleUser, moonshot.RoleUser.String())
	tt.EqualValues(moonshot.RoleAssistant, moonshot.RoleAssistant.String())

	tt.EqualValues(moonshot.ModelMoonshotV18K, moonshot.ModelMoonshotV18K.String())
	tt.EqualValues(moonshot.ModelMoonshotV132K, moonshot.ModelMoonshotV132K.String())
	tt.EqualValues(moonshot.ModelMoonshotV1128K, moonshot.ModelMoonshotV1128K.String())

	tt.EqualValues(moonshot.FinishReasonStop, moonshot.FinishReasonStop.String())
	tt.EqualValues(moonshot.FinishReasonLength, moonshot.FinishReasonLength.String())
}
