package createquote

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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

	zipcodeInt, err := strconv.Atoi(request.Recipient.Address.Zipcode)
	if err != nil {
		return FreightFastOutputDTO{}, fmt.Errorf("failed to convert zipcode to int: %w", err)
	}

	registeredNumber := os.Getenv("API_REGISTER_NUMBER")
	if registeredNumber == "" {
		return FreightFastOutputDTO{}, fmt.Errorf("API_REGISTER_NUMBER is not set")
	}

	shipperToken := os.Getenv("API_SHIPPER_TOKEN")
	if shipperToken == "" {
		return FreightFastOutputDTO{}, fmt.Errorf("API_SHIPPER_TOKEN is not set")
	}

	shipperPlatformCode := os.Getenv("API_SHIPPER_PLATFORM_CODE")
	if shipperPlatformCode == "" {
		return FreightFastOutputDTO{}, fmt.Errorf("API_SHIPPER_PLATFORM_CODE is not set")
	}

	recipientCountry := os.Getenv("API_RECIPIENT_COUNTRY")
	if recipientCountry == "" {
		return FreightFastOutputDTO{}, fmt.Errorf("API_RECIPIENT_COUNTRY is not set")
	}

	recipientTypeStr := os.Getenv("API_RECIPIENT_TYPE")
	recipientType, err := strconv.Atoi(recipientTypeStr)
	if err != nil {
		return FreightFastOutputDTO{}, fmt.Errorf("failed to convert API_RECIPIENT_TYPE to int: %w", err)
	}

	if recipientType < 0 || recipientType > 1 {
		return FreightFastOutputDTO{}, fmt.Errorf("API_RECIPIENT_TYPE is not valid")
	}

	// Montando o corpo da requisição
	requestBody := map[string]interface{}{
		"shipper": map[string]string{
			"registered_number": registeredNumber,
			"token":             shipperToken,
			"platform_code":     shipperPlatformCode,
		},
		"recipient": map[string]interface{}{
			"type":    recipientType,
			"country": recipientCountry,
			"zipcode": zipcodeInt,
		},
		"dispatchers": []map[string]interface{}{
			{
				"registered_number": registeredNumber,
				"zipcode":           zipcodeInt,
				"total_price":       0.0,
				"volumes": func() []map[string]interface{} {
					volumes := []map[string]interface{}{}
					for _, volume := range request.Volumes {

						categoryStr := strconv.Itoa(volume.Category)

						volumes = append(volumes, map[string]interface{}{
							"category":       categoryStr,
							"amount":         volume.Amount,
							"unitary_price":  volume.Price,
							"unitary_weight": volume.UnitaryWeight,
							"sku":            volume.SKU,
							"height":         volume.Height,
							"width":          volume.Width,
							"length":         volume.Length,
						})
					}
					return volumes
				}(),
			},
		},
		"simulation_type": []int{0},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return FreightFastOutputDTO{}, fmt.Errorf("failed to marshal request for Frete Rápido API: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", s.APIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return FreightFastOutputDTO{}, fmt.Errorf("failed to create request for Frete Rápido API: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return FreightFastOutputDTO{}, fmt.Errorf("failed to send request to Frete Rápido API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return FreightFastOutputDTO{}, fmt.Errorf("error querying Frete Rápido API: %s (status code: %d)", string(bodyBytes), resp.StatusCode)
	}

	var outputDTO FreightFastOutputDTO
	err = json.NewDecoder(resp.Body).Decode(&outputDTO)
	if err != nil {
		return FreightFastOutputDTO{}, fmt.Errorf("failed to decode response from Frete Rápido API: %w", err)
	}

	return outputDTO, nil
}
