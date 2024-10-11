package createquote

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

const freightfastURL = "https://sp.freterapido.com/api/v3/quote/simulate"

func GetQuoteFromFreightFast(request CreateQuoteInputDTO) (FreightFastOutputDTO, error) {

	requestBody, err := json.Marshal(request)
	if err != nil {
		return FreightFastOutputDTO{}, err
	}

	req, err := http.NewRequest("POST", freightfastURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return FreightFastOutputDTO{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return FreightFastOutputDTO{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return FreightFastOutputDTO{}, errors.New("error querying Frete Rápido API: " + string(bodyBytes))
	}

	var outputDTO FreightFastOutputDTO
	err = json.NewDecoder(resp.Body).Decode(&outputDTO)
	if err != nil {
		return outputDTO, err
	}

	return outputDTO, nil
}
