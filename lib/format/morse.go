package format

import (
	"fmt"
	"strings"
)

type Morse struct {
	input string
}

func NewMorse(input string) (*Morse, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, fmt.Errorf("input cannot be empty")
	}

	return &Morse{
		input: input,
	}, nil
}

func (m *Morse) Convert() string {
	var result string
	for _, char := range m.input {
		upperChar := strings.ToUpper(string(char))
		if upperChar == " " {
			result += " "
			continue
		}
		if morseCode[upperChar] == "" {
			result += "?"
			continue
		}
		result += morseCode[upperChar] + " "
	}
	return result
}

var morseCode = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	"0": "-----",
}

var _ Format = (*Morse)(nil)
