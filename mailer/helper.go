package mailer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/kelchy/go-lib/log"
)

// DefaultConfig returns a default configuration for the SES mailer
func DefaultSESConfig(awsAccessKeyId string, awsSecretAccessKey string, region string, sessionToken string) aws.Config {
	creds := credentials.NewStaticCredentialsProvider(
		awsAccessKeyId,
		awsSecretAccessKey,
		sessionToken,
	)

	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(region),
	)

	if err != nil {
		fmt.Println("failed to load default config: ", err)
		return aws.Config{}
	}

	return cfg
}

// DefaultLogger returns a default logger for the package
func DefaultLogger() ILogger {
	logger, err := log.New("standard")
	if err != nil {
		fmt.Println("failed to create logger: ", err)
	}
	return logger
}
