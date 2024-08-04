package format

import (
	"fmt"
	"strings"
)

type ASCII struct {
	input string
}

func NewASCII(input string) (Format, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, fmt.Errorf("input cannot be empty")
	}

	return &ASCII{
		input: input,
	}, nil
}

func (f *ASCII) Convert() string {
	var result string
	for _, char := range f.input {
		result += fmt.Sprintf("%d ", char)
	}

	return result
}

var _ Format = (*ASCII)(nil)
