package utils

import "encoding/json"

func GetFormatJSON(data interface{}) []byte {
	myIn, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return json.RawMessage(myIn)
}
