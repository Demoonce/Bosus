package weeb

import (
	"net/http"
)

const (
	api = "Bearer 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"
	url = "https://api-inference.huggingface.co/models/iZELX1/Anything-V3-X"
)

var (
	Client = http.DefaultClient
	Images_Generated = 0
)
