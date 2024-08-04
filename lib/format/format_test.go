package format

import (
	"testing"
)

func TestASCII(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "72 101 108 108 111 "},
		{"123", "49 50 51 "},
		{"Café", "67 97 102 233 "},
	}

	for _, test := range tests {
		ascii, err := NewASCII(test.input)
		if err != nil {
			t.Errorf("NewASCII(%q) returned an error: %v", test.input, err)
			continue
		}

		result := ascii.Convert()
		if result != test.expected {
			t.Errorf("ASCII.Convert() for input %q = %q, want %q", test.input, result, test.expected)
		}
	}
}

func TestBase64(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "SGVsbG8="},
		{"123", "MTIz"},
		{"Café", "Q2Fmw6k="},
	}

	for _, test := range tests {
		base64, err := NewBase64(test.input)
		if err != nil {
			t.Errorf("NewBase64(%q) returned an error: %v", test.input, err)
			continue
		}

		result := base64.Convert()
		if result != test.expected {
			t.Errorf("Base64.Convert() for input %q = %q, want %q", test.input, result, test.expected)
		}
	}
}

func TestBinary(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"A", "1000001 "},
		{"123", "110001 110010 110011 "},
		{"Hi!", "1001000 1101001 100001 "},
	}

	for _, test := range tests {
		binary, err := NewBinary(test.input)
		if err != nil {
			t.Errorf("NewBinary(%q) returned an error: %v", test.input, err)
			continue
		}

		result := binary.Convert()
		if result != test.expected {
			t.Errorf("Binary.Convert() for input %q = %q, want %q", test.input, result, test.expected)
		}
	}
}

func TestHexadecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "48 65 6c 6c 6f "},
		{"123", "31 32 33 "},
		{"Café", "43 61 66 e9 "},
	}

	for _, test := range tests {
		hex, err := NewHexadecimal(test.input)
		if err != nil {
			t.Errorf("NewHexadecimal(%q) returned an error: %v", test.input, err)
			continue
		}

		result := hex.Convert()
		if result != test.expected {
			t.Errorf("Hexadecimal.Convert() for input %q = %q, want %q", test.input, result, test.expected)
		}
	}
}

func TestMorse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"SOS", "... --- ... "},
		{"Hello World", ".... . .-.. .-.. ---  .-- --- .-. .-.. -.. "},
		{"123", ".---- ..--- ...-- "},
	}

	for _, test := range tests {
		morse, err := NewMorse(test.input)
		if err != nil {
			t.Errorf("NewMorse(%q) returned an error: %v", test.input, err)
			continue
		}

		result := morse.Convert()
		if result != test.expected {
			t.Errorf("Morse.Convert() for input %q = %q, want %q", test.input, result, test.expected)
		}
	}
}

func TestUnicode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "U+0048 U+0065 U+006C U+006C U+006F "},
		{"123", "U+0031 U+0032 U+0033 "},
		{"Café", "U+0043 U+0061 U+0066 U+00E9 "},
	}

	for _, test := range tests {
		unicode, err := NewUnicode(test.input)
		if err != nil {
			t.Errorf("NewUnicode(%q) returned an error: %v", test.input, err)
			continue
		}

		result := unicode.Convert()
		if result != test.expected {
			t.Errorf("Unicode.Convert() for input %q = %q, want %q", test.input, result, test.expected)
		}
	}
}

func TestUrlEncoding(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World", "Hello+World"},
		{"a=1&b=2", "a%3D1%26b%3D2"},
		{"Café", "Caf%C3%A9"},
	}

	for _, test := range tests {
		urlEnc, err := NewUrlEncoding(test.input)
		if err != nil {
			t.Errorf("NewUrlEncoding(%q) returned an error: %v", test.input, err)
			continue
		}

		result := urlEnc.Convert()
		if result != test.expected {
			t.Errorf("UrlEncoding.Convert() for input %q = %q, want %q", test.input, result, test.expected)
		}
	}
}

func TestEmptyInput(t *testing.T) {
	formats := []struct {
		name string
		fn   func(string) (Format, error)
	}{
		{"ASCII", NewASCII},
		{"Base64", NewBase64},
		{"Binary", NewBinary},
		{"Hexadecimal", NewHexadecimal},
		{"Morse", NewMorse},
		{"Unicode", NewUnicode},
		{"UrlEncoding", NewUrlEncoding},
	}

	for _, format := range formats {
		_, err := format.fn("")
		if err == nil {
			t.Errorf("%s: expected error for empty input, got nil", format.name)
		}
	}
}
