package util

import (
	"encoding/json"
)

func CheckErr(err error) {
	// This panics if err is not nil which renders a 500 error.
	if err != nil {
		panic(err)
	}
}

func JSON(data map[string]interface{}) []byte {
	b, err := json.Marshal(data)
	CheckErr(err)
	return b
}
