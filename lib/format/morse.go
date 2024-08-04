package format

import (
	"fmt"
	"strings"
	"unicode"
)

type Morse struct {
	input string
}

func NewMorse(input string) (Format, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, fmt.Errorf("input cannot be empty")
	}

	return &Morse{
		input: input,
	}, nil
}

func (f *Morse) Convert() string {
	var result string
	for _, char := range f.input {
		upperChar := unicode.ToUpper(char)
		if upperChar == ' ' {
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

var morseCode = map[rune]string{
	'A':  ".-",
	'B':  "-...",
	'C':  "-.-.",
	'D':  "-..",
	'E':  ".",
	'F':  "..-.",
	'G':  "--.",
	'H':  "....",
	'I':  "..",
	'J':  ".---",
	'K':  "-.-",
	'L':  ".-..",
	'M':  "--",
	'N':  "-.",
	'O':  "---",
	'P':  ".--.",
	'Q':  "--.-",
	'R':  ".-.",
	'S':  "...",
	'T':  "-",
	'U':  "..-",
	'V':  "...-",
	'W':  ".--",
	'X':  "-..-",
	'Y':  "-.--",
	'Z':  "--..",
	'0':  "-----",
	'1':  ".----",
	'2':  "..---",
	'3':  "...--",
	'4':  "....-",
	'5':  ".....",
	'6':  "-....",
	'7':  "--...",
	'8':  "---..",
	'9':  "----.",
	'.':  ".-.-.-",
	',':  "--..--",
	'?':  "..--..",
	'\'': ".----.",
	'!':  "-.-.--",
	'/':  "-..-.",
	'(':  "-.--.",
	')':  "-.--.-",
	'&':  ".-...",
	':':  "---...",
	';':  "-.-.-.",
	'=':  "-...-",
	'+':  ".-.-.",
	'-':  "-....-",
	'_':  "..--.-",
	'"':  ".-..-.",
	'$':  "...-..-",
	'@':  ".--.-.",
	' ':  " ",
}

var _ Format = (*Morse)(nil)
