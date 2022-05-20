package constant

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/spf13/viper"
)

type RequestParam struct {
	Method  string
	URL     string
	Body    io.Reader
	Header  http.Header
	TimeOut time.Duration
}

func Request(param RequestParam) ([]byte, error) {
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
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[%s]%s returns non-200 status: %s", param.Method, param.URL, resp.Status)
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}

type Service string

const (
	ServiceLine Service = "line"
	ServiceUser Service = "user"
)

func RequestToService(s Service, method string, api string, body io.Reader, params interface{}) ([]byte, error) {
	var serviceURL string
	switch s {
	case ServiceUser:
		serviceURL = "http://localhost:8000"
	case ServiceLine:
		serviceURL = "http://localhost:8001"
	}
	uService, err := url.Parse(serviceURL)
	if err != nil {
		return nil, err
	}

	u, err := uService.Parse(api)
	if err != nil {
		return nil, err
	}

	if params != nil {
		val, err := query.Values(params)
		if err != nil {
			return nil, err
		}
		u.RawQuery = val.Encode()
	}

	log.Println(u.String())

	reqParam := RequestParam{
		Method: method,
		URL:    u.String(),
		Body:   body,
	}

	return Request(reqParam)
}
