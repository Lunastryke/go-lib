package mailer

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

// SESClient is the interface for SES clients
type SESClient struct {
	Client *sesv2.Client
}

// SendEmail sends an email using SES
func (s *SESClient) sendEmail(map[string]interface{}) error {
	return nil
}

func sesInit(cfg aws.Config) (*SESClient, error) {
	sesClient := SESClient{}
	sesClient.Client = sesv2.NewFromConfig(cfg)
	return &sesClient, nil
}
