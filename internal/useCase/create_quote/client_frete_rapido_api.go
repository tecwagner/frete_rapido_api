package createquote

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type QuoteService struct {
	APIURL     string
	HTTPClient HTTPClient
}

func NewQuoteService(apiURL string, client HTTPClient) *QuoteService {
	return &QuoteService{
		APIURL:     apiURL,
		HTTPClient: client,
	}
}

func (s *QuoteService) GetQuoteFromFreightFast(ctx context.Context, request CreateQuoteInputDTO) (FreightFastOutputDTO, error) {

	requestBody, err := json.Marshal(request)
	if err != nil {
		return FreightFastOutputDTO{}, fmt.Errorf("failed to marshal request for Frete Rápido API: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", s.APIURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return FreightFastOutputDTO{}, fmt.Errorf("failed to create request for Frete Rápido API: %w", err)

	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return FreightFastOutputDTO{}, fmt.Errorf("failed to send request to Frete Rápido API: %w", err)

	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return FreightFastOutputDTO{}, errors.New("error querying Frete Rápido API: " + string(bodyBytes))
	}

	var outputDTO FreightFastOutputDTO
	err = json.NewDecoder(resp.Body).Decode(&outputDTO)
	if err != nil {
		return FreightFastOutputDTO{}, fmt.Errorf("failed to decode response from Frete Rápido API: %w", err)
	}

	return outputDTO, nil
}
