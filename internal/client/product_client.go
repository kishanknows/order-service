package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kishanknows/order-service/internal/dto"
)

type ProductClient struct {
	baseURL string
	httpClient *http.Client
}

func NewProductClient(baseURL string) *ProductClient {
	return &ProductClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (p *ProductClient) GetProduct(ctx context.Context, productID string) (*dto.ProductData, error) {
	url := p.baseURL + productID

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("product service returned with failure")
		return nil, err
	}

	var productResponse dto.ProductResponse
	if err := json.NewDecoder(resp.Body).Decode(&productResponse); err != nil {
		return nil, err
	}

	return &productResponse.Data, nil
}

func (p *ProductClient) UpdateProduct(ctx context.Context, productID string, stock int) error {
	body, err := json.Marshal(map[string]any{
		"stock": stock,
	})

	if err != nil {
		return err
	}

	url := p.baseURL + productID

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set(
		"Authorization", 
		"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6Imtpc2hhbmtub3dzQGdtYWlsLmNvbSIsInJvbGUiOiJhZG1pbiIsImV4cCI6MTc4MjcyNTc3NywiaWF0IjoxNzgyNjM5Mzc3fQ.om_PPFuW7xOvxHAVwptrEs2oBF_eGRcSCR2gVZ_hjDM",
	)

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("product service returned with failure")
		return err
	}

	return nil
}