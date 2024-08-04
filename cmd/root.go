package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/o4f6bgpac3/string-conversion/lib/format"
	"github.com/spf13/cobra"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	formats string
	input   string
)

var rootCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a string to one or more formats",
	Long: `Convert a given input string to one or more specified output formats.
Currently supported output formats are:
- ascii
- base64
- binary
- hex
- morse
- unicode
- url-encoding`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if input == "" {
			return fmt.Errorf("input string is required. Use --input or -i flag")
		}

		if formats == "" {
			return fmt.Errorf("at least one output format is required. Use --formats or -f flag")
		}

		formatList := strings.Split(formats, ",")
		for i, f := range formatList {
			formatList[i] = strings.TrimSpace(f)
		}

		results := make(map[string]string)

		var f format.Format
		var err error
		for _, outputFormat := range formatList {
			switch strings.ToLower(outputFormat) {
			case "ascii":
				f, err = format.NewASCII(input)
			case "base64":
				f, err = format.NewBase64(input)
			case "binary":
				f, err = format.NewBinary(input)
			case "hex":
				f, err = format.NewHexadecimal(input)
			case "morse":
				f, err = format.NewMorse(input)
			case "unicode":
				f, err = format.NewUnicode(input)
			case "url-encoding":
				f, err = format.NewUrlEncoding(input)
			default:
				return fmt.Errorf("unsupported output format: %s. Supported formats are 'ascii', 'base64, 'binary', 'hex', 'morse', 'unicode' and 'url-encoding", outputFormat)
			}

			if err != nil {
				return err
			}

			results[outputFormat] = f.Convert()
		}

		titleCaser := cases.Title(language.English)

		fmt.Printf("Input: %s\n\n", input)
		for format, result := range results {
			fmt.Printf("%s:\n%s\n\n", titleCaser.String(format), result)
		}

		return nil
	},
}

func init() {
	rootCmd.Flags().StringVarP(&formats, "formats", "f", "", "Comma-separated list of output formats (required)")
	rootCmd.Flags().StringVarP(&input, "input", "i", "", "Input string to convert (required)")
	rootCmd.MarkFlagRequired("formats")
	rootCmd.MarkFlagRequired("input")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
