package utils

import (
	"crypto/tls"
	"math/rand"
	"net/http"
	"time"
)

type BaseConfig struct {
	IdTenantUrl string `env:"ID_TENANT_URL,required"`
	PcloudUrl   string `env:"PCLOUD_URL,required"`
	SafeName    string `env:"SAFE_NAME,required"`
	PlatformID  string `env:"PLATFORM_ID,required"`
	PamUser     string `env:"PAM_USER,required"`
	PamPass     string `env:"PAM_PASS,required,unset"`

	TlsSkipVerify bool `env:"TLS_SKIP_VERIFY" envDefault:"false"`
}

// GetHTTPClient create http client for HTTPS
func GetHTTPClient(timeout time.Duration, skipverify bool) *http.Client {
	client := &http.Client{
		Timeout: timeout, /*time.Second * 30 */
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: skipverify, /* TLS_SKIP_VERIFY */
			},
		},
	}
	return client
}

func TrimQuotes(s string) string {
	x := s[:]
	if len(x) >= 2 {
		if x[0] == '"' && x[len(x)-1] == '"' {
			return x[1 : len(x)-1]
		}
	}
	return x
}

func FirstN(str string, n int) string {
	v := []rune(str)
	if n >= len(v) {
		return str
	}
	return string(v[:n])
}

func MinInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func RandSeq(charlist []rune, numchars int) string {
	b := make([]rune, numchars)
	for i := range b {
		b[i] = charlist[rand.Intn(len(charlist))]
	}
	return string(b)
}
