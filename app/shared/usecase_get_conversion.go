package shared

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/o4f6bgpac3/string-conversion/gen/shared"
	"github.com/o4f6bgpac3/string-conversion/lib/format"
)

func (s *Service) GetConversion(ctx context.Context, r *connect.Request[shared.GetConversionRequest]) (*connect.Response[shared.GetConversionResponse], error) {
	input := r.Msg.Input
	formats := r.Msg.Formats

	conversion := shared.Conversion{
		Input:   input,
		Formats: make([]*shared.Conversion_Format, len(formats)),
	}

	for _, outputFormat := range formats {
		var f format.Format
		var err error

		switch outputFormat {
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
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("unsupported output format: %s", outputFormat))
		}

		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}

		conversion.Formats = append(conversion.Formats, &shared.Conversion_Format{
			Name:  outputFormat,
			Value: f.Convert(),
		})
	}

	response := connect.NewResponse(&shared.GetConversionResponse{
		Conversion: &conversion,
	})

	return response, nil
}
