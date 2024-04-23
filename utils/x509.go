package utils

import (
	"encoding/base64"
	"errors"
	"github.com/tjfoc/gmsm/x509"
)

var (
	errCertExtensionNotFound = errors.New("cert extension not found")
)

func SM2Verify(raw, signature, pubKeyBytes []byte) bool {
	publicKey, err := x509.ParseSm2PublicKey(pubKeyBytes)
	if err != nil {
		return false
	}
	return publicKey.Verify(raw, signature)
}

func ParseX509Certificate(certBase64 string) (*x509.Certificate, error) {
	certBytes, err := base64.StdEncoding.DecodeString(certBase64)
	if err != nil {
		return nil, err
	}
	var cert *x509.Certificate
	cert, err = x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

func GetExtensionValue(id string, cert x509.Certificate) ([]byte, error) {
	for _, extension := range cert.Extensions {
		if id != extension.Id.String() {
			continue
		}
		return extension.Value, nil
	}
	return nil, errCertExtensionNotFound
}
