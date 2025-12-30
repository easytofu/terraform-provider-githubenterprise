package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// ReadAndRestoreResponseBody reads the response body and restores it so it can be read again.
func ReadAndRestoreResponseBody(resp *http.Response) ([]byte, error) {
	if resp == nil || resp.Body == nil {
		return nil, nil
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, nil
}

// DecodeResponseJSON reads, restores, and JSON-decodes the response body.
//
// If the response has an empty body, this returns (nil, nil).
func DecodeResponseJSON(resp *http.Response) (any, error) {
	bodyBytes, err := ReadAndRestoreResponseBody(resp)
	if err != nil {
		return nil, err
	}

	if len(bytes.TrimSpace(bodyBytes)) == 0 {
		return nil, nil
	}

	var decoded any
	if err := json.Unmarshal(bodyBytes, &decoded); err != nil {
		return nil, err
	}

	return decoded, nil
}
