package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"time"

	"mySafarHighLevel/internal/config"
	"mySafarHighLevel/internal/dto"
	"mySafarHighLevel/internal/util"
)

type Client struct {
	cfg    *config.Config
	client *http.Client
}

func New(cfg *config.Config) *Client {
	return &Client{
		cfg: cfg,
		client: &http.Client{
			Timeout: time.Duration(cfg.Timeout) * time.Second,
		},
	}
}

// GetRecommendations sends POST request to API and returns response
func (c *Client) GetRecommendations(ctx context.Context) (*dto.RecommendationResponse, error) {
	payload := dto.Payload{
    Token:        c.cfg.Token,
    IsBaggage:    c.cfg.IsBaggage,
    Lang:         c.cfg.Lang,
    FilterAirlines: nil,
    IsDirectOnly: func() int {
        if c.cfg.IsDirectOnly {
            return 1
        }
        return 0
    }(),
    Src:          0,
    Yth:          0,
    Inf:          0,
    Ins:          0,
    Segments: []dto.Segment{{
        Date: c.cfg.Date,
        From: c.cfg.From,
        To:   c.cfg.To,
    }},
    IsCharter:    c.cfg.IsCharter,
    PriceOrder:   1,
    ArrOrder:     1,
    DepOrder:     1,
    DurationOrder: 1,
    Adt:          c.cfg.Adt,
    Chd:          c.cfg.Chd,
    GdsWhiteList: []string{},
    GdsBlackList: nil,
    Count:        c.cfg.Count,
    Class_:       c.cfg.Class,
}



	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal payload error: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.cfg.URL, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+c.cfg.Token)

	if c.cfg.Debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Println("----- REQUEST DUMP -----")
		fmt.Println(util.MaskToken(string(dump)))
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if c.cfg.Debug {
		dump, _ := httputil.DumpResponse(resp, true)
		fmt.Println("----- RESPONSE DUMP -----")
		fmt.Println(string(dump))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API error: status %d, body=%s", resp.StatusCode, string(body))
	}

	var parsed dto.RecommendationResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return &parsed, nil
}
