package main

import (
	"errors"
	"net/http"
	"strings"
)

func getApiKey(headers http.Header) (string, error) {
	value := headers.Get("Authorization")
	if value == "" {
		return "", errors.New("missing Authorization header")
	}

	keys := strings.Split(value, " ")
	if len(keys) != 2 {
		return "", errors.New("invalid Authorization header")
	}
	if keys[0] != "Bearer" {
		return "", errors.New("invalid Authorization header")
	}

	return keys[1], nil
}
