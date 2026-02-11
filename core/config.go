package core

import (
	"crypto/tls"
	"crypto/x509"
)

type Config struct {
	APIKey          string
	APISecret       string
	Target          string
	Scheme          string
	TLSClientConfig *TLSClientConfig

	LogLevel LogLevel
	Logger   Logger
}

type TLSClientConfig struct {
	// Server should be accessed without verifying the TLS certificate. For testing only.
	InsecureSkipVerify bool

	// Server requires TLS client certificate authentication
	CertFile string
	// Server requires TLS client certificate authentication
	KeyFile string
	// Trusted root certificates for server
	CAFile string

	// CertData holds PEM-encoded bytes (typically read from a client certificate file).
	// CertData takes precedence over CertFile
	CertData []byte
	// KeyData holds PEM-encoded bytes (typically read from a client certificate key file).
	// KeyData takes precedence over KeyFile
	KeyData []byte
	// CAData holds PEM-encoded bytes (typically read from a root certificates bundle).
	// CAData takes precedence over CAFile
	CAData []byte
}

func (conf *TLSClientConfig) ToTLSConfig() *tls.Config {
	if conf == nil {
		return nil
	}
	certpool := x509.NewCertPool()
	certpool.AppendCertsFromPEM(conf.CAData)
	clientKeyPair, _ := tls.X509KeyPair(conf.CertData, conf.KeyData)
	return &tls.Config{
		InsecureSkipVerify: conf.InsecureSkipVerify,
		RootCAs:            certpool,
		Certificates:       []tls.Certificate{clientKeyPair},
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
	}
}
