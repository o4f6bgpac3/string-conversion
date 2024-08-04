package format

import (
	"fmt"
	"net/url"
	"strings"
)

type UrlEncoding struct {
	input string
}

func NewUrlEncoding(input string) (*UrlEncoding, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, fmt.Errorf("input cannot be empty")
	}

	return &UrlEncoding{
		input: input,
	}, nil
}

func (f *UrlEncoding) Convert() string {
	return url.QueryEscape(f.input)
}

var _ Format = (*UrlEncoding)(nil)
