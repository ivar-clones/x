package util

import (
	"bytes"
	"encoding/json"
	"testing"
)

func AssertJSON(actual interface{}, expected interface{}, t *testing.T) {
	actualData, err := json.Marshal(actual)
	if err != nil {
		t.Errorf("an error '%s' was not expected when marshaling actual json data", err)
	}

	expectedData, err := json.Marshal(expected)
	if err != nil {
		t.Errorf("an error '%s' was not expected when marshaling expected json data", err)
	}

	if !bytes.Equal(expectedData, actualData) {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}