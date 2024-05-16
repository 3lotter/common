package utils

import (
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
	cert, _ := ParseX509Certificate(certPEM)

	t.Run("ValidExtension", func(t *testing.T) {
		// 假设我们知道证书中有一个扩展，其ID为"2.5.29.14"（subjectKeyIdentifier）
		extValue, err := GetExtensionValue("2.5.29.14", *cert)
		assert.NoError(t, err)
		assert.NotNil(t, extValue)
	})

	t.Run("InvalidExtension", func(t *testing.T) {
		_, err := GetExtensionValue("invalid-id", *cert)
		assert.Error(t, err)
	})
}
