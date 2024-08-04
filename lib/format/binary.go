package format

import (
	"fmt"
	"strings"
)

type Binary struct {
	input string
}

func NewBinary(input string) (*Binary, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, fmt.Errorf("input cannot be empty")
	}

	return &Binary{
		input: input,
	}, nil
}

func (m *Binary) Convert() string {
	var result string
	for _, char := range m.input {
		binary := fmt.Sprintf("%b", char)
		result += binary + " "
	}

	return result
}

var _ Format = (*Binary)(nil)
