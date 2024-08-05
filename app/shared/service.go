package shared

import "github.com/o4f6bgpac3/string-conversion/gen/shared/sharedconnect"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

var _ sharedconnect.StringConversionServiceHandler = (*Service)(nil)
