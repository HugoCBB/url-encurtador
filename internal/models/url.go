package models

import "time"

type (
	UrlRecord struct {
		Id        string        `json:"id"`
		ShortCode string        `json:"short_code"`
		OldUrl    string        `json:"old_url"`
		Create_at string        `json:"create_at"`
		Exp       time.Duration `json:"exp"`
	}

	UrlRecordInput struct {
		Url string `json:"url"`
	}

	UrlRecordOutput struct {
		Url string `json:"url_converted"`
		Exp string `json:"exp"`
	}
)
