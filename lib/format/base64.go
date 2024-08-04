package format

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type Base64 struct {
	input string
}

func NewBase64(input string) (*Base64, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, fmt.Errorf("input cannot be empty")
	}

	return &Base64{
		input: input,
	}, nil
}

func (f *Base64) Convert() string {
	return base64.StdEncoding.EncodeToString([]byte(f.input))
}
