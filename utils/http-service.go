package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpService struct{}

func (s *HttpService) Get(url string) {
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	s.Do(r)
}

func (s *HttpService) Post(url string) {
	b, _ := json.Marshal(map[string]any{
		"username": "superadmin",
		"password": "123456",
	})
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	s.Do(r)
}

func (s *HttpService) Do(r *http.Request) (data map[string]any, err error) {
	r.Header = http.Header{
		"Content-Type": {"application/json"},
		// "Authorization": {"Bearer <token>"},
	}

	res, err := (&http.Client{}).Do(r)
	if err != nil {
		return
	}

	b, _ := io.ReadAll(res.Body)
	fmt.Println(string(b))

	json.Unmarshal(b, &data)
	return
}
