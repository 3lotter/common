package utils

import (
	"encoding/base64"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

type DecryptMode int

const (
	C1C3C2 DecryptMode = iota
	C1C2C3
)

var (
	errCertExtensionNotFound       = errors.New("cert extension not found")
	errCertExtensionNotASN1Encoded = errors.New("cert extension not asn.1 encoded")
)

func SM2Verify(raw, signature, pubKeyBytes []byte) bool {
	publicKey, err := x509.ParseSm2PublicKey(pubKeyBytes)
	if err != nil {
		return false
	}
	return publicKey.Verify(raw, signature)
}

func SM2Decrypt(keyBase64, cipherTextBase64 string, mode DecryptMode) ([]byte, error) {
	// 实例化私钥
	keyBytes, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return nil, errors.Wrap(err, "base64.StdEncoding.DecodeString fail")
	}

	var sm2Key *sm2.PrivateKey
	sm2Key, err = x509.ParsePKCS8UnecryptedPrivateKey(keyBytes)
	if err != nil {
		return nil, errors.Wrap(err, "x509.ParseSm2PrivateKey fail")
	}

	// 密文解密
	var cipherText []byte
	cipherText, err = base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return nil, errors.Wrap(err, "base64.StdEncoding.DecodeString fail")
	}

	var plainText []byte
	plainText, err = sm2.Decrypt(sm2Key, cipherText, int(mode))
	if err != nil {
		return nil, errors.Wrap(err, "sm2.Decrypt fail")
	}

	return plainText, nil
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
		if len(extension.Value) < 2 {
			return nil, errCertExtensionNotASN1Encoded
		}
		return extension.Value[2:], nil
	}
	return nil, errCertExtensionNotFound
}
