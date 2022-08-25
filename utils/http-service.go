package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpService struct{}

func (s HttpService) Get() {
	r, _ := http.NewRequest(http.MethodGet, "<url>", nil)
	s.Do(r)
}

func (s HttpService) Post() {
	body, _ := json.Marshal(map[string]any{
		"username": "superadmin",
		"password": "123456",
	})
	r, _ := http.NewRequest(http.MethodPost, "<url>", bytes.NewBuffer(body))
	s.Do(r)
}

func (s HttpService) Do(r *http.Request) (map[string]any, error) {
	r.Header = http.Header{
		"Content-Type": {"application/json"},
		// "Authorization": {"Bearer <token>"},
	}

	res, err := (&http.Client{}).Do(r)
	if err != nil {
		return nil, err
	}

	buffer, _ := io.ReadAll(res.Body)
	fmt.Println(string(buffer))

	data := map[string]any{}
	err = json.Unmarshal(buffer, &data)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)

	return data, nil
}
