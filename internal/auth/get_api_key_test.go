package auth

import (
	"net/http"
	"testing"
)

func Test_GetAPIKey_NoHeader(t *testing.T) {
	// arrange
	emptyHeader := http.Header{}
	// (expects)
	expectedError := ErrNoAuthHeaderIncluded

	// act
	_, actualError := GetAPIKey(emptyHeader)

	// assert
	if expectedError == actualError {
		t.Errorf("GetAPIKey returned an unexpected error.")
	}

}

func Test_GetAPIKey_Header(t *testing.T) {
	// arrange
	emptyHeader := http.Header{}

	expectedKey := ""

	// act
	actualKey, _ := GetAPIKey(emptyHeader)

	// assert
	if actualKey != expectedKey {
		t.Errorf("GetAPIKey should have returned empty key.")
	}
}
