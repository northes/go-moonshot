package moonshot_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/northes/go-moonshot"
)

func TestEnumChatCompletions(t *testing.T) {
	tt := require.New(t)

	tt.EqualValues(moonshot.RoleSystem, moonshot.RoleSystem.String())
	tt.EqualValues(moonshot.RoleUser, moonshot.RoleUser.String())
	tt.EqualValues(moonshot.RoleAssistant, moonshot.RoleAssistant.String())
	tt.EqualValues(moonshot.RoleTool, moonshot.RoleTool.String())

	tt.EqualValues(moonshot.ModelMoonshotV18K, moonshot.ModelMoonshotV18K.String())
	tt.EqualValues(moonshot.ModelMoonshotV132K, moonshot.ModelMoonshotV132K.String())
	tt.EqualValues(moonshot.ModelMoonshotV1128K, moonshot.ModelMoonshotV1128K.String())

	tt.EqualValues(moonshot.FinishReasonStop, moonshot.FinishReasonStop.String())
	tt.EqualValues(moonshot.FinishReasonLength, moonshot.FinishReasonLength.String())

	tt.EqualValues(moonshot.ChatCompletionsToolTypeFunction, moonshot.ChatCompletionsToolTypeFunction.String())

	tt.EqualValues(moonshot.ChatCompletionsParametersTypeObject, moonshot.ChatCompletionsParametersTypeObject.String())

	tt.EqualValues(moonshot.ChatCompletionsResponseFormatJSONObject, moonshot.ChatCompletionsResponseFormatJSONObject.String())
	tt.EqualValues(moonshot.ChatCompletionsResponseFormatText, moonshot.ChatCompletionsResponseFormatText.String())
}
