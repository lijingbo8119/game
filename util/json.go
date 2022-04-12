package util

import "encoding/json"

func JsonMustMarshal(v any) []byte {
	d, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return d
}

func JsonMustMarshalString(v any) string {
	return string(JsonMustMarshal(v))
}
