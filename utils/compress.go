package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
)

func CompressString(raw string) (string, error) {
	var buffer bytes.Buffer
	gz := gzip.NewWriter(&buffer)
	if _, err := gz.Write([]byte(raw)); err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

func DeCompressString(compressed string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(compressed)
	if err != nil {
		return "", err
	}

	var gr *gzip.Reader
	gr, err = gzip.NewReader(bytes.NewBuffer(decodedBytes))
	if err != nil {
		return "", err
	}

	var decompressedBytes []byte
	decompressedBytes, err = io.ReadAll(gr)
	if err != nil {
		return "", err
	}

	return string(decompressedBytes), nil
}
