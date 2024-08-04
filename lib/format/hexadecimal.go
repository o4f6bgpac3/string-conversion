package format

import (
	"fmt"
	"strings"
)

type Hexadecimal struct {
	input string
}

func NewHexadecimal(input string) (Format, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, fmt.Errorf("input cannot be empty")
	}

	return &Hexadecimal{
		input: input,
	}, nil
}

func (f *Hexadecimal) Convert() string {
	var result string
	for _, char := range f.input {
		hex := fmt.Sprintf("%x", char)
		result += hex + " "
	}

	return result
}

var _ Format = (*Hexadecimal)(nil)
