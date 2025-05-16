package utils

import (
	"encoding/json"
	"fmt"
)

func Reunmarshal[T any](data interface{}) (*T, error) {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("unamle to marshal data")
	}

	var res T
	err = json.Unmarshal([]byte(marshalledData), &res)
	if err != nil {
		return nil, fmt.Errorf("unamle to unmarshal data")
	}

	return &res, nil
}
