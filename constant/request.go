package constant

import (
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type RequestParam struct {
	Method  string
	URL     string
	Body    io.Reader
	Header  http.Header
	TimeOut time.Duration
}

func Request(param RequestParam) (io.ReadCloser, error) {
	req, err := http.NewRequest(param.Method, param.URL, param.Body)
	if err != nil {
		return nil, err
	}
	for name, value := range param.Header {
		req.Header.Set(name, strings.Join(value, ","))
	}
	if param.TimeOut == 0 {
		param.TimeOut = time.Duration(viper.GetInt("REQUEST_TIMEOUT"))
	}
	client := &http.Client{
		Timeout: param.TimeOut * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   time.Duration(param.TimeOut) * time.Second,
				KeepAlive: time.Duration(param.TimeOut) * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   time.Duration(param.TimeOut) * time.Second,
			ResponseHeaderTimeout: time.Duration(param.TimeOut) * time.Second,
			ExpectContinueTimeout: time.Duration(param.TimeOut) * time.Second,
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
