package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func LookupIP(ip string) (*IPInfo, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(
		fmt.Sprintf("https://ipinfo.io/%s/json", ip),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to fetch IP info: %w",
			err,
		)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"unexpected status code: %d",
			resp.StatusCode,
		)
	}

	var data IPInfo

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf(
			"failed to decode IP info: %w",
			err,
		)
	}

	if data.Bogon {
		return nil, fmt.Errorf(
			"bogon IP address: %s",
			data.IP,
		)
	}

	return &data, nil
}
