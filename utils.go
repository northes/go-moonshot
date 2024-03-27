package moonshot

import "encoding/json"

func MarshalToStringX(in any) string {
	b, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}
	return string(b)
}
