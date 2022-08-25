package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HttpGet() (map[string]any, error) {
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "<url>", nil)
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		// "Authorization": {"Bearer <token>"},
	}
	res, err := client.Do(req)
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

func HttpPost() (map[string]any, error) {
	client := http.Client{}
	body, _ := json.Marshal(map[string]any{
		"username": "superadmin",
		"password": "123456",
	})
	req, _ := http.NewRequest(http.MethodPost, "<url>", bytes.NewBuffer(body))
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		// "Authorization": {"Bearer <token>"},
	}
	res, err := client.Do(req)
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
