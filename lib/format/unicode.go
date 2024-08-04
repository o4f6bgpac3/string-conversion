package format

import (
	"fmt"
	"strings"
)

type Unicode struct {
	input string
}

func NewUnicode(input string) (Format, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, fmt.Errorf("input cannot be empty")
	}

	return &Unicode{
		input: input,
	}, nil
}

func (f *Unicode) Convert() string {
	var result string
	for _, char := range f.input {
		result += fmt.Sprintf("U+%04X ", char)
	}

	return result
}

var _ Format = (*Unicode)(nil)
