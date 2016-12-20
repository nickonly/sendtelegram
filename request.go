package main

import (
	"net/http"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/url"
	"strings"
)

func Request(token string, method string, values url.Values) ([]byte, error) {
	url := "https://api.telegram.org/bot" + token + "/" + method
	pool := x509.NewCertPool()
	tr := &http.Transport{
		DisableCompression: true,
		TLSClientConfig: &tls.Config{RootCAs: pool, InsecureSkipVerify : true},
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/50.0.2661.102 Chrome/50.0.2661.102 Safari/537.36")
	response, err := tr.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return contents, nil
}


