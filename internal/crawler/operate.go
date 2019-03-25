package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"bingwall/internal/entity"
)

func getImageInfos() (*entity.Api, error) {
	resp, err := http.Get(rootUrl + apiPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %v", resp.StatusCode)
	}
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := &entity.Api{}
	if err := json.Unmarshal(buffer, result); err != nil {
		return nil, err
	}
	return result, err
}

func downloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %v", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
