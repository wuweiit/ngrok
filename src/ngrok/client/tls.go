package client

import (
	_ "crypto/sha512"

	"ngrok/log"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
    "io/ioutil"
	"ngrok/client/assets"
)

func LoadTLSConfig(rootCertPaths []string) (*tls.Config, error) {
	pool := x509.NewCertPool()

	mlog := log.NewPrefixLogger("client")

	for _, certPath := range rootCertPaths { 


		rootCrt1 := []byte("")
		if certPath == "assets/client/tls/ngrokroot.crt" {
			rootCrt, err := assets.Asset(certPath)
			if err != nil {
				return nil, err
			}
			rootCrt1 = rootCrt; 
		} else {
			// marker 2019-11-12 新增读取固定的配置
			rootCrt, err := ioutil.ReadFile(certPath)
			if err != nil {
				fmt.Println("File reading error", err)
				return nil, fmt.Errorf("File reading error")
			} 
			rootCrt1 = rootCrt; 
		}
		

		
		mlog.Info("Contents of file: %v", string(rootCrt1))

		pemBlock, _ := pem.Decode(rootCrt1)
		if pemBlock == nil {
			return nil, fmt.Errorf("Bad PEM data")
		}

		certs, err := x509.ParseCertificates(pemBlock.Bytes)
		if err != nil {
			return nil, err
		}

		pool.AddCert(certs[0])
	}

	return &tls.Config{RootCAs: pool}, nil
}
