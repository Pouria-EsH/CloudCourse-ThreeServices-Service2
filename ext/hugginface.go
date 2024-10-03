package ext

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const VITGPT2_APIURL = "https://api-inference.huggingface.co/models/nlpconnect/vit-gpt2-image-captioning"

type HuggingFace struct {
	apikey string
}

func (hf HuggingFace) GetDiscription(fileBytes *bytes.Buffer) (string, error) {
	if fileBytes == nil {
		return "", fmt.Errorf("filebyte is null")
	}

	resp, err := hf.sendHFHttpRequest(fileBytes)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	return hf.parseResponse(respBody)
}

func (hf HuggingFace) sendHFHttpRequest(img io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", VITGPT2_APIURL, img)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", hf.apikey))
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{}
	return client.Do(req)
}

func (hf HuggingFace) parseResponse(resp []byte) (string, error) {
	type result struct {
		Text string `json:"generated_text"`
	}
	type errorlist struct {
		Errorlist []string `json:"error"`
	}

	var results []result
	if err := json.Unmarshal(resp, &results); err != nil {
		var errs errorlist
		if err := json.Unmarshal(resp, &errs); err != nil {
			return "", fmt.Errorf("error decoding JSON: %w", err)
		}
		return "", fmt.Errorf("inference Error: %s", errs.Errorlist)
	}

	return results[0].Text, nil
}
