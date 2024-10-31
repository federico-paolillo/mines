package id

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
)

const (
	idSizeInBytes = 256
	emptyId       = ""
)

func Generate() (string, error) {
	rawId := make([]byte, idSizeInBytes)

	_, err := rand.Read(rawId)
	if err != nil {
		//nolint:errorlint // We do not want to wrap and leak errors that are not under our control
		return emptyId, fmt.Errorf(
			"id: could not generate id random bytes. %v",
			err,
		)
	}

	var sb strings.Builder

	encoder := base64.NewEncoder(base64.URLEncoding, &sb)

	_, err = encoder.Write(rawId)
	if err != nil {
		//nolint:errorlint // We do not want to wrap and leak errors that are not under our control
		return emptyId, fmt.Errorf(
			"id: could not base64 encode id. %v",
			err,
		)
	}

	err = encoder.Close()
	if err != nil {
		//nolint:errorlint // We do not want to wrap and leak errors that are not under our control
		return emptyId, fmt.Errorf(
			"id: could not base64 encode id. %v",
			err,
		)
	}

	return sb.String(), nil
}
