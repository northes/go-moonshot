package test

import "encoding/json"

func MarshalJsonToStringX(t any) string {
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(b)
}
