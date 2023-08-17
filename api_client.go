package amadeusgo

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

func apiRequest(url string, method string, bodyParams interface{}, headers map[string]string) ([]byte, int) {
	client := &http.Client{}
	var err error
	var req *http.Request
	var reader io.ReadCloser

	if method == "POST" {
		req, err = http.NewRequest(method, url, bytes.NewBuffer([]byte(bodyParams.(string))))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		fmt.Println(err)
		return nil, 500
	}
	req.Header.Add("Accept", "application/json")
	for key, val := range headers {
		req.Header.Add(key, val)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, 500
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(res.Body)

	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(res.Body)
		defer func(reader io.ReadCloser) {
			err := reader.Close()
			if err != nil {
				fmt.Println(err.Error())
			}
		}(reader)
	default:
		reader = res.Body
	}
	body, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		return nil, 500
	}

	return body, res.StatusCode
}

func writeResponse(response http.ResponseWriter, result []byte, status int) {
	response.WriteHeader(status)
	_, err := response.Write(result)
	if err != nil {
		fmt.Println(err.Error())
	}
}
