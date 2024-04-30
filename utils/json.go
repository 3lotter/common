package utils

import (
	"encoding/json"
	"reflect"
)

func JsonCompare(jsonA []byte, jsonB []byte) (bool, error) {
	credentialJsonMap := make(map[string]any)
	credentialSubjectJsonMap := make(map[string]any)
	if err := json.Unmarshal(jsonA, &credentialJsonMap); err != nil {
		return false, err
	}
	if err := json.Unmarshal(jsonB, &credentialSubjectJsonMap); err != nil {
		return false, err
	}

	return reflect.DeepEqual(credentialJsonMap, credentialSubjectJsonMap), nil
}
