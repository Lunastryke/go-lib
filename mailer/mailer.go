package mailer

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"golang.org/x/time/rate"
)

// Mailer is the mailer struct
type Mailer struct {
	Client      IClient      // Client for the mailer. Interface as we can have multiple different clients
	rateLimiter rate.Limiter // Rate limiter for the mailer
	logger      ILogger      // Logger for the mailer
}
type IMailer interface {
	SendEmail(map[string]interface{}) error
}

type IClient interface {
	sendEmail(map[string]interface{}) error
}

// ILogger is the interface for logger
type ILogger interface {
	Debug(key string, message string)
	Out(key string, message string)
	Error(key string, err error)
}

// New SES creates a new SES client.
func NewSES(cfg aws.Config, logger ILogger) (*Mailer, error) {
	mailer := Mailer{}
	mailer.logger = logger
	// Default rate limit is 1 event per second
	mailer.rateLimiter = *rate.NewLimiter(rate.Limit(1), 1)
	sesClient, err := sesInit(cfg)
	if err != nil {
		logger.Error("SES Client Init", err)
		return nil, err
	}
	mailer.Client = sesClient
	return &mailer, nil
}

// SetRateLimit sets the rate limit for the mailer. limit is the number of events allowed per second
func (m *Mailer) SetRateLimit(limit float64) error {
	m.rateLimiter.SetLimit(rate.Limit(limit))
	return nil
}

// SendEmail sends an email using the mailer
func (m *Mailer) SendEmail(request map[string]interface{}) error {
	m.Client.sendEmail(request)
	return nil
}
