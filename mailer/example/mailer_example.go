package main

import (
	"fmt"

	"github.com/kelchy/go-lib/mailer"
)

func main() {
	// Initialize the mailer, credentials should not be hardcoded
	ses, err := mailer.NewSES(mailer.DefaultSESConfig(
		"AKIAIOSFODNN7EXAMPLE",
		"wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
		"ap-southeast-1",
		""),
		mailer.DefaultLogger())

	if err != nil {
		fmt.Println("Failed to start up mailer: ", err)
		return
	}

	ses.SendEmail(map[string]interface{}{})
}
