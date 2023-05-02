package helpers

import "encoding/json"

func MakeJson(toJson any) ([]byte, error) {
	return json.Marshal(toJson)
}
