package backend

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HTTPBackend struct{}

func NewHttpBackend() *HTTPBackend {
	return &HTTPBackend{}
}

func (be *HTTPBackend) Load() error {
	return nil
}

func (be *HTTPBackend) Set(key, value string) error {
	urlStr := fmt.Sprintf("http://localhost:8080/kv/%s", key)
	bodyData := map[string]string{
		"value": value,
	}

	data, err := json.Marshal(bodyData)
	if err != nil {
		return err
	}

	buffer := bytes.NewBuffer(data)

	response, err := http.Post(urlStr, "application/json", buffer)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("failed to set")
	}

	return nil
}

func (be *HTTPBackend) Get(key string) (string, error) {
	urlStr := fmt.Sprintf("http://localhost:8080/kv/%s", key)

	response, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New("failed to get")
	}

	respBody := map[string]string{}

	err = json.NewDecoder(response.Body).Decode(&respBody)
	if err != nil {
		return "", err
	}

	value, ok := respBody["value"]
	if !ok {
		return "", errors.New("missing 'value' key on response body")
	}

	return value, nil
}

func (be *HTTPBackend) Delete(string) error {
	return nil
}
