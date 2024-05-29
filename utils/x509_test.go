package utils

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"github.com/tjfoc/gmsm/x509"
	"testing"

	"github.com/stretchr/testify/assert"
)

const certPEM = `MIIEGzCCA8CgAwIBAgIQMbQ+pdcKG4+QMTz13qJEpjAKBggqgRzPVQGDdTA0MQswCQYDVQQGEwJDTjERMA8GA1UECgwIVW5pVHJ1c3QxEjAQBgNVBAMMCVNIRUNBIFNNMjAeFw0yNDA0MTEwNTIyNDlaFw0yNDA1MTExNTU5NTlaMFkxCzAJBgNVBAYTAkNOMQswCQYDVQQIEwJTSDELMAkGA1UEBxMCU0gxCzAJBgNVBAoTAlNIMQswCQYDVQQLEwJTSDEWMBQGA1UEAwwN5rWL6K+V5rOV5Lq6MjBZMBMGByqGSM49AgEGCCqBHM9VAYItA0IABJr2N2yTLP0I2EragHXxQ4zu03A3WG/yj0ikzV4qUheeQVgLihngUcF7Cs3vzBjDgEYsQ6KXxeqKBceMIE/ZHpejggKNMIICiTAiBgNVHSMBAf8EGDAWgBSJMQSRe0Oqqpq/hB2bhu7wuHCZoDAgBgNVHQ4BAf8EFgQUAsTO3xRPEnLsaz0wpO2fCLKurakwDgYDVR0PAQH/BAQDAgbAMBMGA1UdJQQMMAoGCCsGAQUFBwMCMEIGA1UdIAQ7MDkwNwYJKoEcAYbvOoEVMCowKAYIKwYBBQUHAgEWHGh0dHA6Ly93d3cuc2hlY2EuY29tL3BvbGljeS8wCQYDVR0TBAIwADCBzwYDVR0fBIHHMIHEMHOgcaBvhm1sZGFwOi8vbGRhcDIuc2hlY2EuY29tOjM4OS8sY249OTMzZjRiYTczNmU3ZDljZC5jcmwsb3U9YTMyNTRiMzUsb3U9OTRmNCxvdT0wYWU5LG91PWQ0MTA0YjRkLG91PWNybCxvPVVuaVRydXN0ME2gS6BJhkdodHRwOi8vbGRhcDIuc2hlY2EuY29tL2Q0MTA0YjRkLzBhZTkvOTRmNC9hMzI1NGIzNS85MzNmNGJhNzM2ZTdkOWNkLmNybDCBgQYIKwYBBQUHAQEEdTBzMDgGCCsGAQUFBzABhixodHRwOi8vb2NzcDMuc2hlY2EuY29tL29jc3Avc2hlY2Evc2hlY2Eub2NzcDA3BggrBgEFBQcwAoYraHR0cDovL2xkYXAyLnNoZWNhLmNvbS9yb290L3NoZWNhc20yc3ViLmRlcjAbBggqgRzQFAQBAwQPEw0xMTExMTExMTExMTExMBAGCSqBHIbvOguBTgQDEwExMCAGCSqBHIbvOguBUgQTExFYWFgxMTFYWFhYMDAwMDIwMjAmBgkqgRyG7zoLgU0EGRMXMjAxQFhZWFhYMTExWFhYWDAwMDAyMDIwCgYIKoEcz1UBg3UDSQAwRgIhAKGnZ2IBltpagnk2rshCMJD6X5ZkIK/7fReBYQiwnFlpAiEA+ZZo/dIUjLrcmNOzgdQ4PoWLQnRWGvCU2i81k+a0kd4=`

func TestParseX509Certificate(t *testing.T) {
	t.Run("ValidCertificate", func(t *testing.T) {
		cert, err := ParseX509Certificate(certPEM)
		assert.NoError(t, err)
		assert.NotNil(t, cert)
	})

	t.Run("InvalidBase64", func(t *testing.T) {
		_, err := ParseX509Certificate("invalid-base64")
		assert.Error(t, err)
	})
}

func TestGetExtensionValue(t *testing.T) {
	// 正面测试用例：有效的扩展ID和值
	t.Run("有效扩展", func(t *testing.T) {
		extID := "1.2.3.4"
		extValue := []byte{0x30, 0x03, 0x01, 0x02, 0x03}
		cert := x509.Certificate{
			Extensions: []pkix.Extension{
				{Id: asn1.ObjectIdentifier{1, 2, 3, 4}, Value: extValue},
			},
		}
		expected := extValue[2:]
		result, err := GetExtensionValue(extID, cert)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !isEqual(result, expected) {
			t.Fatalf("expected %v, got %v", expected, result)
		}
	})

	// 负面测试用例：扩展ID未找到
	t.Run("扩展未找到", func(t *testing.T) {
		extID := "1.2.3.5"
		cert := x509.Certificate{
			Extensions: []pkix.Extension{
				{Id: asn1.ObjectIdentifier{1, 2, 3, 4}, Value: []byte{0x30, 0x03, 0x01, 0x02, 0x03}},
			},
		}
		_, err := GetExtensionValue(extID, cert)
		if err != errCertExtensionNotFound {
			t.Fatalf("expected error %v, got %v", errCertExtensionNotFound, err)
		}
	})

	// 负面测试用例：扩展值不是ASN.1编码
	t.Run("扩展值不是ASN.1编码", func(t *testing.T) {
		extID := "1.2.3.4"
		cert := x509.Certificate{
			Extensions: []pkix.Extension{
				{Id: asn1.ObjectIdentifier{1, 2, 3, 4}, Value: []byte{0x01}},
			},
		}
		_, err := GetExtensionValue(extID, cert)
		if err != errCertExtensionNotASN1Encoded {
			t.Fatalf("expected error %v, got %v", errCertExtensionNotASN1Encoded, err)
		}
	})

	// 边界情况：证书扩展为空
	t.Run("证书扩展为空", func(t *testing.T) {
		extID := "1.2.3.4"
		cert := x509.Certificate{
			Extensions: []pkix.Extension{},
		}
		_, err := GetExtensionValue(extID, cert)
		if err != errCertExtensionNotFound {
			t.Fatalf("expected error %v, got %v", errCertExtensionNotFound, err)
		}
	})

	// 边界情况：扩展值正好为2字节
	t.Run("扩展值正好为2字节", func(t *testing.T) {
		extID := "1.2.3.4"
		extValue := []byte{0x30, 0x02}
		cert := x509.Certificate{
			Extensions: []pkix.Extension{
				{Id: asn1.ObjectIdentifier{1, 2, 3, 4}, Value: extValue},
			},
		}
		expected := []byte{}
		result, err := GetExtensionValue(extID, cert)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !isEqual(result, expected) {
			t.Fatalf("expected %v, got %v", expected, result)
		}
	})
}

func isEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
