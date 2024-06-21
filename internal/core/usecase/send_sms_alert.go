package Usecase

import (
	"fmt"
	"os"

	"github.com/Domains18/food-delivery/internal/core/models"
	"github.com/Domains18/food-delivery/internal/core/repository"
	"github.com/Domains18/food-delivery/pkg"
	"github.com/MikeMwita/africastalking-go/pkg/sms"
)

// SendSmsAlert sends an SMS alert to a customer when an order is added
func SendSmsAlert(order models.Order) error {
	client := sms.SmsSender{
		ApiKey:  os.Getenv("AFRICASTALKING_API_KEY"),
		ApiUser: os.Getenv("AFRICASTALKING_API_USER"),
		Sender:  os.Getenv("AFRICASTALKING_SENDER"),
	}

	customerRepo := repository.NewCustomerRepo(pkg.DB)
	customer, err := customerRepo.GetCustomerByID(order.CustomerID)
	if err != nil {
		return err
	}

	// Construct the SMS message with personalized information
	message := fmt.Sprintf("Hello %s, your order for %s has been placed successfully. Thank you for choosing us.", customer.Name, order.Item)

	client.Recipients = []string{customer.Phone}

	client.Message = message

	// Send the SMS
	response, err := client.SendSMS()
	if err != nil {
		return err
	}

	fmt.Printf("SMS Response: %+v\n", response)

	return nil
}
