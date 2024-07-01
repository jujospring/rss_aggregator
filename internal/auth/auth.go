package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extracts api key from the headers of an http request
// example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication information found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authentication")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first section of authentication header")
	}

	return vals[1], nil
}
