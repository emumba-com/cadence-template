package communicationService

import (
	"context"
	"fmt"
)

func SendEmail(ctx context.Context, userID string) error {
	fmt.Println("Email sent successfully to userID:", userID)

	return nil
}

func SendSMS(ctx context.Context, userID string) error {
	fmt.Println("SMS sent successfully to userID:", userID)

	return nil
}
